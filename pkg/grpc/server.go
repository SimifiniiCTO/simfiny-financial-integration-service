package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/plaid/plaid-go/v14/plaid"
	"github.com/stripe/stripe-go/v74/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	telemetry "github.com/SimifiniiCTO/core/core-telemetry"
	"github.com/SimifiniiCTO/simfiny-core-lib/database/redis"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	taskprocessor "github.com/SimifiniiCTO/simfiny-core-lib/task-processor"
	clickhousedatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	database "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	taskhandler "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/task-handler"
	transactionmanager "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transaction_manager"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	openai "github.com/sashabaranov/go-openai"
)

// Server is the grpc server object
type Server struct {
	proto.UnimplementedFinancialServiceServer
	// proto.UnimplementedFinancialServiceServer
	logger                      *zap.Logger
	config                      *Config
	instrumentation             *instrumentation.Client
	conn                        *database.Db
	clickhouseConn              *clickhousedatabase.Db
	plaidClient                 *plaidhandler.PlaidWrapper
	MetricsEngine               *telemetry.MetricsEngine
	stripeClient                *client.API
	kms                         secrets.KeyManagement
	TransactionManager          *transactionmanager.TransactionManager
	InMemoryWebhookVerification plaidhandler.WebhookVerification
	redisDb                     *redis.Client
	Taskprocessor               *taskprocessor.TaskProcessor
	OpenAiClient                *openai.Client
	th                          *taskhandler.TaskHandler
}

// Config is the config for the grpc server initialization
type Config struct {
	Port                                  int              `mapstructure:"grpc-port"`
	GatewayPort                           int              `mapstructure:"grpc-gateway-port"`
	ServiceName                           string           `mapstructure:"grpc-service-name"`
	NewRelicLicense                       string           `mapstructure:"newrelic-api-key"`
	Environment                           string           `mapstructure:"env"`
	PlaidProducts                         []plaid.Products `mapstructure:"plaid-products"`
	RpcTimeout                            time.Duration    `mapstructure:"rpc-timeout"`
	StripeApiKey                          string           `mapstructure:"stripe-api-key"`
	PlaidClientID                         string           `mapstructure:"plaid-client-id"`
	PlaidSecretKey                        string           `mapstructure:"plaid-secret-key"`
	PlaidEnv                              string           `mapstructure:"plaid-env"`
	PlaidOauthDomain                      string           `mapstructure:"plaid-oauth-domain"`
	PlaidWebhooksEnabled                  bool             `mapstructure:"plaid-webhooks-enabled"`
	PlaidWebhookOauthDomain               string           `mapstructure:"plaid-webhook-oauth-domain"`
	AwsAccessKeyID                        string           `mapstructure:"aws-access-key-id"`
	AwsRegion                             string           `mapstructure:"aws-region"`
	AwsSecretAccessKey                    string           `mapstructure:"aws-secret-access-key"`
	AwsKmsKeyID                           string           `mapstructure:"aws-kms-key-id"`
	MaxPlaidLinks                         int              `mapstructure:"max-plaid-links"`
	BillingEnabled                        bool             `mapstructure:"stripe-enabled"`
	WorkflowExecutionTimeout              time.Duration    `mapstructure:"workflow-execution-timeout"`
	WorkflowTaskTimeout                   time.Duration    `mapstructure:"workflow-task-timeout"`
	WorkflowRunTimeout                    time.Duration    `mapstructure:"workflow-run-timeout"`
	TaskProcessorWorkers                  int              `mapstructure:"task-processor-workers"`
	OpenAiMaxTokens                       int32            `mapstructure:"openai-max-tokens"`
	OpenAiTopP                            float32          `mapstructure:"openai-top-p"`
	OpenAiFrequencyPenalty                float32          `mapstructure:"openai-frequency-penalty"`
	OpenAiPresencePenalty                 int              `mapstructure:"openai-presence-penalty"`
	OpenAiTemperature                     float32          `mapstructure:"openai-temperature"`
	BatchJobIntervalRecurringTransactions string           `mapstructure:"batch-job-interval-recurring-transactions"`
	BatchJobIntervalActionableInsights    string           `mapstructure:"batch-job-interval-actionable-insights"`
	BatchJobIntervalSyncAllAccounts       string           `mapstructure:"batch-job-interval-sync-all-accounts"`
}

var _ proto.FinancialServiceServer = (*Server)(nil)

// Params is the params for the grpc server initialization
type Params struct {
	Config             *Config
	Logger             *zap.Logger
	Instrumentation    *instrumentation.Client
	Db                 *database.Db
	KeyManagement      secrets.KeyManagement
	PlaidWrapper       *plaidhandler.PlaidWrapper
	TransactionManager *transactionmanager.TransactionManager
	ClickhouseDb       *clickhousedatabase.Db
	RedisDb            *redis.Client
	OpenAiToken        *string
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

	if p.OpenAiToken == nil {
		return errors.New("open ai token is nil")
	}

	return nil
}

func configureTaskHandler(param *Params, client *openai.Client) *taskhandler.TaskHandler {
	opts := []taskhandler.Option{
		taskhandler.WithLogger(param.Logger),
		taskhandler.WithInstrumentationClient(param.Instrumentation),
		taskhandler.WithClickhouseDb(param.ClickhouseDb),
		taskhandler.WithPostgresDb(param.Db),
		taskhandler.WithPlaidClient(param.PlaidWrapper),
		taskhandler.WithOpenAIClient(client),
	}

	return taskhandler.NewTaskHandler(opts...)
}

// NewServer creates a new grpc server
func NewServer(param *Params) (*Server, error) {
	if param.Validate() != nil {
		return nil, errors.New("invalid param")
	}

	// initialize a stripe connection for the user
	sc := &client.API{}
	sc.Init(param.Config.StripeApiKey, nil)

	openAiClient := openai.NewClient(*param.OpenAiToken)
	th := configureTaskHandler(param, openAiClient)

	// generatee task procerssor
	opts := []taskprocessor.Option{
		taskprocessor.WithLoggerOpt(param.Logger),
		taskprocessor.WithInstrumentationClientOpt(param.Instrumentation),
		taskprocessor.WithRedisAddressOpt(param.RedisDb.URI),
		taskprocessor.WithConcurrencyFactorOpt(&param.Config.TaskProcessorWorkers),
		taskprocessor.WithTaskHandlerOpt(th),
	}

	tp, err := taskprocessor.NewTaskProcessor(opts...)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		logger:                      param.Logger,
		config:                      param.Config,
		conn:                        param.Db,
		plaidClient:                 param.PlaidWrapper,
		instrumentation:             param.Instrumentation,
		stripeClient:                sc,
		kms:                         param.KeyManagement,
		TransactionManager:          param.TransactionManager,
		InMemoryWebhookVerification: plaidhandler.NewInMemoryWebhookVerification(5*time.Minute, param.Logger, param.PlaidWrapper),
		clickhouseConn:              param.ClickhouseDb,
		redisDb:                     param.RedisDb,
		Taskprocessor:               tp,
		th:                          th,
	}

	if param.OpenAiToken == nil {
		return nil, errors.New("open ai token is required")
	}

	srv.OpenAiClient = openAiClient

	if err := srv.registerBatchJobs(); err != nil {
		return nil, err
	}

	return srv, nil
}

// registerBatchJobs registers the batch jobs
func (s *Server) registerBatchJobs() error {
	// we first enqueue the task to routinely pull transactions from plaid for all users across all accounts (this should run every 12 hours) (sync)
	syncAllAccountsBatchJob, err := taskhandler.NewSyncAllPlatformConnectedPlaidAccounts()
	if err != nil {
		return err
	}

	generateActionableInsightsBatchJob, err := taskhandler.NewGenerateActionableInsights()
	if err != nil {
		return err
	}

	generateRecurringTransactionsBatchJob, err := taskhandler.NewSyncRecurringTransactions()
	if err != nil {
		return err
	}

	// TODO: we enqueue the task to compute the net worth of all users across all accounts (this should run every 24 hours) (net worth)
	_, err = s.Taskprocessor.EnqueueRecurringTask(
		context.Background(),
		generateRecurringTransactionsBatchJob,
		stringToProcessingInterval(&s.config.BatchJobIntervalRecurringTransactions))
	if err != nil {
		return err
	}

	// TODO: this should be config driven
	_, err = s.Taskprocessor.EnqueueRecurringTask(
		context.Background(),
		syncAllAccountsBatchJob,
		stringToProcessingInterval(&s.config.BatchJobIntervalSyncAllAccounts))
	if err != nil {
		return err
	}

	// TODO: we enqueue the task to compute actionable insights for all users across all accounts (this should run every 24 hours) - use openai for this (insights)
	_, err = s.Taskprocessor.EnqueueRecurringTask(
		context.Background(),
		generateActionableInsightsBatchJob,
		stringToProcessingInterval(&s.config.BatchJobIntervalActionableInsights))
	if err != nil {
		return err
	}

	return nil
}

// stringToProcessingInterval converts a string to a processing interval
func stringToProcessingInterval(str *string) taskprocessor.ProcessingInterval {
	if str == nil {
		// handle nil input in some way, I'll return EveryDay as a default for this example
		return taskprocessor.EveryDay
	}
	switch *str {
	case "@yearly":
		return taskprocessor.EveryYear
	case "@monthly":
		return taskprocessor.EveryMonth
	case "@weekly":
		return taskprocessor.EveryWeek
	case "@midnight":
		return taskprocessor.EveryDayAtMidnight
	case "@daily":
		return taskprocessor.EveryDay
	case "@every 24h":
		return taskprocessor.Every24Hours
	case "@every 12h":
		return taskprocessor.Every12Hours
	case "@every 6h":
		return taskprocessor.Every6Hours
	case "@every 3h":
		return taskprocessor.Every3Hours
	case "@every 1h":
		return taskprocessor.EveryHour
	case "@every 30m":
		return taskprocessor.Every30Minutes
	case "@every 15m":
		return taskprocessor.Every15Minutes
	case "@every 10m":
		return taskprocessor.Every10Minutes
	case "@every 5m":
		return taskprocessor.Every5Minutes
	case "@every 3m":
		return taskprocessor.Every3Minutes
	case "@every 1m":
		return taskprocessor.Every1Minutes
	case "@every 30s":
		return taskprocessor.Every30Seconds
	default:
		// If the string doesn't match any of the predefined intervals, handle it accordingly.
		// I'll return EveryDay as a default for this example.
		return taskprocessor.EveryDay
	}
}
