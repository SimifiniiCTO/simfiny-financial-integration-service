package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/version"
)

// ReadEnvVars reads environment variables and flags
func ReadEnvVars() {
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
	fs.Duration("rpc-timeout", 30*time.Second, "client timeout duration")
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
	fs.String("newrelic-api-key", "62fd721c712d5863a4e75b8f547b7c1ea884NRAL", "new relic license key")
	// database connection environment variables
	fs.String("dbhost", "db", "database host string")
	fs.Int("dbport", 5432, "database port")
	fs.String("dbuser", "postgres", "database user string")
	fs.String("dbpassword", "postgres", "database password string")
	fs.String("dbname", "postgres", "database name")
	fs.String("dbsslmode", "disable", "wether tls connection is enabled")
	fs.Int("max-db-conn-attempts", 1, "max database connection attempts")
	fs.Int("max-db-conn-retries", 1, "max database connection attempts")
	fs.Duration("max-db-retry-timeout", 500*time.Millisecond, "max time until a db connection request is seen as timing out")
	fs.Duration("max-db-retry-sleep-interval", 100*time.Millisecond, "max time to sleep in between db connection attempts")
	fs.Duration("max-query-timeout", 500*time.Millisecond, "max time until a db query is seen as timing out")

	// plaid specific keys
	fs.String("plaid-client-id", "61eb5d49ea3b4700127560d1", "plaid client id")
	fs.String("plaid-secret-key", "465686056e8fd1b87db3d993d096d8", "plaid secret key")
	fs.String("plaid-env", "sandbox", "plaid environment")
	fs.String("plaid-webhook-url", "http://localhost:3000/webhook", "plaid webhook url")
	fs.String("plaid-redirect-url", "http://localhost:3000/", "plaid redirect url")
	fs.StringSlice("plaid-products", []string{"transactions", "auth", "balance", "investments", "liabilities"}, "plaid products to enable")
	fs.String("plaid-oauth-domain", "simfiny", "plaid oauth domain")
	fs.String("plaid-webhook-oauth-domain", "simfiny", "plaid webhook oauth domain")
	fs.Bool("plaid-webhook-enabled", true, "enable plaid webhook")
	fs.Uint64("max-plaid-links", 5, "plaid link limit")

	fs.String("grpc-service-endpoint", "http://localhost:9896", "grpc api endpoint for service")
	fs.Int("grpc-gateway-port", 8090, "grpc gateway port")

	fs.String("service-environment", "dev", "environment in which service is running")
	fs.String("service-documentation", "https://github.com/SimifiniiCTO/simfinii/blob/main/src/backend/services/simfiny-financial-integration-service/documentation/setup.md", "location of service docs")
	fs.String("point-of-contact", "yoanyomba", "service point of contact")
	fs.Bool("metrics-reporting-enabled", true, "enable metrics reporting")

	// aws configs
	// to access aws credentials, ref the following link
	// ref: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html
	fs.String("aws-access-key-id", "AKIA5HFOAJRN7YDEYPST", "aws kms key id")
	fs.String("aws-region", "us-east-2", "aws kms region")
	fs.String("aws-secret-access-key", "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", "aws kms secret key")
	fs.String("aws-kms-id", "mrk-e44f269bc0034feb95ede34154c3cfe4", "aws s3 bucket")

	// billing configs
	fs.Bool("stripe-enabled", false, "enable billing")
	fs.String("stripe-api-key", "sk_test_4eC39HqLyjWDarjtT1zdp7dc", "")

	// temporal configurations
	fs.String("temporal-cluster-endpoint", "temporal:7233", "base endpoint of the temporal cluster")
	fs.String("temporal-namespace", "simfiny", "temporal namespace to operate on")
	fs.Duration("temporal-retry-initial-interval", 100*time.Millisecond, "retry initial interval")
	fs.Float64("temporal-backoff-coefficient", 2.5, "backoff coefficient")
	fs.Int("temporal-max-retry-attempts", 5, "maximum number of retry attempts")
	fs.Duration("temporal-max-retry-interval", 5*time.Second, "max retry interval")
	fs.Duration("temporal-rpc-timeout", 2*time.Second, "end to end rpc timeout interval")
	fs.String("temporal-task-queue", "simfiny", "temporal task queue")
	fs.Int64("temporal-concurrent-activity-execution-size", 5000, "temporal concurrent activity execution size")
	fs.Int64("temporal-concurrent-workflow-task-pollers", 100, "temporal concurrent workflow task pollers")
	fs.Duration("workflow-execution-timeout", 1*time.Second, "e timeout for duration of workflow execution. It includes retries and continue as new. Use WorkflowRunTimeout to limit execution time of a single workflow run.")
	fs.Duration("workflow-task-timeout", 1*time.Second, "The timeout for processing workflow task from the time the worker pulled this task. If a workflow task is lost, it is retried after this timeout. The resolution is seconds.")
	fs.Duration("workflow-run-timeout", 1*time.Second, "The timeout for duration of a single workflow run. The resolution is seconds. Optional: defaulted to WorkflowExecutionTimeout.")
	defaultLogger := zap.L()

	// parse flags
	parseFlags(fs)

	// load environment variables
	loadEnvVariables(fs)

	// validate environment configurations
	if err := validateEnvConfigs(fs); err != nil {
		defaultLogger.Panic(err.Error())
	}
	startStressTestsIfEnabled()
}

func parseFlags(fs *pflag.FlagSet) {
	versionFlag := fs.BoolP("version", "v", false, "get version number")

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
	viper.BindPFlags(fs)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}

func loadEnvVariables(fs *pflag.FlagSet) {
	viper.AddConfigPath("/go/src/github.com/SimifiniiCTO/simfiny-financial-integration-service")
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

func validateEnvConfigs(fs *pflag.FlagSet) error {
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
		return fmt.Errorf("`--random-delay-max` should be greater than `--random-delay-min`")
	}

	switch delayUnit := viper.GetString("random-delay-unit"); delayUnit {
	case
		"s",
		"ms":
		break
	default:
		return fmt.Errorf("`random-delay-unit` accepted values are: s|ms")
	}

	return nil
}

func startStressTestsIfEnabled() {
	// start stress tests if any
	beginStressTest(viper.GetInt("stress-cpu"), viper.GetInt("stress-memory"), zap.L())

}

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

		stressMemoryPayload, err := ioutil.ReadFile(path)
		f.Close()
		os.Remove(path)
		if err != nil {
			logger.Error("memory stress failed", zap.Error(err))
		}
		logger.Info("starting CPU stress", zap.Int("memory", len(stressMemoryPayload)))
	}
}
