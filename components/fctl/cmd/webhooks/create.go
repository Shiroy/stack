package webhooks

import (
	"fmt"
	"net/url"

	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	secretFlag = "secret"
)

type CreateWebhookController struct {
	store *CreateWebhookStore
}

type CreateWebhookStore struct {
	Webhook shared.WebhooksConfig `json:"webhook"`
}

var _ fctl.Controller[*CreateWebhookStore] = (*CreateWebhookController)(nil)

func NewDefaultCreateWebhookStore() *CreateWebhookStore {
	return &CreateWebhookStore{
		Webhook: shared.WebhooksConfig{},
	}
}

func NewCreateWebhookController() *CreateWebhookController {
	return &CreateWebhookController{
		store: NewDefaultCreateWebhookStore(),
	}
}
func (c *CreateWebhookController) GetStore() *CreateWebhookStore {
	return c.store
}

func (c *CreateWebhookController) Run(cmd *cobra.Command, args []string) (fctl.Renderable, error) {
	cfg, err := fctl.GetConfig(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "fctl.GetConfig")
	}

	organizationID, err := fctl.ResolveOrganizationID(cmd, cfg)
	if err != nil {
		return nil, err
	}

	stack, err := fctl.ResolveStack(cmd, cfg, organizationID)
	if err != nil {
		return nil, err
	}

	if !fctl.CheckStackApprobation(cmd, stack, "You are about to create a webhook") {
		return nil, fctl.ErrMissingApproval
	}

	client, err := fctl.NewStackClient(cmd, cfg, stack)
	if err != nil {
		return nil, errors.Wrap(err, "creating stack client")
	}

	if _, err := url.Parse(args[0]); err != nil {
		return nil, errors.Wrap(err, "invalid endpoint URL")
	}

	secret := fctl.GetString(cmd, secretFlag)

	response, err := client.Webhooks.InsertConfig(cmd.Context(), shared.ConfigUser{
		Endpoint:   args[0],
		EventTypes: args[1:],
		Secret:     &secret,
	})

	if err != nil {
		return nil, errors.Wrap(err, "creating config")
	}

	if response.ErrorResponse != nil {
		return nil, fmt.Errorf("%s: %s", response.ErrorResponse.ErrorCode, response.ErrorResponse.ErrorMessage)
	}

	if response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	c.store.Webhook = response.ConfigResponse.Data

	return c, nil
}

func (c *CreateWebhookController) Render(cmd *cobra.Command, args []string) error {
	pterm.Success.WithWriter(cmd.OutOrStdout()).Printfln("Config created successfully")
	return nil
}

func NewCreateCommand() *cobra.Command {

	return fctl.NewCommand("create <endpoint> [<event-type>...]",
		fctl.WithShortDescription("Create a new config. At least one event type is required."),
		fctl.WithAliases("cr"),
		fctl.WithConfirmFlag(),
		fctl.WithArgs(cobra.MinimumNArgs(2)),
		fctl.WithStringFlag(secretFlag, "", "Bring your own webhooks signing secret. If not passed or empty, a secret is automatically generated. The format is a string of bytes of size 24, base64 encoded. (larger size after encoding)"),
		fctl.WithController[*CreateWebhookStore](NewCreateWebhookController()),
	)
}
