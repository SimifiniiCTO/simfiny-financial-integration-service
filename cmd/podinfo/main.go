package main

import (
	"context"
	"fmt"

	newrelic "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/plaid/plaid-go/plaid"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	core_database "github.com/SimifiniiCTO/core/core-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/env"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/api"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/grpc"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/signals"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/version"
)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

func main() {
	// reads in all environment variables
	env.ReadEnvVars()
	ctx := context.Background()

	// Create a new boot instance.
	boot := rkboot.NewBoot()
	// configure logging
	logger := rkentry.GlobalAppCtx.GetLoggerEntry("zap-logger").Logger
	logger.Info("starting service ....")

	logger.Info("successfully initialized newrelic sdk ....")
	instrumentation := configureServiceTelemetryInstance(logger)

	// configure plaid client
	plaidClient, err := configurePlaidSDK()
	if err != nil {
		logger.Panic(err.Error())
	}

	logger.Info("initializing database ....")
	db, err := configureDatabaseConn(ctx, logger, instrumentation)
	if err != nil {
		logger.Panic(err.Error())
	}

	logger.Info("successfully initialized database ....")

	conn, err := db.Conn.Engine.DB()
	if err != nil {
		logger.Panic(err.Error())
	}
	defer conn.Close()

	logger.Info("successfully initialized database ....")

	// load gRPC server config
	var grpcCfg grpc.Config
	if err := viper.Unmarshal(&grpcCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	// initialize gRPC server
	grpcSrv, err := grpc.NewServer(&grpc.Params{
		Config:      &grpcCfg,
		Logger:      logger,
		Db:          db,
		PlaidClient: plaidClient,
	})
	if err != nil {
		logger.Panic(err.Error())
	}

	grpcEntry := rkgrpc.GetGrpcEntry("financial-integration-service")
	grpcEntry.AddRegFuncGrpc(grpcSrv.RegisterGrpcServer)
	grpcEntry.AddRegFuncGw(proto.RegisterFinancialServiceHandlerFromEndpoint)

	// TODO: add grpc interceptor middleware to emit metrics on various gRPC calls
	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())

	// load HTTP server config
	var srvCfg api.Config
	if err := viper.Unmarshal(&srvCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	// log version and port
	logger.Info("Starting financial integration service",
		zap.String("version", viper.GetString("version")),
		zap.String("revision", viper.GetString("revision")),
		zap.String("port", srvCfg.Port),
	)

	// start HTTP server
	srv, _ := api.NewServer(&srvCfg, logger, instrumentation, db, plaidClient)
	stopCh := signals.SetupSignalHandler()
	srv.ListenAndServe(stopCh)
}

// configurNewrelicSDK configures the new relic sdk with metadata specific to this service
func configureNewrelicSDK(logger *zap.Logger) (newrelic.Application, error) {
	var newrelicLicenseKey = viper.GetString("newrelic-key")
	var serviceName = viper.GetString("grpc-service-name")

	if newrelicLicenseKey != "" {
		cfg := newrelic.NewConfig(serviceName, newrelicLicenseKey)
		cfg.ErrorCollector.CaptureEvents = true
		cfg.ErrorCollector.Enabled = true
		cfg.TransactionEvents.Enabled = true
		cfg.Enabled = true
		cfg.TransactionEvents.Enabled = true
		cfg.Attributes.Enabled = true
		cfg.BrowserMonitoring.Enabled = true
		cfg.TransactionTracer.Enabled = true
		cfg.SpanEvents.Enabled = true
		cfg.RuntimeSampler.Enabled = true
		cfg.DistributedTracer.Enabled = true
		cfg.AppName = serviceName
		cfg.BrowserMonitoring.Enabled = true
		cfg.CustomInsightsEvents.Enabled = true
		cfg.DatastoreTracer.InstanceReporting.Enabled = true
		cfg.DatastoreTracer.QueryParameters.Enabled = true
		cfg.DatastoreTracer.DatabaseNameReporting.Enabled = true
		cfg.Logger = nrzap.Transform(logger.Named("newrelic"))
		cfg.DistributedTracer.Enabled = true
		cfg.Enabled = true

		app, err := newrelic.NewApplication(cfg)
		if err != nil {
			return nil, err
		}

		return app, nil
	}

	return nil, fmt.Errorf("invalid input parameter. param: newrelicLicenseKey = %s", newrelicLicenseKey)
}

func configurePlaidSDK() (*plaid.APIClient, error) {
	var (
		plaidClientID  = viper.GetString("plaid-client-id")
		plaidSecretKey = viper.GetString("plaid-secret-key")
		plaidEnv       = viper.GetString("plaid-env")
	)

	if plaidClientID == "" || plaidSecretKey == "" {
		return nil, fmt.Errorf(
			"invalid input arguments. plaidClientId: %s, plaidSecretKey: %s",
			plaidClientID,
			plaidSecretKey)
	}

	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", plaidClientID)
	configuration.AddDefaultHeader("PLAID-SECRET", plaidSecretKey)
	configuration.UseEnvironment(environments[plaidEnv])
	return plaid.NewAPIClient(configuration), nil
}

// configureDatabaseConn configures the database connection
func configureDatabaseConn(ctx context.Context, logger *zap.Logger, instrumentation *instrumentation.ServiceTelemetry) (*database.Db, error) {
	host := viper.GetString("dbhost")
	port := viper.GetInt("dbport")
	user := viper.GetString("dbuser")
	password := viper.GetString("dbpassword")
	dbname := viper.GetString("dbname")
	sslMode := viper.GetString("dbsslmode")

	maxDBConnAttempts := viper.GetInt("max-db-conn-attempts")
	maxRetriesPerDBConnectionAttempt := viper.GetInt("max-db-conn-retries")
	maxDBRetryTimeout := viper.GetDuration("max-db-retry-timeout")
	maxDBSleepInterval := viper.GetDuration("max-db-retry-sleep-interval")
	dbQueryTimeout := viper.GetDuration("db-query-timeout")

	connectionString := database.ConfigureConnectionString(host, user, password, dbname, sslMode, port)
	logger.Info(fmt.Sprintf("%s", connectionString))

	logger.Info("establishing database connection")

	// establish db connections
	conn := core_database.NewDatabaseConn(
		&core_database.Parameters{
			QueryTimeout:              &dbQueryTimeout,
			MaxConnectionRetries:      &maxDBConnAttempts,
			MaxConnectionRetryTimeout: &maxDBRetryTimeout,
			RetrySleep:                &maxDBSleepInterval,
			ConnectionString:          &connectionString,
		})

	logger.Info("establishing database connection")

	queryOperator := dal.Use(conn.Engine)
	opts := []database.Option{
		database.WithConnection(conn),
		database.WithDatabaseMaxConnectionAttempts(maxDBConnAttempts),
		database.WithDatabaseMaxRetriesPerOperation(maxRetriesPerDBConnectionAttempt),
		database.WithDatabaseRetryTimeOut(maxDBRetryTimeout),
		database.WithDatabaseOperationSleepInterval(maxDBSleepInterval),
		database.WithDatabaseLogger(logger),
		database.WithDatabaseInstrumentation(instrumentation),
		database.WithDatabaseQueryOperator(queryOperator),
	}

	db, err := database.New(ctx, opts...)
	if err != nil {
		return nil, err
	}

	logger.Info("successfully initialized database connection object")
	return db, nil
}

// configureServiceTelemetryInstance configures the service telemetry instance object
func configureServiceTelemetryInstance(logger *zap.Logger) *instrumentation.ServiceTelemetry {
	// configure new relic sdk
	app, err := configureNewrelicSDK(logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	metricsEnabled := viper.GetBool("metrics-reporting-enabled")
	env := viper.GetString("service-environment")
	serviceName := viper.GetString("grpc-service-name")

	opts := []instrumentation.ServiceTelemetryOption{
		instrumentation.WithServiceName(serviceName),
		instrumentation.WithServiceVersion(version.VERSION),
		instrumentation.WithServiceEnvironment(env),
		instrumentation.WithMetricsEnabled(metricsEnabled),
		instrumentation.WithNewrelicSdk(app),
	}

	return instrumentation.NewServiceTelemetry(opts...)
}
