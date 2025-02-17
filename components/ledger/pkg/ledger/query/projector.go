package query

import (
	"context"
	"fmt"
	"sync"

	"github.com/formancehq/ledger/pkg/core"
	"github.com/formancehq/ledger/pkg/ledger/utils/batching"
	"github.com/formancehq/ledger/pkg/opentelemetry/metrics"
	storageerrors "github.com/formancehq/ledger/pkg/storage"
	"github.com/formancehq/stack/libs/go-libs/collectionutils"
	"github.com/formancehq/stack/libs/go-libs/logging"
)

type logPersistenceParts struct {
	mu           sync.Mutex
	parts        map[string]struct{}
	onTerminated func()
}

func (p *logPersistenceParts) Store(id string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.parts[id] = struct{}{}
}

func (p *logPersistenceParts) Delete(id string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.parts, id)

	if len(p.parts) == 0 {
		p.onTerminated()
	}
}

func newLogPersistenceParts(onTerminated func()) *logPersistenceParts {
	return &logPersistenceParts{
		mu:           sync.Mutex{},
		parts:        map[string]struct{}{},
		onTerminated: onTerminated,
	}
}

type Projector struct {
	store           Store
	monitor         Monitor
	metricsRegistry metrics.PerLedgerRegistry

	queue      chan []*core.ActiveLog
	stopChan   chan chan struct{}
	activeLogs *collectionutils.LinkedList[*core.ActiveLog]

	txUpdater              *batching.Batcher[core.Transaction]
	txMetadataUpdater      *metadataUpdater
	accountMetadataUpdater *metadataUpdater

	moveUpdater   *moveBuffer
	limitReadLogs int
}

func (p *Projector) QueueLog(logs ...*core.ActiveLog) {
	p.queue <- logs
}

func (p *Projector) Stop(ctx context.Context) {
	logger := logging.FromContext(ctx).WithField("component", "projector")
	logger.Infof("Stop")
	ch := make(chan struct{})
	p.stopChan <- ch
	<-ch
}

func (p *Projector) Start(ctx context.Context) {
	logger := logging.FromContext(ctx).WithField("component", "projector")
	logger.Infof("Start")

	ctx = logging.ContextWithLogger(ctx, logger)

	go p.moveUpdater.Run(logging.ContextWithField(ctx, "component", "moves buffer"))
	go p.txUpdater.Run(logging.ContextWithField(ctx, "component", "transactions buffer"))
	go p.accountMetadataUpdater.Run(logging.ContextWithField(ctx, "component", "accounts metadata buffer"))
	go p.txMetadataUpdater.Run(logging.ContextWithField(ctx, "component", "transactions metadata buffer"))

	p.syncLogs(ctx)

	go func() {

		for {
			select {
			case ch := <-p.stopChan:
				logger.Debugf("Close move buffer")
				p.moveUpdater.Close()

				logger.Debugf("Stop transaction worker")
				p.txUpdater.Close()

				logger.Debugf("Stop account metadata worker")
				p.accountMetadataUpdater.Close()

				logger.Debugf("Stop transaction metadata worker")
				p.txMetadataUpdater.Close()

				close(ch)
				return
			case logs := <-p.queue:
				logger.Debugf("Got %d new logs to project", len(logs))
				p.processLogs(ctx, logs)
			}
		}
	}()
}

func (p *Projector) syncLogs(ctx context.Context) error {
	lastReadLogID, err := p.store.GetNextLogID(ctx)
	if err != nil && !storageerrors.IsNotFoundError(err) {
		panic(err)
	}

	logging.FromContext(ctx).Infof("Project logs since id: %d", lastReadLogID)

	for {
		logs, err := p.store.ReadLogsRange(ctx, lastReadLogID, lastReadLogID+uint64(p.limitReadLogs))
		if err != nil {
			panic(err)
		}

		if len(logs) == 0 {
			// No logs, nothing to do
			return nil
		}

		p.processLogs(ctx, collectionutils.Map(logs, func(from core.ChainedLog) *core.ActiveLog {
			return core.NewActiveLog(&from)
		}))

		lastReadLogID = logs[len(logs)-1].ID + 1

		if len(logs) < p.limitReadLogs {
			// Nothing to do anymore, no need to read more logs
			return nil
		}
	}
}

func (p *Projector) processLogs(ctx context.Context, logs []*core.ActiveLog) {
	p.metricsRegistry.QueryInboundLogs().Add(ctx, int64(len(logs)))
	p.activeLogs.Append(logs...)

	for _, log := range logs {
		log := log
		markLogAsProjected := func() {
			log.SetProjected()
			if err := p.store.MarkedLogsAsProjected(ctx, log.ID); err != nil {
				panic(err)
			}
			p.metricsRegistry.QueryProcessedLogs().Add(ctx, 1)
		}
		dispatchTransaction := func(l *logPersistenceParts, log *core.ActiveLog, tx core.Transaction) {
			logger := logging.FromContext(ctx).WithFields(map[string]any{
				"log-id": log.ID,
			})
			moves := tx.GetMoves()
			moveKey := func(move *core.Move) string {
				return fmt.Sprintf("move/%d/%v/%s", move.PostingIndex, move.IsSource, move.Account)
			}
			l.Store("tx")
			for _, move := range moves {
				l.Store(moveKey(move))
			}

			p.txUpdater.Append(tx, func() {
				logger.Debugf("Transaction projected")
				l.Delete("tx")
			})

			for _, move := range moves {
				move := move
				p.moveUpdater.AppendMove(move, func() {
					logger.WithFields(map[string]any{
						"asset":     move.Asset,
						"is_source": move.IsSource,
						"account":   move.Account,
					}).Debugf("Move projected")
					l.Delete(moveKey(move))
				})
			}
		}
		switch payload := log.Log.Data.(type) {
		case core.NewTransactionLogPayload:
			l := newLogPersistenceParts(func() {
				markLogAsProjected()
				p.monitor.CommittedTransactions(ctx, *payload.Transaction)
			})

			dispatchTransaction(l, log, *payload.Transaction)
		case core.SetMetadataLogPayload:
			switch payload.TargetType {
			case core.MetaTargetTypeAccount:
				p.accountMetadataUpdater.Append(payload.TargetID, payload.Metadata, func() {
					markLogAsProjected()
					p.monitor.SavedMetadata(ctx, payload.TargetType, fmt.Sprint(payload.TargetID), payload.Metadata)
				})
			case core.MetaTargetTypeTransaction:
				p.txMetadataUpdater.Append(payload.TargetID, payload.Metadata, func() {
					markLogAsProjected()
					p.monitor.SavedMetadata(ctx, payload.TargetType, fmt.Sprint(payload.TargetID), payload.Metadata)
				})
			}
		case core.RevertedTransactionLogPayload:
			l := newLogPersistenceParts(func() {
				markLogAsProjected()
				p.activeLogs.RemoveValue(log)

				revertedTx, err := p.store.GetTransaction(ctx, payload.RevertedTransactionID)
				if err != nil {
					panic(err)
				}
				p.monitor.RevertedTransaction(ctx, payload.RevertTransaction, &revertedTx.Transaction)
			})
			l.Store("metadata")
			dispatchTransaction(l, log, *payload.RevertTransaction)
			p.txMetadataUpdater.Append(payload.RevertedTransactionID, core.RevertedMetadata(payload.RevertTransaction.ID), func() {
				l.Delete("metadata")
			})
		}
	}
}

func NewProjector(store Store, monitor Monitor, metricsRegistry metrics.PerLedgerRegistry) *Projector {
	return &Projector{
		store:           store,
		monitor:         monitor,
		metricsRegistry: metricsRegistry,
		txUpdater: batching.NewBatcher(
			store.InsertTransactions,
			batching.NoOpOnBatchProcessed[core.Transaction](),
			2,
			512,
		),
		accountMetadataUpdater: newMetadataUpdater(func(ctx context.Context, update ...*MetadataUpdate) error {
			return store.UpdateAccountsMetadata(ctx, collectionutils.Map(update, func(from *MetadataUpdate) core.Account {
				return core.Account{
					Address:  from.ID.(string),
					Metadata: from.Metadata,
				}
			})...)
		}, 1, 512),
		txMetadataUpdater: newMetadataUpdater(func(ctx context.Context, update ...*MetadataUpdate) error {
			return store.UpdateTransactionsMetadata(ctx, collectionutils.Map(update, func(from *MetadataUpdate) core.TransactionWithMetadata {
				return core.TransactionWithMetadata{
					ID:       from.ID.(uint64),
					Metadata: from.Metadata,
				}
			})...)
		}, 1, 512),
		moveUpdater:   newMoveUpdater(store.InsertMoves, 5, 100),
		activeLogs:    collectionutils.NewLinkedList[*core.ActiveLog](),
		queue:         make(chan []*core.ActiveLog, 1024),
		stopChan:      make(chan chan struct{}),
		limitReadLogs: 10000,
	}
}
