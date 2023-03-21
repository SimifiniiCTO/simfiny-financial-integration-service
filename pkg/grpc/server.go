package grpc

import (
	"errors"
	"time"

	"github.com/plaid/plaid-go/plaid"
	"github.com/stripe/stripe-go/v74/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	telemetry "github.com/SimifiniiCTO/core/core-telemetry"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
)

// Server is the grpc server object
type Server struct {
	proto.UnimplementedFinancialServiceServer

	// proto.UnimplementedFinancialServiceServer
	logger          *zap.Logger
	config          *Config
	instrumentation *instrumentation.ServiceTelemetry
	conn            *database.Db
	plaidClient     *plaidhandler.PlaidWrapper
	MetricsEngine   *telemetry.MetricsEngine
	stripeClient    *client.API
	kms             secrets.KeyManagement
}

// Config is the config for the grpc server initialization
type Config struct {
	Port                    int              `mapstructure:"grpc-port"`
	GatewayPort             int              `mapstructure:"grpc-gateway-port"`
	ServiceName             string           `mapstructure:"grpc-service-name"`
	PlaidWebhookURI         string           `mapstructure:"plaid-webhook-url"`
	PlaidRedirectURI        string           `mapstructure:"plaid-redirect-url"`
	NewRelicLicense         string           `mapstructure:"newrelic-key"`
	Environment             string           `mapstructure:"env"`
	PlaidProducts           []plaid.Products `mapstructure:"plaid-products"`
	RpcTimeout              time.Duration    `mapstructure:"rpc-timeout"`
	StripeApiKey            string           `mapstructure:"stripe-api-key"`
	PlaidClientID           string           `mapstructure:"plaid-client-id"`
	PlaidSecretKey          string           `mapstructure:"plaid-secret-key"`
	PlaidEnv                string           `mapstructure:"plaid-env"`
	PlaidOauthDomain        string           `mapstructure:"plaid-oauth-domain"`
	PlaidWebhooksEnabled    bool             `mapstructure:"plaid-webhooks-enabled"`
	PlaidWebhookOauthDomain string           `mapstructure:"plaid-webhook-oauth-domain"`
	AwsKeyID                string           `mapstructure:"aws-key-id"`
	AwsRegion               string           `mapstructure:"aws-region"`
	AwsSecretKey            string           `mapstructure:"aws-secret-key"`
	MaxPlaidLinks           int              `mapstructure:"max-plaid-links"`
	BillingEnabled          bool             `mapstructure:"billing-enabled"`
}

var _ proto.FinancialServiceServer = (*Server)(nil)

// Params is the params for the grpc server initialization
type Params struct {
	Config          *Config
	Logger          *zap.Logger
	Instrumentation *instrumentation.ServiceTelemetry
	Db              *database.Db
	PlaidClient     *plaid.APIClient
}

// RegisterGrpcServer registers the grpc server object
func (server *Server) RegisterGrpcServer(srv *grpc.Server) {
	proto.RegisterFinancialServiceServer(srv, server)
}

// Validate validates the params
func (p *Params) Validate() error {
	if p.Config == nil {
		return errors.New("config is nil")
	}

	if p.Logger == nil {
		return errors.New("logger is nil")
	}

	if p.Db == nil {
		return errors.New("db is nil")
	}

	if p.PlaidClient == nil {
		return errors.New("plaid client is nil")
	}

	return nil
}

// NewServer creates a new grpc server
func NewServer(param *Params) (*Server, error) {
	if param.Validate() != nil {
		return nil, errors.New("invalid param")
	}

	// configure aws kms
	keyManagement, err := secrets.NewAWSKMS(secrets.AWSKMSConfig{
		Region:    param.Config.AwsRegion,
		KeyID:     param.Config.AwsKeyID,
		SecretKey: param.Config.AwsSecretKey,
	})
	if err != nil {
		return nil, err
	}

	opts := []plaidhandler.Option{
		plaidhandler.WithEnvironment(plaid.Environment(param.Config.PlaidEnv)),
		plaidhandler.WithClientID(param.Config.PlaidClientID),
		plaidhandler.WithSecretKey(param.Config.PlaidSecretKey),
		plaidhandler.WithInstrumentation(param.Instrumentation),
		plaidhandler.WithLogger(param.Logger),
		plaidhandler.WithOauthDomain(param.Config.PlaidOauthDomain),
		plaidhandler.WithWebhooksDomain(param.Config.PlaidWebhookOauthDomain),
		plaidhandler.WithWebhooksEnabled(param.Config.PlaidWebhooksEnabled),
	}

	// declare plaid wrapper
	handler, err := plaidhandler.New(opts...)
	if err != nil {
		return nil, err
	}

	sc := &client.API{}
	sc.Init(param.Config.StripeApiKey, nil)

	return &Server{
		logger:          param.Logger,
		config:          param.Config,
		conn:            param.Db,
		plaidClient:     handler,
		instrumentation: param.Instrumentation,
		stripeClient:    sc,
		kms:             keyManagement,
	}, nil
}
