package plaidhandler

import (
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"
)

var (
	PLAID_CLIENT_ID_TEST_UTILS  = "61eb5d49ea3b4700127560d1"
	PLAID_SECRET_KEY_TEST_UTILS = "465686056e8fd1b87db3d993d096d8"
	NEW_RELIC_KEY_TEST_UTILS    = "62fd721c712d5863a4e75b8f547b7c1ea884NRAL"
	SERVICE_TEST_UTILS          = "test"
	TEST_ACCESSTOKEN            = ""
	plaidTestClient             *PlaidWrapper
)

func NewMockPlaidClient() *plaid.APIClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID_TEST_UTILS)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET_KEY_TEST_UTILS)
	configuration.UseEnvironment(plaid.Sandbox)
	return plaid.NewAPIClient(configuration)
}

// NewNewRelicClient configures the new relic sdk with metadata specific to this service
func newNewRelicClient(logger *zap.Logger) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(SERVICE_TEST_UTILS),
		newrelic.ConfigLicense(NEW_RELIC_KEY_TEST_UTILS),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = true
		},
		// Use nrzap to register the logger with the agent:
		nrzap.ConfigLogger(logger.Named("newrelic")),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(true),
	)
}

func newLogger() *zap.Logger {
	return zap.L()
}

func GetPlaidWrapperForTest() (*PlaidWrapper, error) {
	l := newLogger()
	c, err := newNewRelicClient(l)
	if err != nil {
		return nil, err
	}

	return &PlaidWrapper{
		Client:         NewMockPlaidClient(),
		NewRelicClient: c,
		Logger:         l,
	}, nil
}
