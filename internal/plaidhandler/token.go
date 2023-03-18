package plaidhandler

import (
	"context"
	"fmt"

	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
)

func (p *PlaidWrapper) CreateLinkToken(ctx context.Context, options *LinkTokenOptions) (LinkToken, error) {
	var redirectUri *string
	if p.OAuthDomain != "" {
		redirectUri = pointer.StringP(fmt.Sprintf("https://%s/plaid/oauth-return", p.OAuthDomain))
	}

	var webhooksUrl *string
	if p.WebhooksEnabled {
		if p.WebhooksDomain == "" {
			p.Logger.Error("BUG: Plaid webhook domain is not present but webhooks are enabled.")
		} else {
			webhooksUrl = pointer.StringP(p.GetWebhooksURL())
		}
	}

	// based on the env config value ... we decipher wether to use the sandbox or production plaid client
	request := p.client.PlaidApi.
		LinkTokenCreate(ctx).
		LinkTokenCreateRequest(plaid.LinkTokenCreateRequest{
			ClientName:   PlaidClientName,
			Language:     PlaidLanguage,
			CountryCodes: PlaidCountries,
			User: plaid.LinkTokenCreateRequestUser{
				ClientUserId:             options.ClientUserID,
				LegalName:                &options.LegalName,
				PhoneNumber:              options.PhoneNumber,
				PhoneNumberVerifiedTime:  options.PhoneNumberVerifiedTime,
				EmailAddress:             &options.EmailAddress,
				EmailAddressVerifiedTime: options.EmailAddressVerifiedTime,
				Ssn:                      nil,
				DateOfBirth:              nil,
			},
			Products:    &PlaidProducts,
			Webhook:     webhooksUrl,
			RedirectUri: redirectUri,
		})

	result, _, err := request.Execute()
	if err != nil {
		p.Logger.Error("failed to create link token with Plaid", zap.Error(err))
		return nil, err
	}

	return &PlaidLinkToken{
		LinkToken: result.LinkToken,
		Expires:   result.Expiration,
	}, nil
}

func (p *PlaidWrapper) ExchangePublicToken(ctx context.Context, publicToken string) (*ItemToken, error) {
	request := p.client.PlaidApi.
		ItemPublicTokenExchange(ctx).
		ItemPublicTokenExchangeRequest(plaid.ItemPublicTokenExchangeRequest{
			PublicToken: publicToken,
		})

	result, _, err := request.Execute()
	if err != nil {
		p.Logger.Error("failed to exchange public token with Plaid", zap.Error(err))
		return nil, err
	}

	token, err := NewItemTokenFromPlaid(result)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (p *PlaidWrapper) getAccessTokenForSandboxAcct() (string, error) {
	ctx := context.Background()

	// If not testing in Sandbox, remove these four lines and instead use a publicToken obtained from Link
	result, err := p.getPlublicTokenForSandboxAcct(ctx)
	if err != nil {
		return "", err
	}

	req := plaid.NewItemPublicTokenExchangeRequest(result.GetPublicToken())

	// Exchange the publicToken for an accessToken
	resp, _, err := p.
		client.
		PlaidApi.
		ItemPublicTokenExchange(ctx).
		ItemPublicTokenExchangeRequest(*req).
		Execute()

	if err != nil {
		return "", err
	}

	return resp.GetAccessToken(), nil
}

func (p *PlaidWrapper) getPlublicTokenForSandboxAcct(ctx context.Context) (plaid.SandboxPublicTokenCreateResponse, error) {
	const (
		sandboxInstitution = "ins_109508"
	)

	testProducts := []plaid.Products{
		plaid.PRODUCTS_AUTH,
		plaid.PRODUCTS_BALANCE,
		plaid.PRODUCTS_INVESTMENTS,
		plaid.PRODUCTS_LIABILITIES,
		plaid.PRODUCTS_TRANSACTIONS,
	}

	req := plaid.NewSandboxPublicTokenCreateRequest(
		sandboxInstitution,
		testProducts,
	)

	resp, _, err := p.
		client.
		PlaidApi.
		SandboxPublicTokenCreate(ctx).
		SandboxPublicTokenCreateRequest(*req).
		Execute()
	return resp, err
}

// configureClient configures the plaid client
func (p *PlaidWrapper) configureClient() {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", p.ClientID)
	configuration.AddDefaultHeader("PLAID-SECRET", p.SecretKey)
	configuration.UseEnvironment(environments[string(p.Environment)])
	p.client = plaid.NewAPIClient(configuration)
}
