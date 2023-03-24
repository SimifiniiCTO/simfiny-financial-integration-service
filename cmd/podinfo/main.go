package main

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
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
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
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

	// load gRPC server config
	var grpcCfg grpc.Config
	if err := viper.Unmarshal(&grpcCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	keyManagement, err := configureKeyManagement(logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	db, err := configureDatabaseConn(ctx, logger, instrumentation, keyManagement)
	if err != nil {
		logger.Panic(err.Error())
	}

	conn, err := db.Conn.Engine.DB()
	if err != nil {
		logger.Panic(err.Error())
	}
	defer conn.Close()

	plaidWrapper, err := configurePlaidWrapper(instrumentation, logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	// initialize gRPC server
	grpcSrv, err := grpc.NewServer(&grpc.Params{
		Config:          &grpcCfg,
		Logger:          logger,
		Db:              db,
		PlaidClient:     plaidClient,
		Instrumentation: instrumentation,
		KeyManagement:   keyManagement,
		PlaidWrapper:    plaidWrapper,
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

// configureKeyManagement configures the key management (aws) sdk to be used by the service
func configureKeyManagement(logger *zap.Logger) (secrets.KeyManagement, error) {
	region := viper.GetString("aws-region")
	accessKey := viper.GetString("aws-access-key-id")
	secretAccessKey := viper.GetString("aws-secret-access-key")
	kmsKeyID := viper.GetString("aws-kms-id")

	keyManagement, err := secrets.NewAWSKMS(secrets.AWSKMSConfig{
		Region:    region,
		AccessKey: accessKey,
		SecretKey: secretAccessKey,
		KmsKeyID:  kmsKeyID,
		Log:       logger,
	})
	if err != nil {
		return nil, err
	}

	return keyManagement, nil
}

// configurNewrelicSDK configures the new relic sdk with metadata specific to this service
func configureNewrelicSDK(logger *zap.Logger) (*newrelic.Application, error) {
	var newrelicLicenseKey = viper.GetString("newrelic-api-key")
	var serviceName = viper.GetString("grpc-service-name")

	if newrelicLicenseKey != "" {
		return newrelic.NewApplication(
			newrelic.ConfigAppName(serviceName),
			newrelic.ConfigLicense(newrelicLicenseKey),
			newrelic.ConfigAppLogForwardingEnabled(true),
			func(cfg *newrelic.Config) {
				cfg.ErrorCollector.RecordPanics = true
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
				cfg.Logger = nrzap.Transform(logger)
			},
			// Use nrzap to register the logger with the agent:
			nrzap.ConfigLogger(logger.Named("newrelic")),
			newrelic.ConfigDistributedTracerEnabled(true),
			newrelic.ConfigEnabled(true),
		)

	}

	return nil, fmt.Errorf("invalid input parameter. param: newrelicLicenseKey = %s", newrelicLicenseKey)
}

// configurePlaidWrapper configures the plaid sdk wrapper to be userd by the service
func configurePlaidWrapper(instrimentation *instrumentation.ServiceTelemetry, logger *zap.Logger) (*plaidhandler.PlaidWrapper, error) {
	var (
		plaidClientID       = viper.GetString("plaid-client-id")
		plaidSecretKey      = viper.GetString("plaid-secret-key")
		plaidEnv            = viper.GetString("plaid-env")
		oauthDomain         = viper.GetString("plaid-oauth-domain")
		webhooksEnabled     = viper.GetBool("plaid-webhooks-enabled")
		webhooksOauthDomain = viper.GetString("plaid-webhooks-oauth-domain")
	)

	env, err := decipherPlaidEnvironment(plaidEnv)
	if err != nil {
		return nil, err
	}

	opts := []plaidhandler.Option{
		plaidhandler.WithEnvironment(env),
		plaidhandler.WithClientID(plaidClientID),
		plaidhandler.WithSecretKey(plaidSecretKey),
		plaidhandler.WithInstrumentation(instrimentation),
		plaidhandler.WithLogger(logger),
		plaidhandler.WithOauthDomain(oauthDomain),
		plaidhandler.WithWebhooksDomain(webhooksOauthDomain),
		plaidhandler.WithWebhooksEnabled(webhooksEnabled),
	}

	// declare plaid wrapper
	handler, err := plaidhandler.New(opts...)
	if err != nil {
		return nil, err
	}

	return handler, nil

}

// decipherPlaidEnvironment deciphers the plaid environment from the config env
func decipherPlaidEnvironment(environment string) (plaid.Environment, error) {
	var (
		env plaid.Environment
	)

	if environment == "" {
		return "", fmt.Errorf("environment cannot be empty")
	}

	switch environment {
	case "sandbox":
		env = plaid.Sandbox
	case "development":
		env = plaid.Development
	case "production":
		env = plaid.Production
	default:
		return "", fmt.Errorf("invalid plaid environment: %s", environment)
	}

	return env, nil
}

// configurePlaidSDK configures the plaid sdk to be used by the service
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
func configureDatabaseConn(ctx context.Context, logger *zap.Logger, instrumentation *instrumentation.ServiceTelemetry, keyManagment secrets.KeyManagement) (*database.Db, error) {
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
	dbQueryTimeout := viper.GetDuration("max-query-timeout")

	connectionString := database.ConfigureConnectionString(host, user, password, dbname, sslMode, port)

	// establish db connections
	conn := core_database.NewDatabaseConn(
		&core_database.Parameters{
			QueryTimeout:              &dbQueryTimeout,
			MaxConnectionRetries:      &maxDBConnAttempts,
			MaxConnectionRetryTimeout: &maxDBRetryTimeout,
			RetrySleep:                &maxDBSleepInterval,
			ConnectionString:          &connectionString,
		})

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
		database.WithKms(keyManagment),
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
