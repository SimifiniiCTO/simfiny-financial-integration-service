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
)

// Server is the grpc server object
type Server struct {
	proto.UnimplementedFinancialServiceServer

	// proto.UnimplementedFinancialServiceServer
	logger               *zap.Logger
	config               *Config
	instrumentation      *instrumentation.ServiceTelemetry
	conn                 *database.Db
	plaidClient          *plaid.APIClient
	plaidInternalHandler *plaidhandler.PlaidWrapper
	MetricsEngine        *telemetry.MetricsEngine
	stripeClient         *client.API
}

// Config is the config for the grpc server initialization
type Config struct {
	Port             int              `mapstructure:"grpc-port"`
	GatewayPort      int              `mapstructure:"grpc-gateway-port"`
	ServiceName      string           `mapstructure:"grpc-service-name"`
	PlaidWebhookURI  string           `mapstructure:"plaid-webhook-url"`
	PlaidRedirectURI string           `mapstructure:"plaid-redirect-url"`
	NewRelicLicense  string           `mapstructure:"newrelic-key"`
	Environment      string           `mapstructure:"env"`
	PlaidProducts    []plaid.Products `mapstructure:"plaid-products"`
	RpcTimeout       time.Duration    `mapstructure:"rpc-timeout"`
	StripeApiKey     string           `mapstructure:"stripe-api-key"`
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

	// declare plaid wrapper
	pHandler, err := plaidhandler.NewPlaidWrapper(
		param.PlaidClient,
		param.Instrumentation,
		param.Logger)
	if err != nil {
		return nil, err
	}

	sc := &client.API{}
	sc.Init(param.Config.StripeApiKey, nil)

	return &Server{
		logger:               param.Logger,
		config:               param.Config,
		conn:                 param.Db,
		plaidClient:          param.PlaidClient,
		plaidInternalHandler: pHandler,
		instrumentation:      param.Instrumentation,
		stripeClient:         sc,
	}, nil
}
