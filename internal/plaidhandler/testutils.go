package plaidhandler

import (
	"github.com/plaid/plaid-go/v14/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
)

var (
	PLAID_CLIENT_ID_TEST_UTILS  = "61eb5d49ea3b4700127560d1"
	PLAID_SECRET_KEY_TEST_UTILS = "465686056e8fd1b87db3d993d096d8"
	NEW_RELIC_KEY_TEST_UTILS    = "62fd721c712d5863a4e75b8f547b7c1ea884NRAL"
	SERVICE_TEST_UTILS          = "test"
	TEST_ACCESSTOKEN            = ""
	plaidTestClient             *PlaidWrapper
	PLAID_WEBHOOKS_DOMAIN       = "d30b-50-35-101-189.ngrok-free.app"
)

// NewMockPlaidClient returns a mock Plaid client for testing
func NewMockPlaidClient() *plaid.APIClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID_TEST_UTILS)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET_KEY_TEST_UTILS)
	configuration.UseEnvironment(plaid.Sandbox)
	return plaid.NewAPIClient(configuration)
}

// GetPlaidWrapperForTest returns a PlaidWrapper for testing
func GetPlaidWrapperForTest() (*PlaidWrapper, error) {
	return &PlaidWrapper{
		client:             NewMockPlaidClient(),
		InstrumentationSdk: &instrumentation.Client{},
		Logger:             zap.L(),
		Environment:        plaid.Sandbox,
		ClientID:           PLAID_CLIENT_ID_TEST_UTILS,
		SecretKey:          PLAID_SECRET_KEY_TEST_UTILS,
		OAuthDomain:        "",
		WebhooksDomain:     PLAID_WEBHOOKS_DOMAIN,
		WebhooksEnabled:    false,
	}, nil
}
