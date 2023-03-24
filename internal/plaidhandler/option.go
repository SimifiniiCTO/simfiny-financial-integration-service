package plaidhandler

import (
	"errors"

	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
)

type Option func(*PlaidWrapper)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

// WithInstrumentation is a functional option to set the instrumentation service
func WithInstrumentation(instrumentation *instrumentation.ServiceTelemetry) Option {
	return func(p *PlaidWrapper) {
		p.InstrumentationSdk = instrumentation
	}
}

// WithLogger is a functional option to set the logger
func WithLogger(logger *zap.Logger) Option {
	return func(p *PlaidWrapper) {
		p.Logger = logger
	}
}

// WithEnvironment is a functional option to set the environment
func WithEnvironment(environment plaid.Environment) Option {
	return func(p *PlaidWrapper) {
		p.Environment = environment
	}
}

// WithClientID is a functional option to set the client id
func WithClientID(clientID string) Option {
	return func(p *PlaidWrapper) {
		p.ClientID = clientID
	}
}

// WithSecretKey is a functional option to set the secret key
func WithSecretKey(secretKey string) Option {
	return func(p *PlaidWrapper) {
		p.SecretKey = secretKey
	}
}

// WithOAuthDomain is a functional option to set the oauth domain
func WithOauthDomain(domain string) Option {
	return func(p *PlaidWrapper) {
		p.OAuthDomain = domain
	}
}

// WithWebhooksDomain is a functional option to set the webhooks domain
func WithWebhooksDomain(domain string) Option {
	return func(p *PlaidWrapper) {
		p.WebhooksDomain = domain
	}
}

// WithWebhooksEnabled is a functional option to set the webhooks enabled
func WithWebhooksEnabled(enabled bool) Option {
	return func(p *PlaidWrapper) {
		p.WebhooksEnabled = enabled
	}
}

// New creates a new PlaidWrapper instance
func New(opts ...Option) (*PlaidWrapper, error) {
	plaidWrapper := &PlaidWrapper{}

	for _, opt := range opts {
		opt(plaidWrapper)
	}

	if plaidWrapper.Logger == nil ||
		plaidWrapper.Environment == "" ||
		plaidWrapper.ClientID == "" ||
		plaidWrapper.SecretKey == "" ||
		plaidWrapper.InstrumentationSdk == nil {
		return nil, errors.New("invalid input arguments. check input parameters")
	}

	// configure plaid client
	plaidWrapper.configureClient()

	return plaidWrapper, nil
}
