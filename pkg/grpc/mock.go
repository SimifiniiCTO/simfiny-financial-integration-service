package grpc

import (
	"context"
	"log"
	"net"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type MockDialOption func(context.Context, string) (net.Conn, error)

// dialer creates an in memory full duplex connection
func dialer() func() MockDialOption {
	return func() MockDialOption {
		listener := bufconn.Listen(1024 * 1024)

		server := grpc.NewServer()
		s := NewMockServer()
		proto.RegisterFinancialIntegrationServiceServer(server, s)

		go func() {
			if err := server.Serve(listener); err != nil {
				log.Fatal(err)
			}
		}()

		return func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}
	}
}

// MockGRPCService creates and returns a mock grpc service connection
func MockGRPCService(ctx context.Context) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// NewMockServer creates a new mock server instance
func NewMockServer() *Server {
	config := &Config{
		Port:            9897,
		ServiceName:     "FinancialIntegrationService",
		Environment:     "dev",
		NewRelicLicense: "62fd721c712d5863a4e75b8f547b7c1ea884NRAL",
	}

	dbConfigs := &database.DefaultConnInitializationParams
	telemetry, _ := newrelic.NewApplication(
		newrelic.ConfigAppName(config.ServiceName),
		newrelic.ConfigLicense(config.NewRelicLicense),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = true
		},
		// Use nrzap to register the logger with the agent:
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(false),
	)
	dbConfigs.Telemetry = telemetry

	dbConn, err := database.New(context.Background(), dbConfigs)
	if err != nil {
		log.Fatal(err)
	}

	plaidHandler, err := plaidhandler.GetPlaidWrapperForTest()
	if err != nil {
		log.Fatal(err)
	}

	client := plaidhandler.NewMockPlaidClient()

	srv := &Server{
		logger:                        zap.S().Desugar(),
		config:                        config,
		telemetry:                     telemetry,
		grpcCompatibleInstrumentation: nil,
		conn:                          dbConn,
		plaidClient:                   client,
		plaidInternalHandler:          plaidHandler,
	}

	return srv
}
