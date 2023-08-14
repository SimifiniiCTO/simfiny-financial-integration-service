package plaidhandler

import (
	"fmt"

	"github.com/hashicorp/go-set"
	"github.com/plaid/plaid-go/v14/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
)

type Option func(*PlaidWrapper)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

// WithInstrumentation is a functional option to set the instrumentation service
func WithInstrumentation(instrumentation *instrumentation.Client) Option {
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

// WithProducts is a functional option to set the plaid products
func WithProducts(products []string) Option {
	var plaidProducts = map[string]plaid.Products{
		string(plaid.PRODUCTS_AUTH):                   plaid.PRODUCTS_AUTH,
		string(plaid.PRODUCTS_ASSETS):                 plaid.PRODUCTS_ASSETS,
		string(plaid.PRODUCTS_INVESTMENTS):            plaid.PRODUCTS_INVESTMENTS,
		string(plaid.PRODUCTS_LIABILITIES):            plaid.PRODUCTS_LIABILITIES,
		string(plaid.PRODUCTS_BALANCE):                plaid.PRODUCTS_BALANCE,
		string(plaid.PRODUCTS_IDENTITY):               plaid.PRODUCTS_IDENTITY,
		string(plaid.PRODUCTS_PAYMENT_INITIATION):     plaid.PRODUCTS_PAYMENT_INITIATION,
		string(plaid.PRODUCTS_TRANSACTIONS):           plaid.PRODUCTS_TRANSACTIONS,
		string(plaid.PRODUCTS_CREDIT_DETAILS):         plaid.PRODUCTS_CREDIT_DETAILS,
		string(plaid.PRODUCTS_INCOME):                 plaid.PRODUCTS_INCOME,
		string(plaid.PRODUCTS_INCOME_VERIFICATION):    plaid.PRODUCTS_INCOME_VERIFICATION,
		string(plaid.PRODUCTS_DEPOSIT_SWITCH):         plaid.PRODUCTS_DEPOSIT_SWITCH,
		string(plaid.PRODUCTS_STANDING_ORDERS):        plaid.PRODUCTS_STANDING_ORDERS,
		string(plaid.PRODUCTS_TRANSFER):               plaid.PRODUCTS_TRANSFER,
		string(plaid.PRODUCTS_RECURRING_TRANSACTIONS): plaid.PRODUCTS_RECURRING_TRANSACTIONS,
	}
	return func(p *PlaidWrapper) {
		// ref: https://plaid.com/docs/link/initializing-products/#required-if-supported-products

		// iterate over set of strings and convert to plaid product
		enabledProducts := set.New[plaid.Products](0)
		requiredProductsIfSupported := set.New[plaid.Products](0)
		requiredProductsIfSupported.Insert(plaid.PRODUCTS_INVESTMENTS)
		requiredProductsIfSupported.Insert(plaid.PRODUCTS_LIABILITIES)

		for _, element := range products {
			if value, ok := plaidProducts[element]; ok {
				enabledProducts.Insert(value)
			}
		}

		requiredProductsIfSupported.ForEach(func(element plaid.Products) bool {
			if enabledProducts.Contains(element) {
				fmt.Printf("removing %s from enabled products\n", element)
				enabledProducts.Remove(element)
				return true
			}

			return false
		})

		fmt.Println("enabled products: ", enabledProducts.Slice())
		fmt.Println("required if enabled products: ", requiredProductsIfSupported.Slice())

		p.RequiredProductsIfSupported = requiredProductsIfSupported.Slice()
		p.EnabledProducts = enabledProducts.Slice()
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
		plaidWrapper.OAuthDomain == "" ||
		plaidWrapper.WebhooksDomain == "" ||
		plaidWrapper.InstrumentationSdk == nil ||
		len(plaidWrapper.EnabledProducts) == 0 {
		return nil, fmt.Errorf("invalid input arguments. check input parameters. params: %s", plaidWrapper.ToJSON())
	}

	// configure plaid client
	plaidWrapper.configureClient()

	return plaidWrapper, nil
}
