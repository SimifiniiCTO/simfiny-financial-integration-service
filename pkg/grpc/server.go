package grpc

import (
	telemetry "github.com/SimifiniiCTO/core/core-telemetry"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	nrGrpcCompatibleVersion "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedFinancialIntegrationServiceServer
	logger                        *zap.Logger
	config                        *Config
	telemetry                     *newrelic.Application
	grpcCompatibleInstrumentation nrGrpcCompatibleVersion.Application
	conn                          *database.Db
	plaidClient                   *plaid.APIClient
	plaidInternalHandler          *plaidhandler.PlaidWrapper
	MetricsEngine                 *telemetry.MetricsEngine
}

type Config struct {
	Port             int    `mapstructure:"grpc-port"`
	GatewayPort      int    `mapstructure:"grpc-gateway-port"`
	ServiceName      string `mapstructure:"grpc-service-name"`
	PlaidWebhookURI  string `mapstructure:"plaid-webhook-url"`
	PlaidRedirectURI string `mapstructure:"plaid-redirect-url"`
	NewRelicLicense  string `mapstructure:"newrelic-key"`
	Environment      string `mapstructure:"env"`
}

func NewServer(config *Config, logger *zap.Logger, telemetry *newrelic.Application, db *database.Db, plaidClient *plaid.APIClient, engine *telemetry.MetricsEngine) (*Server, error) {
	pHandler, err := plaidhandler.NewPlaidWrapper(plaidClient, telemetry, logger)
	if err != nil {
		return nil, err
	}

	cfg := nrGrpcCompatibleVersion.NewConfig(config.ServiceName, config.NewRelicLicense)
	grpcCompatibleInstrmtAgent, err := nrGrpcCompatibleVersion.NewApplication(cfg)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		logger:                        logger,
		config:                        config,
		telemetry:                     telemetry,
		conn:                          db,
		plaidClient:                   plaidClient,
		plaidInternalHandler:          pHandler,
		grpcCompatibleInstrumentation: grpcCompatibleInstrmtAgent,
		MetricsEngine:                 engine,
	}

	return srv, nil
}

// RegisterGrpcServer registers the grpc server object
func (server *Server) RegisterGrpcServer(srv *grpc.Server) {
	proto.RegisterFinancialIntegrationServiceServer(srv, server)
}
