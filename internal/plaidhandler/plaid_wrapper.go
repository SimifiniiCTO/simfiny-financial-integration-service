package plaidhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/plaid/plaid-go/v14/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// The PlaidWrapper type is a struct that contains various fields related to the Plaid API client,
// including the client ID, secret key, enabled products, and webhooks settings.
// @property client - The Plaid API client used to make requests to the Plaid API.
// @property InstrumentationSdk - InstrumentationSdk is a client for Plaid's instrumentation API, which
// allows for monitoring and tracking of API usage and performance. It is used to collect metrics and
// traces for Plaid API requests and responses.
// @property Logger - The Logger property is a pointer to a zap.Logger object, which is a logging
// library used for structured logging in Go. It is used to log information about Plaid API requests
// and responses, as well as any errors or warnings that occur during the execution of the PlaidWrapper
// methods.
// @property Environment - The Plaid environment (e.g. sandbox, development, production) that the
// PlaidWrapper is configured to use.
// @property {string} ClientID - The Plaid client ID, which is used to authenticate API requests.
// @property {string} SecretKey - The Plaid secret key is a confidential API key used to authenticate
// requests made to the Plaid API. It is used to verify the identity of the client making the request
// and to ensure that the request is authorized to access the requested resources.
// @property {string} OAuthDomain - OAuthDomain is a string property used to specify the domain name
// that the user will be brought to upon returning to simfiny after authenticating to a bank that
// requires OAuth. This is typically a UI domain name and should not include a protocol or a path. The
// protocol is automatically inserted as `https
// @property {string} WebhooksDomain - WebhooksDomain is a string property that is used to specify the
// domain name that Plaid will send webhooks to. This is typically a UI domain name and should not
// include a protocol or a path. The protocol is automatically inserted as `https` as it is the only
// protocol supported.
// @property {bool} WebhooksEnabled - WebhooksEnabled is a boolean property that specifies whether
// webhooks are enabled or not for the Plaid API client. If webhooks are enabled, the WebhooksDomain
// property must also be specified.
// @property {[]plaid.Products} EnabledProducts - EnabledProducts is a slice of plaid.Products that
// specifies the products that are enabled for the Plaid API client. These products include
// AccountInfo, Auth, Balance, Identity, Income, Transactions, Assets, Liabilities, Investments, and
// PaymentInitiation.
type PlaidWrapper struct {
	// Plaid API client
	client *plaid.APIClient
	// InstrumentationSdk is the instrumentation client
	InstrumentationSdk *instrumentation.Client
	// Logger is the logger
	Logger *zap.Logger
	// Environment is the Plaid environment
	Environment plaid.Environment
	// ClientID is the Plaid client ID
	ClientID string
	// SecretKey is the Plaid secret key
	SecretKey string
	// OAuthDomain is used to specify the domain name that the user will be brought to upon returning to simfiny after
	// authenticating to a bank that requires OAuth. This will typically be a UI domain name and should not include a
	// protocol or a path. The protocol is auto inserted as `https` as it is the only protocol supported. The path is
	// currently hard coded until a need for different paths arises?
	OAuthDomain string
	// WebhooksDomain is used to specify the domain name that Plaid will send webhooks to. This will typically be a
	// UI domain name and should not include a protocol or a path. The protocol is auto inserted as `https` as it is the
	// only protocol supported.
	WebhooksDomain string
	// WebhooksEnabled is used to specify whether webhooks are enabled or not. If webhooks are enabled, the
	// WebhooksDomain must be specified.
	WebhooksEnabled bool
	// EnabledProducts is used to specify the products that are enabled for the Plaid API client
	EnabledProducts []plaid.Products

	/**
		In order to require a secondary product when possible, without excluding users who don't have access
		to that product from linking their account, you can specify the secondary product in the required_if_supported_products
		array when calling /link/token/create. (It is possible to specify multiple products in this field as well, as long as
		at least one always-required product is specified in the regular products array.)

		When a product is listed in required_if_supported_products, Plaid will require that product if possible. If the institution
		the end user is attempting to link supports the secondary product, and the user grants Plaid access to an account at the
		institution that is compatible with it, Plaid will treat the product as though it was listed in the products array.
		If the institution does not support the product, if the user doesn't have accounts compatible with it, or if the user
		has compatible accounts but doesn't grant access to them, Plaid will treat the product as though it had not been specified
		in the /link/token/create call. (To determine whether a product was successfully added to the Item, you can
		all /item/get after obtaining your access_token and look at the value of the products array.)
	**/
	RequiredProductsIfSupported []plaid.Products
}

// The PlaidWrapperImpl interface defines methods for interacting with the Plaid API, including
// creating link tokens, exchanging public tokens for access tokens, retrieving webhook verification
// keys and institutions, synchronizing items, getting accounts and transactions, and deleting items.
// @property CreateLinkToken - Creates a link token for the given options. A link token is a
// short-lived token that is used to initialize the Plaid Link flow.
// @property ExchangePublicToken - This method exchanges a public token for an access token. A public
// token is a short-lived token that can be safely shared with a client-side application. An access
// token, on the other hand, is a long-lived token that provides access to a user's financial data.
// This method is used to obtain
// @property GetWebhookVerificationKey - GetWebhookVerificationKey is a method that retrieves the
// webhook verification key from Plaid. This key is used to verify the authenticity of webhook events
// sent by Plaid.
// @property GetInstitution - This method retrieves the institution details from Plaid for a given
// institution ID.
// @property Sync - Sync is a method that synchronizes the given item with Plaid. It takes in a
// context, a cursor, and an access token as parameters and returns a SyncResult and an error. The
// SyncResult contains information about the synchronization process, such as the number of new
// transactions and the number of updated
// @property GetAccounts - GetAccounts is a method that returns all accounts associated with a given
// access token and user ID. It takes in a context.Context object and the access token as a string, and
// returns a slice of schema.BankAccount objects and an error.
// @property GetAllTransactions - GetAllTransactions is a method that returns all transactions for the
// given access token within a specified time range and for specified account IDs. It takes in a
// context, access token, start and end time, and an array of account IDs as parameters and returns an
// array of schema.Transaction and an error.
// @property GetTransactions - GetTransactions is a method that retrieves transactions for the given
// access token within a specified time range and with optional filters for account IDs. It returns a
// slice of schema.Transaction objects and an error if the request fails. The method takes in the
// following parameters:
// @property {string} GetWebhooksURL - GetWebhooksURL is a method that returns the webhook URL for the
// given webhook type.
// @property {string} GetRedirectURI - GetRedirectURI is a method that returns the redirect URI for the
// given webhook type.
// @property GetInvestmentAccount - GetInvestmentAccount is a method that retrieves investment accounts
// associated with a given Plaid access token and user ID. It takes in a context, user ID, and access
// token as parameters and returns a slice of InvestmentAccount structs and an error.
// @property {error} DeleteItem - DeleteItem is a method that deletes the item associated with the
// given access token from Plaid.
// @property getAccessTokenForSandboxAcct - A method that returns the access token for a sandbox
// account.
// @property getPlublicTokenForSandboxAcct - This is a method that retrieves a public token for a
// sandbox account in Plaid.
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
	GetAccessTokenForSandboxAcct() (*plaid.ItemPublicTokenExchangeResponse, error)
	// UpdateLinkToken returns the access token for the link account associated with the link account
	UpdateLinkToken(ctx context.Context, options *LinkTokenOptions, accessToken *string) (LinkToken, error)
	// GetPlublicTokenForSandboxAcct returns the public token for the sandbox account
	getPlublicTokenForSandboxAcct(ctx context.Context) (plaid.SandboxPublicTokenCreateResponse, error)
}

// The above code is declaring a variable `_` of type `PlaidWrapperImpl` and initializing it with a new
// instance of `PlaidWrapper`. The `&` symbol is used to get the memory address of the newly created
// `PlaidWrapper` instance. This code is written in Go programming language.
var _ PlaidWrapperImpl = &PlaidWrapper{}

// GetWebhookVerificationKey retrieves the webhook verification key from Plaid
func (p *PlaidWrapper) GetWebhookVerificationKey(ctx context.Context, keyId string) (*WebhookVerificationKey, error) {
	p.Logger.Info("data", zap.Any("details", p))

	request := p.client.PlaidApi.
		WebhookVerificationKeyGet(ctx).
		WebhookVerificationKeyGetRequest(plaid.WebhookVerificationKeyGetRequest{
			ClientId: &p.ClientID,
			Secret:   &p.SecretKey,
			KeyId:    keyId,
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

// GetInstitution retrieves the institution from Plaid
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

// GetWebhooksURL returns the webhook url for the given webhook type
func (p *PlaidWrapper) GetWebhooksURL() string {
	return fmt.Sprintf("https://%s/plaid/webhook", p.WebhooksDomain)
}

// GetRedirectURI returns the redirect uri for the given webhook type
func (p *PlaidWrapper) GetRedirectURI() string {
	return fmt.Sprintf("https://%s/plaid/oauth-return", p.OAuthDomain)
}

// ToJSON returns the JSON representation of the PlaidWrapper
func (p *PlaidWrapper) ToJSON() string {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(b)
}

// EnabledProductsToString returns the string representation of the enabled products
func (p *PlaidWrapper) EnabledProductsToString() []string {
	results := make([]string, 0, len(p.EnabledProducts))
	for _, element := range p.EnabledProducts {
		results = append(results, string(element))
	}

	return results
}
