package transactionmanager

import (
	"time"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	database "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"
)

type Configs struct {
	ClientOptions                      *client.Options
	Logger                             *zap.Logger
	TelemetrySDK                       *instrumentation.Client
	DatabaseConn                       *database.Db
	RetryPolicy                        *Policy
	RpcTimeout                         *time.Duration
	MetricsEnabled                     bool
	MaxConcurrentActivityExecutionSize uint64
	MaxConcurrentWorkflowTaskPollers   uint64
	ServiceTaskQueue                   string
}

// Option is a function that configures the transaction manager
type Option func(*Configs)

// WithClientOptions is an option that configures the transaction manager with a temporal client options
func WithClientOptions(clientOptions *client.Options) Option {
	return func(c *Configs) {
		c.ClientOptions = clientOptions
	}
}

// WithLogger is an option that configures the transaction manager with a logger
func WithLogger(logger *zap.Logger) Option {
	return func(c *Configs) {
		c.Logger = logger
	}
}

// WithTelemetrySDK is an option that configures the transaction manager with a telemetry sdk
func WithTelemetrySDK(telemetrySDK *instrumentation.Client) Option {
	return func(c *Configs) {
		c.TelemetrySDK = telemetrySDK
	}
}

// WithDatabaseConn is an option that configures the transaction manager with a database connection
func WithDatabaseConn(databaseConn *database.Db) Option {
	return func(c *Configs) {
		c.DatabaseConn = databaseConn
	}
}

// WithRetryPolicy is an option that configures the transaction manager with a retry policy
func WithRetryPolicy(retryPolicy *Policy) Option {
	return func(c *Configs) {
		c.RetryPolicy = retryPolicy
	}
}

// WithRpcTimeout is an option that configures the transaction manager with a rpc timeout
func WithRpcTimeout(rpcTimeout *time.Duration) Option {
	return func(c *Configs) {
		c.RpcTimeout = rpcTimeout
	}
}

// WithMetricsEnabled is an option that configures the transaction manager with a metrics enabled
func WithMetricsEnabled(metricsEnabled bool) Option {
	return func(c *Configs) {
		c.MetricsEnabled = metricsEnabled
	}
}

// WithMaxConcurrentActivityExecutionSize is an option that configures the transaction manager with a max concurrent activity execution size
func WithMaxConcurrentActivityExecutionSize(maxConcurrentActivityExecutionSize uint64) Option {
	return func(c *Configs) {
		c.MaxConcurrentActivityExecutionSize = maxConcurrentActivityExecutionSize
	}
}

// WithMaxConcurrentWorkflowTaskPollers is an option that configures the transaction manager with a max concurrent workflow task pollers
func WithMaxConcurrentWorkflowTaskPollers(maxConcurrentWorkflowTaskPollers uint64) Option {
	return func(c *Configs) {
		c.MaxConcurrentWorkflowTaskPollers = maxConcurrentWorkflowTaskPollers
	}
}

// WithServiceTaskQueue is an option that configures the transaction manager with a service task queue
func WithServiceTaskQueue(serviceTaskQueue string) Option {
	return func(c *Configs) {
		c.ServiceTaskQueue = serviceTaskQueue
	}
}

// NewConfigs creates a new transaction manager config
func NewConfigs(opts ...Option) (*Configs, error) {
	configs := &Configs{}

	for _, opt := range opts {
		opt(configs)
	}

	if err := configs.validate(); err != nil {
		return nil, err
	}

	return configs, nil
}

// validate validates the transaction manager config
func (c *Configs) validate() error {
	if c == nil {
		return ErrInvalidTransactionManagerConfig
	}

	if c.RpcTimeout == nil {
		return ErrInvalidRpcTimeout
	}

	if c.RetryPolicy == nil {
		return ErrInvalidRetryPolicy
	}

	if c.ClientOptions == nil {
		return ErrInvalidClientOptions
	}

	if c.Logger == nil {
		return ErrInvalidLogger
	}

	if c.TelemetrySDK == nil {
		return ErrInvalidTelemetrySDK
	}

	return nil
}
