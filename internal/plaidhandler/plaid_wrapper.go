package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
)

type PlaidWrapper struct {
	Client             *plaid.APIClient
	InstrumentationSdk *instrumentation.ServiceTelemetry
	Logger             *zap.Logger
}

// NewPlaidWrapper returns a pointer reference to the plaid wrapper object
func NewPlaidWrapper(plaidClient *plaid.APIClient, instrumentation *instrumentation.ServiceTelemetry, logger *zap.Logger) (*PlaidWrapper, error) {
	if plaidClient == nil || logger == nil {
		return nil, errors.New("invalid input arguments. check input parameters")
	}

	return &PlaidWrapper{
		Client:             plaidClient,
		InstrumentationSdk: instrumentation,
		Logger:             logger,
	}, nil
}

func (p *PlaidWrapper) GetAccessTokenForSandboxAcct() (string, error) {
	ctx := context.Background()

	// If not testing in Sandbox, remove these four lines and instead use a publicToken obtained from Link
	sandboxPublicTokenResp, err := p.GetPlublicTokenForSandboxAcct(ctx)
	if err != nil {
		return "", err
	}

	publicToken := sandboxPublicTokenResp.GetPublicToken()

	// Exchange the publicToken for an accessToken
	exchangePublicTokenResp, _, err := p.Client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()
	if err != nil {
		return "", err
	}

	return exchangePublicTokenResp.GetAccessToken(), nil
}

func (p *PlaidWrapper) GetPlublicTokenForSandboxAcct(ctx context.Context) (plaid.SandboxPublicTokenCreateResponse, error) {
	sandboxInstitution := "ins_109508"
	testProducts := []plaid.Products{plaid.PRODUCTS_AUTH}
	sandboxPublicTokenResp, _, err := p.Client.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		*plaid.NewSandboxPublicTokenCreateRequest(
			sandboxInstitution,
			testProducts,
		),
	).Execute()
	return sandboxPublicTokenResp, err
}
