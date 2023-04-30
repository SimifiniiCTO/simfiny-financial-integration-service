package plaidhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
)

type PlaidWrapper struct {
	client             *plaid.APIClient
	InstrumentationSdk *instrumentation.ServiceTelemetry
	Logger             *zap.Logger
	Environment        plaid.Environment
	ClientID           string
	SecretKey          string
	// OAuthDomain is used to specify the domain name that the user will be brought to upon returning to simfiny after
	// authenticating to a bank that requires OAuth. This will typically be a UI domain name and should not include a
	// protocol or a path. The protocol is auto inserted as `https` as it is the only protocol supported. The path is
	// currently hard coded until a need for different paths arises?
	OAuthDomain     string
	WebhooksDomain  string
	WebhooksEnabled bool
	EnabledProducts []plaid.Products
}

type PlaidWrapperImpl interface {
	// CreateLinkToken creates a link token for the given options
	CreateLinkToken(ctx context.Context, options *LinkTokenOptions) (LinkToken, error)
	// ExchangePublicToken exchanges a public token for an access token
	ExchangePublicToken(ctx context.Context, publicToken string) (*ItemToken, error)
	// GetWebhookVerificationKey retrieves the webhook verification key from Plaid
	GetWebhookVerificationKey(ctx context.Context, keyId string) (*WebhookVerificationKey, error)
	// GetInstitution retrieves the institution from Plaid
	GetInstitution(ctx context.Context, institutionId string) (*plaid.Institution, error)
	// Sync Plaid synchronizes the given item with Plaid
	Sync(ctx context.Context, cursor, accessToken *string) (*SyncResult, error)
	// GetAccounts returns all accounts for the given access token
	GetAccounts(ctx context.Context, accessToken string, userId uint64) ([]*schema.BankAccount, error)
	// GetAllTransactions returns all transactions for the given access token
	GetAllTransactions(ctx context.Context, accessToken string, start time.Time, end time.Time, accountIds []string) ([]*schema.Transaction, error)
	// GetTransactions returns transactions for the given access token
	GetTransactions(ctx context.Context, accessToken string, start, end time.Time, count, offset int32, bankAccountIds []string) ([]*schema.Transaction, error)
	// GetWebhooksURL returns the webhook url for the given webhook type
	GetWebhooksURL() string
	// GetRedirectURI returns the redirect uri for the given webhook type
	GetRedirectURI() string
	// GetInvestmentAccount returns the investment account for the given access token
	GetInvestmentAccount(ctx context.Context, userID uint64, accessToken string) ([]*schema.InvestmentAccount, error)
	// DeleteItem deletes the item for the given access token
	DeleteItem(ctx context.Context, accessToken *string) error
	// GetAccessTokenForSandboxAcct returns the access token for the sandbox account
	getAccessTokenForSandboxAcct() (string, error)
	// GetPlublicTokenForSandboxAcct returns the public token for the sandbox account
	getPlublicTokenForSandboxAcct(ctx context.Context) (plaid.SandboxPublicTokenCreateResponse, error)
}

var _ PlaidWrapperImpl = &PlaidWrapper{}

func (p *PlaidWrapper) GetWebhookVerificationKey(ctx context.Context, keyId string) (*WebhookVerificationKey, error) {
	request := p.client.PlaidApi.
		WebhookVerificationKeyGet(ctx).
		WebhookVerificationKeyGetRequest(plaid.WebhookVerificationKeyGetRequest{
			KeyId: keyId,
		})

	result, _, err := request.Execute()
	if err != nil {
		p.Logger.Error("failed to retrieve webhook verification key from Plaid", zap.Error(err))
		return nil, err
	}

	webhook, err := NewWebhookVerificationKeyFromPlaid(result.Key)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (p *PlaidWrapper) GetInstitution(ctx context.Context, institutionId string) (*plaid.Institution, error) {
	request := p.client.PlaidApi.
		InstitutionsGetById(ctx).
		InstitutionsGetByIdRequest(plaid.InstitutionsGetByIdRequest{
			InstitutionId: institutionId,
			CountryCodes:  PlaidCountries,
			Options: &plaid.InstitutionsGetByIdRequestOptions{
				IncludeOptionalMetadata:          pointer.BoolP(true),
				IncludeStatus:                    pointer.BoolP(true),
				IncludeAuthMetadata:              pointer.BoolP(false),
				IncludePaymentInitiationMetadata: pointer.BoolP(false),
			},
		})

	result, _, err := request.Execute()
	if err != nil {
		p.Logger.Error("failed to retrieve institution from Plaid", zap.Error(err))
		return nil, err
	}

	return &result.Institution, nil
}

func (p *PlaidWrapper) GetWebhooksURL() string {
	return fmt.Sprintf("https://%s/plaid/webhook", p.WebhooksDomain)
}

func (p *PlaidWrapper) GetRedirectURI() string {
	return fmt.Sprintf("https://%s/plaid/oauth-return", p.OAuthDomain)
}

func (p *PlaidWrapper) ToJSON() string {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(b)
}
