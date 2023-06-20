package main

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/v12/plaid"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"

	clickhousedb "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"
	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	clickhouseDatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/env"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	database "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/redis"
	"github.com/SimifiniiCTO/simfiny-core-lib/signals"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	transactionmanager "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transaction_manager"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/api"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/grpc"
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
	instrumentation, err := configureInstrumentationClient(logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	// load gRPC server config
	var grpcCfg grpc.Config
	if err := viper.Unmarshal(&grpcCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	var httpCfg api.Config
	if err := viper.Unmarshal(&httpCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	keyManagement, err := configureKeyManagement(logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	redisConn, err := configureRedisConn(logger, instrumentation)
	if err != nil {
		logger.Panic(err.Error())
	}
	defer redisConn.Close()

	db, err := configureDatabaseConn(ctx, logger, instrumentation, keyManagement)
	if err != nil {
		logger.Panic(err.Error())
	}

	conn, err := db.Conn.Engine.DB()
	if err != nil {
		logger.Panic(err.Error())
	}
	defer conn.Close()

	clickHouseDb, err := configureClickhouseConn(ctx, logger, instrumentation)
	if err != nil {
		logger.Panic(err.Error())
	}

	clickHouseConn, err := clickHouseDb.Conn.Engine.DB()
	if err != nil {
		logger.Panic(err.Error())
	}
	defer clickHouseConn.Close()

	plaidWrapper, err := configurePlaidWrapper(instrumentation, logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	logger.Info("successfully initialized plaid wrapper ....")

	transactionManager, err := configureTransactionManager(logger, db, instrumentation)
	if err != nil {
		logger.Panic(err.Error())
	}

	logger.Info("successfully initialized temporal client ....")

	// initialize gRPC server
	grpcSrv, err := grpc.NewServer(&grpc.Params{
		Config:             &grpcCfg,
		Logger:             logger,
		Db:                 db,
		Instrumentation:    instrumentation,
		KeyManagement:      keyManagement,
		PlaidWrapper:       plaidWrapper,
		TransactionManager: transactionManager,
		ClickhouseDb:       clickHouseDb,
		RedisDb:            redisConn,
	})
	if err != nil {
		logger.Panic(err.Error())
	}

	logger.Info("successfully initialized grpc server ....")

	serviceName := viper.GetString("grpc-service-name")

	grpcEntry := rkgrpc.GetGrpcEntry(serviceName)
	grpcEntry.AddRegFuncGrpc(grpcSrv.RegisterGrpcServer)
	grpcEntry.AddRegFuncGw(proto.RegisterFinancialServiceHandlerFromEndpoint)

	logger.Info("successfully configured grpc server ....")

	// TODO: add grpc interceptor middleware to emit metrics on various gRPC calls
	// Bootstrap

	// ensure we can optimatelly close the transaction manager and all associated resources
	defer grpcSrv.TransactionManager.Close()

	// start the worker asynchronously
	go grpcSrv.TransactionManager.StartWorker()

	// load HTTP server config
	var srvCfg api.Config
	if err := viper.Unmarshal(&srvCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	// log version and port
	logger.Info("Starting financial integration service (grpc and http service)",
		zap.String("version", viper.GetString("version")),
		zap.String("revision", viper.GetString("revision")),
		zap.String("port", srvCfg.Port),
	)

	// TODO: clean this up
	srv, err := api.NewServer(&httpCfg, logger, instrumentation, db, nil, plaidWrapper, keyManagement, grpcSrv.Taskprocessor)
	if err != nil {
		logger.Panic(err.Error())
	}

	httpServer, httpSecureServer := srv.ListenAndServe()

	logger.Info("successfully initialized http server ....", zap.Any("srv", srv))

	boot.Bootstrap(context.Background())
	boot.AddShutdownHookFunc("shutdown http server", func() {
		srv.Shutdown(ctx, httpServer, httpSecureServer)
		// TODO: add shutdown hook for grpc server
	})

	// Wait for shutdown sig dgfs
	boot.WaitForShutdownSig(context.Background())
}

func configureTransactionManager(log *zap.Logger, db *database.Db, telemetrySdk *instrumentation.Client) (*transactionmanager.TransactionManager, error) {
	rpcTimeout := viper.GetDuration("temporal-rpc-timeout")
	metricsEnabled := viper.GetBool("metrics-reporting-enabled")
	concurrentActivityExecutionSize := viper.GetUint64("temporal-max-concurrent-activity-execution-size")
	concurrentWorkflowTaskPollers := viper.GetUint64("temporal-max-concurrent-workflow-task-pollers")

	initialRetryInterval := viper.GetDuration("temporal-retry-initial-interval")
	retryBackoffCoefficient := viper.GetFloat64("temporal-retry-backoff-coefficient")
	maximumInterval := viper.GetDuration("temporal-max-retry-interval")
	maxRetryAttempts := viper.GetInt("temporal-max-retry-attempts")

	clusterEndpoint := viper.GetString("temporal-cluster-endpoint")
	namespace := viper.GetString("temporal-namespace")

	taskQueue := viper.GetString("temporal-task-queue")

	policy := &transactionmanager.Policy{
		RetryInitialInterval:    &initialRetryInterval,
		RetryBackoffCoefficient: retryBackoffCoefficient,
		MaximumInterval:         maximumInterval,
		MaximumAttempts:         maxRetryAttempts,
	}
	//
	clientOpts := &client.Options{
		HostPort:  clusterEndpoint,
		Namespace: namespace,
	}

	opts := []transactionmanager.Option{
		transactionmanager.WithClientOptions(clientOpts),
		transactionmanager.WithLogger(log),
		transactionmanager.WithTelemetrySDK(telemetrySdk),
		transactionmanager.WithDatabaseConn(db),
		transactionmanager.WithRetryPolicy(policy),
		transactionmanager.WithRpcTimeout(&rpcTimeout),
		transactionmanager.WithMetricsEnabled(metricsEnabled),
		transactionmanager.WithMaxConcurrentActivityExecutionSize(concurrentActivityExecutionSize),
		transactionmanager.WithMaxConcurrentWorkflowTaskPollers(concurrentWorkflowTaskPollers),
		transactionmanager.WithServiceTaskQueue(taskQueue),
	}

	transactionManager, err := transactionmanager.NewTransactionManager(opts...)
	if err != nil {
		return nil, err
	}

	return transactionManager, nil
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
func configurePlaidWrapper(instrumentation *instrumentation.Client, logger *zap.Logger) (*plaidhandler.PlaidWrapper, error) {
	var (
		plaidClientID       = viper.GetString("plaid-client-id")
		plaidSecretKey      = viper.GetString("plaid-secret-key")
		plaidEnv            = viper.GetString("plaid-env")
		oauthDomain         = viper.GetString("plaid-oauth-domain")
		webhooksEnabled     = viper.GetBool("plaid-webhook-enabled")
		webhooksOauthDomain = viper.GetString("plaid-webhook-oauth-domain")
		products            = viper.GetStringSlice("plaid-products")
	)

	env, err := decipherPlaidEnvironment(plaidEnv)
	if err != nil {
		return nil, err
	}

	opts := []plaidhandler.Option{
		plaidhandler.WithEnvironment(env),
		plaidhandler.WithClientID(plaidClientID),
		plaidhandler.WithSecretKey(plaidSecretKey),
		plaidhandler.WithInstrumentation(instrumentation),
		plaidhandler.WithLogger(logger),
		plaidhandler.WithOauthDomain(oauthDomain),
		plaidhandler.WithWebhooksDomain(webhooksOauthDomain),
		plaidhandler.WithWebhooksEnabled(webhooksEnabled),
		plaidhandler.WithProducts(products),
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

func configureClickhouseConn(ctx context.Context, logger *zap.Logger, instrumentation *instrumentation.Client) (*clickhouseDatabase.Db, error) {
	host := viper.GetString("clickhouse-connection-uri")
	maxConnectionRetries := viper.GetInt("max-db-conn-attempts")
	maxDBRetryTimeout := viper.GetDuration("max-db-retry-timeout")
	maxDBSleepInterval := viper.GetDuration("max-db-retry-sleep-interval")
	queryTimeout := viper.GetDuration("max-query-timeout")

	maxIdleConnections := viper.GetInt("max-db-idle-connections")
	maxOpenConnections := viper.GetInt("max-db-open-connections")
	maxConnectionLifetime := viper.GetDuration("max-db-connection-lifetime")

	opts := []clickhousedb.Option{
		clickhousedb.WithConnectionString(&host),
		clickhousedb.WithQueryTimeout(&queryTimeout),
		clickhousedb.WithMaxConnectionRetries(&maxConnectionRetries),
		clickhousedb.WithMaxConnectionRetryTimeout(&maxDBRetryTimeout),
		clickhousedb.WithRetrySleep(&maxDBSleepInterval),
		clickhousedb.WithMaxIdleConnections(&maxIdleConnections),
		clickhousedb.WithMaxOpenConnections(&maxOpenConnections),
		clickhousedb.WithMaxConnectionLifetime(&maxConnectionLifetime),
		clickhousedb.WithInstrumentationClient(instrumentation),
		clickhousedb.WithLogger(logger),
	}

	conn, err := clickhousedb.New(opts...)
	if err != nil {
		return nil, err
	}

	queryOperator := dal.Use(conn.Engine)
	databaseOpts := []clickhouseDatabase.Option{
		clickhouseDatabase.WithDatabaseClient(conn),
		clickhouseDatabase.WithDatabaseLogger(logger),
		clickhouseDatabase.WithDatabaseInstrumentation(instrumentation),
		clickhouseDatabase.WithDatabaseQueryOperator(queryOperator),
	}

	db, err := clickhouseDatabase.New(ctx, databaseOpts...)
	if err != nil {
		return nil, err
	}

	logger.Info("successfully initialized database connection object")
	return db, nil
}

func configureRedisConn(logger *zap.Logger, instrumentation *instrumentation.Client) (*redis.Client, error) {
	host := viper.GetString("cache-server")
	serviceName := viper.GetString("grpc-service-name")
	cacheTTLInSeconds := viper.GetInt("cache-ttl-in-seconds")
	tlsEnabled := viper.GetBool("cache-tls-enabled")

	stopCh := signals.SetupSignalHandler()
	opts := []redis.Option{
		redis.WithLogger(logger),
		redis.WithTelemetrySdk(instrumentation),
		redis.WithURI(host),
		redis.WithServiceName(serviceName),
		redis.WithCacheTTLInSeconds(cacheTTLInSeconds),
		redis.WithTlsEnabled(tlsEnabled),
	}

	c, err := redis.New(stopCh, opts...)
	if err != nil {
		return nil, err
	}

	logger.Info("successfully initialized redis connection object")
	return c, nil
}

// configureDatabaseConn configures the database connection
func configureDatabaseConn(ctx context.Context, logger *zap.Logger, instrumentation *instrumentation.Client, keyManagment secrets.KeyManagement) (*database.Db, error) {
	host := viper.GetString("dbhost")
	port := viper.GetInt("dbport")
	user := viper.GetString("dbuser")
	password := viper.GetString("dbpassword")
	dbname := viper.GetString("dbname")
	sslMode := viper.GetString("dbsslmode")

	maxConnectionRetries := viper.GetInt("max-db-conn-attempts")
	maxRetriesPerDBConnectionAttempt := viper.GetInt("max-db-conn-retries")
	maxDBRetryTimeout := viper.GetDuration("max-db-retry-timeout")
	maxDBSleepInterval := viper.GetDuration("max-db-retry-sleep-interval")
	queryTimeout := viper.GetDuration("max-query-timeout")

	maxIdleConnections := viper.GetInt("max-db-idle-connections")
	maxOpenConnections := viper.GetInt("max-db-open-connections")
	maxConnectionLifetime := viper.GetDuration("max-db-connection-lifetime")

	connectionString := database.ConfigureConnectionString(host, user, password, dbname, sslMode, port)
	opts := []postgresdb.Option{
		postgresdb.WithConnectionString(&connectionString),
		postgresdb.WithQueryTimeout(&queryTimeout),
		postgresdb.WithMaxConnectionRetries(&maxConnectionRetries),
		postgresdb.WithMaxConnectionRetryTimeout(&maxDBRetryTimeout),
		postgresdb.WithRetrySleep(&maxDBSleepInterval),
		postgresdb.WithMaxIdleConnections(&maxIdleConnections),
		postgresdb.WithMaxOpenConnections(&maxOpenConnections),
		postgresdb.WithMaxConnectionLifetime(&maxConnectionLifetime),
		postgresdb.WithInstrumentationClient(instrumentation),
		postgresdb.WithLogger(logger),
	}

	conn, err := postgresdb.New(opts...)
	if err != nil {
		return nil, err
	}

	queryOperator := dal.Use(conn.Engine)
	databaseOpts := []database.Option{
		database.WithConnection(conn),
		database.WithDatabaseMaxConnectionAttempts(maxConnectionRetries),
		database.WithDatabaseMaxRetriesPerOperation(maxRetriesPerDBConnectionAttempt),
		database.WithDatabaseRetryTimeOut(maxDBRetryTimeout),
		database.WithDatabaseOperationSleepInterval(maxDBSleepInterval),
		database.WithDatabaseLogger(logger),
		database.WithDatabaseInstrumentation(instrumentation),
		database.WithDatabaseQueryOperator(queryOperator),
		database.WithKms(keyManagment),
	}

	db, err := database.New(ctx, databaseOpts...)
	if err != nil {
		return nil, err
	}

	logger.Info("successfully initialized database connection object")
	return db, nil
}

func configureInstrumentationClient(logger *zap.Logger) (*instrumentation.Client, error) {
	// configure new relic sdk
	opts := []instrumentation.Option{
		instrumentation.WithServiceName(viper.GetString("grpc-service-name")),
		instrumentation.WithServiceVersion(version.VERSION),
		instrumentation.WithServiceEnvironment(viper.GetString("service-environment")),
		instrumentation.WithEnabled(viper.GetBool("metrics-reporting-enabled")), // enables instrumentation
		instrumentation.WithNewrelicKey(viper.GetString("newrelic-api-key")),
		instrumentation.WithLogger(logger),
	}

	return instrumentation.New(opts...)
}
