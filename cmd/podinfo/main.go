package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	telemetry "github.com/SimifiniiCTO/core/core-telemetry"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/labstack/gommon/log"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/plaid"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"

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
	// flags definition
	fs := pflag.NewFlagSet("default", pflag.ContinueOnError)
	fs.String("host", "", "Host to bind service to")
	fs.Int("port", 9898, "HTTP port to bind service to")
	fs.Int("secure-port", 0, "HTTPS port")
	fs.Int("port-metrics", 0, "metrics port")
	fs.Int("grpc-port", 9896, "gRPC port")
	fs.String("grpc-service-name", "financial-integration-service", "gPRC service name")
	fs.String("level", "info", "log level debug, info, warn, error, flat or panic")
	fs.StringSlice("backend-url", []string{}, "backend service URL")
	fs.Duration("http-client-timeout", 2*time.Minute, "client timeout duration")
	fs.Duration("http-server-timeout", 30*time.Second, "server read and write timeout duration")
	fs.Duration("http-server-shutdown-timeout", 5*time.Second, "server graceful shutdown timeout duration")
	fs.String("data-path", "/data", "data local path")
	fs.String("config-path", "", "config dir path")
	fs.String("cert-path", "/data/cert", "certificate path for HTTPS port")
	fs.String("config", "config.yaml", "config file name")
	fs.String("ui-path", "./ui", "UI local path")
	fs.String("ui-logo", "", "UI logo")
	fs.String("ui-color", "#34577c", "UI color")
	fs.String("ui-message", fmt.Sprintf("greetings from financial-integration-service v%v", version.VERSION), "UI message")
	fs.Bool("h2c", false, "allow upgrading to H2C")
	fs.Bool("random-delay", false, "between 0 and 5 seconds random delay by default")
	fs.String("random-delay-unit", "s", "either s(seconds) or ms(milliseconds")
	fs.Int("random-delay-min", 0, "min for random delay: 0 by default")
	fs.Int("random-delay-max", 5, "max for random delay: 5 by default")
	fs.Bool("random-error", false, "1/3 chances of a random response error")
	fs.Bool("unhealthy", false, "when set, healthy state is never reached")
	fs.Bool("unready", false, "when set, ready state is never reached")
	fs.Int("stress-cpu", 0, "number of CPU cores with 100 load")
	fs.Int("stress-memory", 0, "MB of data to load into memory")
	fs.String("cache-server", "", "Redis address in the format <host>:<port>")
	// TODO: ensure this comes from env variables
	fs.String("newrelic-key", "62fd721c712d5863a4e75b8f547b7c1ea884NRAL", "new relic license key")
	// database connection environment variables
	fs.String("dbhost", "financial_integration_svc_db", "database host string")
	fs.Int("dbport", 5432, "database port")
	fs.String("dbuser", "financial_integration_svc", "database user string")
	fs.String("dbpassword", "financial_integration_svc", "database password string")
	fs.String("dbname", "financial_integration_svc", "database name")
	fs.String("dbsslmode", "disable", "wether tls connection is enabled")
	fs.Int("max-db-conn-attempts", 2, "max database connection attempts")
	fs.Int("max-db-conn-retries", 2, "max database connection attempts")
	fs.Duration("max-db-retry-timeout", 500*time.Millisecond, "max time until a db connection request is seen as timing out")
	fs.Duration("max-db-retry-sleep-interval", 100*time.Millisecond, "max time to sleep in between db connection attempts")
	// plaid specific keys
	fs.String("plaid-client-id", "61eb5d49ea3b4700127560d1", "plaid client id")
	fs.String("plaid-secret-key", "465686056e8fd1b87db3d993d096d8", "plaid secret key")
	fs.String("plaid-env", "sandbox", "plaid environment")
	fs.String("plaid-webhook-url", "", "plaid webhook url")
	fs.String("plaid-redirect-url", "", "plaid redirect url")
	fs.String("env", "dev", "current environment")
	fs.String("grpc-service-endpoint", "http://localhost:9896", "grpc api endpoint for service")
	fs.Int("grpc-gateway-port", 8090, "grpc gateway port")

	fs.String("service-environment", "dev", "environment in which service is running")
	fs.String("service-documentation", "https://github.com/SimifiniiCTO/simfinii/blob/main/src/backend/services/user-service/documentation/setup.md", "location of service docs")
	fs.String("point-of-contact", "yoanyomba", "service point of contact")
	fs.Bool("metrics-reporting-enabled", true, "enable metrics reporting")

	ctx := context.Background()
	versionFlag := fs.BoolP("version", "v", false, "get version number")

	// parse flags
	err := fs.Parse(os.Args[1:])
	switch {
	case err == pflag.ErrHelp:
		os.Exit(0)
	case err != nil:
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err.Error())
		fs.PrintDefaults()
		os.Exit(2)
	case *versionFlag:
		fmt.Println(version.VERSION)
		os.Exit(0)
	}

	// bind flags and environment variables
	LoadEnvVariables(fs)
	// Create a new boot instance.
	boot := rkboot.NewBoot()
	// configure logging
	logger := rkentry.GlobalAppCtx.GetLoggerEntry("zap-logger").Logger

	// load config from file
	if _, fileErr := os.Stat(filepath.Join(viper.GetString("config-path"), viper.GetString("config"))); fileErr == nil {
		viper.SetConfigName(strings.Split(viper.GetString("config"), ".")[0])
		viper.AddConfigPath(viper.GetString("config-path"))
		if readErr := viper.ReadInConfig(); readErr != nil {
			fmt.Printf("Error reading config file, %v\n", readErr)
		}
	}

	// configure new relic sdk
	nrClient, err := InitializeNewrelicSDK(logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	svcTelemetry, err := initTelemetrySDK(logger, nrClient)
	if err != nil {
		log.Panic(err.Error())
	}

	// configure plaid client
	var plaidClientID = viper.GetString("plaid-client-id")
	var plaidSecretKey = viper.GetString("plaid-secret-key")
	var plaidEnv = viper.GetString("plaid-env")
	plaidClient, err := InitializePlaidClient(plaidClientID, plaidSecretKey, plaidEnv)
	if err != nil {
		logger.Panic(err.Error())
	}

	// start stress tests if any
	beginStressTest(viper.GetInt("stress-cpu"), viper.GetInt("stress-memory"), logger)

	// validate port
	if _, err := strconv.Atoi(viper.GetString("port")); err != nil {
		port, _ := fs.GetInt("port")
		viper.Set("port", strconv.Itoa(port))
	}

	// validate secure port
	if _, err := strconv.Atoi(viper.GetString("secure-port")); err != nil {
		securePort, _ := fs.GetInt("secure-port")
		viper.Set("secure-port", strconv.Itoa(securePort))
	}

	// validate random delay options
	if viper.GetInt("random-delay-max") < viper.GetInt("random-delay-min") {
		logger.Panic("`--random-delay-max` should be greater than `--random-delay-min`")
	}

	switch delayUnit := viper.GetString("random-delay-unit"); delayUnit {
	case
		"s",
		"ms":
		break
	default:
		logger.Panic("`random-delay-unit` accepted values are: s|ms")
	}

	db, err := InitializeDbConn(ctx, logger, nrClient)
	if err != nil {
		logger.Panic(err.Error())
	}

	conn, err := db.Conn.Engine.DB()
	if err != nil {
		logger.Panic(err.Error())
	}
	defer conn.Close()

	// load gRPC server config
	var grpcCfg grpc.Config
	if err := viper.Unmarshal(&grpcCfg); err != nil {
		logger.Panic("config unmarshal failed", zap.Error(err))
	}

	// start gRPC server
	if grpcCfg.Port > 0 {
		grpcSrv, _ := grpc.NewServer(&grpcCfg, logger, nrClient, db, plaidClient, svcTelemetry.Engine)

		grpcEntry := rkgrpc.GetGrpcEntry("financial-integration-service")
		grpcEntry.AddRegFuncGrpc(grpcSrv.RegisterGrpcServer)
		grpcEntry.AddRegFuncGw(proto.RegisterFinancialIntegrationServiceHandlerFromEndpoint)

		grpcEntry.AddUnaryInterceptors(
			svcTelemetry.RequestLatencyUnaryServerInterceptor(),
			svcTelemetry.RequestCountUnaryServerInterceptor(),
			svcTelemetry.ErrorCountUnaryServerInterceptor(),
			svcTelemetry.RequestTimeUnaryServerInterceptor(),
			svcTelemetry.TxnUnaryServerInterceptor())

		grpcEntry.AddStreamInterceptors(
			svcTelemetry.RequestLatencyStreamServerInterceptor(),
			svcTelemetry.RequestCountStreamServerInterceptor(),
			svcTelemetry.ErrorCountStreamServerInterceptor())

		// Bootstrap
		boot.Bootstrap(context.Background())

		// Wait for shutdown sig
		boot.WaitForShutdownSig(context.Background())
	}

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
	srv, _ := api.NewServer(&srvCfg, logger, nrClient, db, plaidClient)
	stopCh := signals.SetupSignalHandler()
	srv.ListenAndServe(stopCh)
}

// InitializeNewrelicSDK configures the new relic sdk with metadata specific to this service
func InitializeNewrelicSDK(logger *zap.Logger) (*newrelic.Application, error) {
	var newrelicLicenseKey = viper.GetString("newrelic-key")
	var serviceName = viper.GetString("grpc-service-name")

	if newrelicLicenseKey != "" {
		return newrelic.NewApplication(
			newrelic.ConfigAppName(serviceName),
			newrelic.ConfigLicense(newrelicLicenseKey),
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
			},
			// Use nrzap to register the logger with the agent:
			nrzap.ConfigLogger(logger.Named("newrelic")),
			newrelic.ConfigDistributedTracerEnabled(true),
			newrelic.ConfigEnabled(true),
		)
	}

	return nil, errors.New(fmt.Sprintf("invalid input parameter. param: newrelicLicenseKey = %s", newrelicLicenseKey))
}

// initTelemetrySDK initializes the telemetry engine with the parameters provided in the config or environment variable file
func initTelemetrySDK(logger *zap.Logger, nrClient *newrelic.Application) (*telemetry.Telemetry, error) {
	serviceName := viper.GetString("grpc-service-name")
	licenseKey := viper.GetString("newrelic-key")
	version := viper.GetString("version")
	environment := viper.GetString("service-environment")
	pointOfContact := viper.GetString("point-of-contact")
	metricsReportingEnabled := viper.GetBool("metrics-reporting-enabled")

	params := &telemetry.TelemetryParams{
		ServiceName:        serviceName,
		Version:            version,
		PointOfContact:     pointOfContact,
		Environment:        environment,
		NewRelicLicenseKey: licenseKey,
		Logger:             logger,
		Enabled:            metricsReportingEnabled,
		Client:             nrClient,
	}

	engine, err := telemetry.New(params)
	if err != nil {
		return nil, err
	}

	return engine, nil
}

func InitializePlaidClient(plaidClientID, plaidSecretKey, plaidEnv string) (*plaid.APIClient, error) {
	if plaidClientID == "" || plaidSecretKey == "" {
		return nil, errors.New(fmt.Sprintf("invalid input arguments. plaidClientId: %s, plaidSecretKey: %s", plaidClientID, plaidSecretKey))
	}
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", plaidClientID)
	configuration.AddDefaultHeader("PLAID-SECRET", plaidSecretKey)
	configuration.UseEnvironment(environments[plaidEnv])
	return plaid.NewAPIClient(configuration), nil
}

func InitializeDbConn(ctx context.Context, logger *zap.Logger, telemetry *newrelic.Application) (*database.Db, error) {
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

	// initialize a newrelic txn to trace the db connection event
	txn := telemetry.StartTransaction("database-connection")
	defer txn.End()

	initializationParams := &database.ConnectionInitializationParams{
		ConnectionParams: &database.ConnectionParameters{
			Host:         host,
			User:         user,
			Password:     password,
			DatabaseName: dbname,
			Port:         port,
			SslMode:      sslMode,
		},
		Logger:                 logger,
		MaxConnectionAttempts:  maxDBConnAttempts,
		MaxRetriesPerOperation: maxRetriesPerDBConnectionAttempt,
		RetryTimeOut:           maxDBRetryTimeout,
		RetrySleepInterval:     maxDBSleepInterval,
		Telemetry:              telemetry,
	}

	db, err := database.New(ctx, initializationParams)
	if err != nil {
		return nil, err
	}

	logger.Info("successfully initialized database connection object")
	return db, nil
}

var stressMemoryPayload []byte

func beginStressTest(cpus int, mem int, logger *zap.Logger) {
	done := make(chan int)
	if cpus > 0 {
		logger.Info("starting CPU stress", zap.Int("cores", cpus))
		for i := 0; i < cpus; i++ {
			go func() {
				for {
					select {
					case <-done:
						return
					default:

					}
				}
			}()
		}
	}

	if mem > 0 {
		path := "/tmp/financial_integration_service.data"
		f, err := os.Create(path)

		if err != nil {
			logger.Error("memory stress failed", zap.Error(err))
		}

		if err := f.Truncate(1000000 * int64(mem)); err != nil {
			logger.Error("memory stress failed", zap.Error(err))
		}

		stressMemoryPayload, err = ioutil.ReadFile(path)
		f.Close()
		os.Remove(path)
		if err != nil {
			logger.Error("memory stress failed", zap.Error(err))
		}
		logger.Info("starting CPU stress", zap.Int("memory", len(stressMemoryPayload)))
	}
}

// LoadEnvVariables binds a set of flags to and loads environment variables
func LoadEnvVariables(fs *pflag.FlagSet) {
	viper.AddConfigPath("/go/src/github.com/SimifiniiCTO/simfinii/src/backend/services/user-service")
	viper.BindPFlags(fs)
	viper.RegisterAlias("BACKEND_SERVICE_URLS", "BACKEND_URL")
	viper.SetConfigName("financial-integration-service")
	viper.SetConfigType("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	hostname, _ := os.Hostname()
	viper.SetDefault("JWT_SECRET", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	viper.SetDefault("UI_LOGO", "https://raw.githubusercontent.com/stefanprodan/podinfo/gh-pages/cuddle_clap.gif")
	viper.Set("hostname", hostname)
	viper.Set("version", version.VERSION)
	viper.Set("revision", version.REVISION)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}
}
