package routes

import (
	"net/http"

	"github.com/formancehq/ledger/pkg/api/controllers"
	"github.com/formancehq/ledger/pkg/api/middlewares"
	"github.com/formancehq/ledger/pkg/opentelemetry/metrics"
	"github.com/formancehq/stack/libs/go-libs/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/riandyrn/otelchi"
)

func NewRouter(
	backend controllers.Backend,
	healthController *health.HealthController,
	globalMetricsRegistry metrics.GlobalRegistry,
) chi.Router {
	router := chi.NewMux()

	router.Use(
		cors.New(cors.Options{
			AllowOriginFunc: func(r *http.Request, origin string) bool {
				return true
			},
			AllowCredentials: true,
		}).Handler,
		middlewares.MetricsMiddleware(globalMetricsRegistry),
		middleware.Recoverer,
	)

	router.Get("/_healthcheck", healthController.Check)

	router.Group(func(router chi.Router) {
		router.Use(otelchi.Middleware("ledger"))
		router.Get("/_info", controllers.GetInfo(backend))

		router.Route("/{ledger}", func(router chi.Router) {
			router.Use(func(handler http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					handler.ServeHTTP(w, r)
				})
			})
			router.Use(middlewares.LedgerMiddleware(backend))

			// LedgerController
			router.Get("/_info", controllers.GetLedgerInfo)
			router.Get("/stats", controllers.GetStats)
			router.Get("/logs", controllers.GetLogs)

			// AccountController
			router.Get("/accounts", controllers.GetAccounts)
			router.Head("/accounts", controllers.CountAccounts)
			router.Get("/accounts/{address}", controllers.GetAccount)
			router.Post("/accounts/{address}/metadata", controllers.PostAccountMetadata)

			// TransactionController
			router.Get("/transactions", controllers.GetTransactions)
			router.Head("/transactions", controllers.CountTransactions)

			router.Post("/transactions", controllers.PostTransaction)

			router.Get("/transactions/{txid}", controllers.GetTransaction)
			router.Post("/transactions/{txid}/revert", controllers.RevertTransaction)
			router.Post("/transactions/{txid}/metadata", controllers.PostTransactionMetadata)

			// BalanceController
			router.Get("/balances", controllers.GetBalances)
			// TODO: Rename to /aggregatedBalances
			router.Get("/aggregate/balances", controllers.GetBalancesAggregated)
		})
	})

	return router
}
