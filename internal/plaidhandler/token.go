package plaidhandler

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"
)

func (p *PlaidWrapper) TriggerWebhookForTest(ctx context.Context, clientId, secret, accesstoken, webhookCode string) error {
	req := plaid.NewSandboxItemFireWebhookRequest(accesstoken, webhookCode)
	_, _, err := p.client.PlaidApi.SandboxItemFireWebhook(ctx).SandboxItemFireWebhookRequest(*req).Execute()
	return err
}

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

	// if we are in sandbox mode, we need to create a public token for the sandbox account
	// based on the env config value ... we decipher wether to use the sandbox or production plaid client
	if p.Environment == plaid.Sandbox {
		publicToken, err := p.getPlublicTokenForSandboxAcct(ctx)
		if err != nil {
			p.Logger.Error("failed to create link token with Plaid", zap.Error(err))
			return nil, err
		}

		return &PlaidLinkToken{
			LinkToken: publicToken.GetPublicToken(),
		}, nil
	}

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId:             options.ClientUserID,
		LegalName:                &options.LegalName,
		PhoneNumber:              options.PhoneNumber,
		PhoneNumberVerifiedTime:  *plaid.NewNullableTime(options.PhoneNumberVerifiedTime),
		EmailAddress:             &options.EmailAddress,
		EmailAddressVerifiedTime: *plaid.NewNullableTime(options.EmailAddressVerifiedTime),
	}

	request := plaid.NewLinkTokenCreateRequest(
		PlaidClientName,
		PlaidLanguage,
		PlaidCountries,
		user,
	)

	request.SetProducts(p.EnabledProducts)
	request.SetLinkCustomizationName(PlaidClientName)
	request.SetWebhook(*webhooksUrl)
	request.SetRedirectUri(*redirectUri)

	result, _, err := p.client.PlaidApi.
		LinkTokenCreate(ctx).
		LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &PlaidLinkToken{
		LinkToken: result.LinkToken,
		Expires:   result.Expiration,
	}, nil
}

func (p *PlaidWrapper) ExchangePublicToken(ctx context.Context, publicToken string) (*ItemToken, error) {
	if p.Environment == plaid.Sandbox {
		res, err := p.getAccessTokenForSandboxAcct()
		if err != nil {
			p.Logger.Error("failed to create link token with Plaid", zap.Error(err))
			return nil, err
		}

		return &ItemToken{
			AccessToken: res.AccessToken,
			ItemId:      res.ItemId,
		}, nil
	}

	request := p.client.PlaidApi.
		ItemPublicTokenExchange(ctx).
		ItemPublicTokenExchangeRequest(plaid.ItemPublicTokenExchangeRequest{
			PublicToken: publicToken,
		})

	result, resp, err := request.Execute()
	if err != nil {
		// ensure we close the response body handle
		defer resp.Body.Close()

		// Read response body
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		p.Logger.Error("failed to exchange public token with Plaid", zap.Error(err), zap.Any("details", string(bodyBytes)))
		return nil, err
	}

	token, err := NewItemTokenFromPlaid(result)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (p *PlaidWrapper) getAccessTokenForSandboxAcct() (*plaid.ItemPublicTokenExchangeResponse, error) {
	ctx := context.Background()

	// If not testing in Sandbox, remove these four lines and instead use a publicToken obtained from Link
	result, err := p.getPlublicTokenForSandboxAcct(ctx)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &resp, nil
}

func (p *PlaidWrapper) getPlublicTokenForSandboxAcct(ctx context.Context) (plaid.SandboxPublicTokenCreateResponse, error) {
	const (
		sandboxInstitution = "ins_109508"
	)

	testProducts := []plaid.Products{
		plaid.PRODUCTS_AUTH,
		plaid.PRODUCTS_INVESTMENTS,
		plaid.PRODUCTS_LIABILITIES,
		plaid.PRODUCTS_TRANSACTIONS,
	}

	req := plaid.NewSandboxPublicTokenCreateRequest(
		sandboxInstitution,
		testProducts,
	)

	webhookUrl := p.GetWebhooksURL()
	req.Options = &plaid.SandboxPublicTokenCreateRequestOptions{
		Webhook: &webhookUrl,
	}

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
	configuration.UseEnvironment(p.Environment)
	p.client = plaid.NewAPIClient(configuration)
}
