package database

import (
	"context"
	"time"

	"go.uber.org/zap"

	core_database "github.com/SimifiniiCTO/core/core-database"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/service_errors"
)

// DatabaseOperations provides an interface which any database tied to this service should implement
type DatabaseOperations interface {
	// Profile Operations
	// CreateUserProfile creates a user profile
	CreateUserProfile(ctx context.Context, profile *schema.UserProfile) (*schema.UserProfile, error)
	// DeleteUserProfileByUserId deletes a user profile by user id
	DeleteUserProfileByUserID(ctx context.Context, userID uint64) error
	// UpdateUserProfile updates a user profile
	UpdateUserProfile(ctx context.Context, profile *schema.UserProfile) error
	// GetUserProfileByUserID retrieves a user profile by user id
	GetUserProfileByUserID(ctx context.Context, userID uint64) (*schema.UserProfile, error)

	// BankAccount Operations
	// CreateBankAccount creates a bank account
	CreateBankAccount(ctx context.Context, userID uint64, bankAccount *schema.BankAccount) (*schema.BankAccount, error)
	// GetBankAccount retrieves a bank account by id
	GetBankAccount(ctx context.Context, bankAccountID uint64) (*schema.BankAccount, error)
	// DeleteBankAccount deletes a bank account by id
	DeleteBankAccount(ctx context.Context, bankAccountID uint64) error
	// UpdateBankAccount updates a bank account
	UpdateBankAccount(ctx context.Context, bankAccount *schema.BankAccount) error

	// Pocket Operations
	// CreatePocket creates a pocket
	CreatePocket(ctx context.Context, bankAccountID uint64, pocket *schema.Pocket) (*schema.Pocket, error)
	// GetPocket retrieves a pocket by id
	GetPocket(ctx context.Context, pocketID uint64) (*schema.Pocket, error)
	// DeletePocket deletes a pocket by id
	DeletePocket(ctx context.Context, pocketID uint64) error
	// UpdatePocket updates a pocket
	UpdatePocket(ctx context.Context, pocket *schema.Pocket) error

	// Goal Operations
	// CreateGoal creates a goal
	CreateGoal(ctx context.Context, pocketID uint64, goal *schema.SmartGoal) (*schema.SmartGoal, error)
	// GetGoal retrieves a goal by id
	GetGoal(ctx context.Context, goalID uint64) (*schema.SmartGoal, error)
	// DeleteGoal deletes a goal by id
	DeleteGoal(ctx context.Context, goalID uint64) error
	// UpdateGoal updates a goal
	UpdateGoal(ctx context.Context, goal *schema.SmartGoal) error
	// GetGoalsByPocketID retrieves all goals by pocket id
	GetGoalsByPocketID(ctx context.Context, pocketID uint64) ([]*schema.SmartGoal, error)

	// Milestone Operations
	// CreateMilestone creates a milestone
	CreateMilestone(ctx context.Context, goalID uint64, milestone *schema.Milestone) (*schema.Milestone, error)
	// DeleteMilestone deletes a milestone by id
	DeleteMilestone(ctx context.Context, milestoneID uint64) error
	// UpdateMilestone updates a milestone
	UpdateMilestone(ctx context.Context, milestone *schema.Milestone) error
	// GetMilestone retrieves a milestone by id
	GetMilestone(ctx context.Context, milestoneID uint64) (*schema.Milestone, error)
	// GetMilestonesByGoalID retrieves all milestones by goal id
	GetMilestonesByGoalID(ctx context.Context, goalID uint64) ([]*schema.Milestone, error)

	// Forecast Operations
	// CreateForecast creates a forecast
	CreateForecast(ctx context.Context, goalID uint64, forecast *schema.Forecast) (*schema.Forecast, error)
	// GetForecast retrieves a forecast by id
	GetForecast(ctx context.Context, forecastID uint64) (*schema.Forecast, error)
	// UpdateForecast updates a forecast
	UpdateForecast(ctx context.Context, forecast *schema.Forecast) error
	// DeleteForecast deletes a forecast by id
	DeleteForecast(ctx context.Context, forecastID uint64) error

	// CreateBudget creates a budget
	CreateBudget(ctx context.Context, milestoneID uint64, budget *schema.Budget) (*schema.Budget, error)
	// GetBudget retrieves a budget by id
	GetBudget(ctx context.Context, budgetID uint64) (*schema.Budget, error)
	// UpdateBudget updates a budget
	UpdateBudget(ctx context.Context, budget *schema.Budget) error
	// DeleteBudget deletes a budget by id
	DeleteBudget(ctx context.Context, budgetID uint64) error
}

// Db withholds connection to a postgres database as well as a logging handler
type Db struct {
	// Conn serves as the actual database connection object
	Conn *core_database.DatabaseConn
	// Logger is the logging utility used by this object
	Logger *zap.Logger
	// MaxConnectionAttempts outlines the maximum connection attempts
	// to initiate against the database
	MaxConnectionAttempts int
	// MaxRetriesPerOperation defines the maximum retries to attempt per failed database
	// connection attempt
	MaxRetriesPerOperation int
	// RetryTimeOut defines the maximum time until a retry operation is observed as a
	// timed out operation
	RetryTimeOut time.Duration
	// OperationSleepInterval defines the amount of time between retry operations
	// that the system sleeps
	OperationSleepInterval time.Duration
	// Telemetry defines the object by which we will emit metrics, trace requests, and database operations
	Instrumentation *instrumentation.ServiceTelemetry
	// QueryOperator is the object that will be used to execute database queries
	QueryOperator *dal.Query
}

var _ DatabaseOperations = (*Db)(nil)

// ConnectionInitializationParams represents connection initialization parameters for the database
type ConnectionInitializationParams struct {
	// ConnectionParams outlines database connection parameters
	ConnectionParams *ConnectionParameters
	// Logger is the logging utility used by this object
	Logger *zap.Logger
	// MaxConnectionAttempts outlines the maximum connection attempts
	// to initiate against the database
	MaxConnectionAttempts int
	// MaxRetriesPerOperation defines the maximum retries to attempt per failed database
	// connection attempt
	MaxRetriesPerOperation int
	// RetryTimeOut defines the maximum time until a retry operation is observed as a
	// timed out operation
	RetryTimeOut time.Duration
	// RetrySleepInterval defines the amount of time between retry operations
	// that the system sleeps
	RetrySleepInterval time.Duration
	// Telemetry defines the object by which we will emit metrics, trace requests, and database operations
	Instrumentation *instrumentation.ServiceTelemetry
}

// New creates a database connection and returns the connection object
func New(ctx context.Context, params *ConnectionInitializationParams) (*Db,
	error) {
	errInvalidInputParams := service_errors.ErrInvalidInputParam

	// TODO: generate a span for the database connection attempt
	if params == nil || (params != nil && (params.ConnectionParams == nil || params.Logger == nil)) {
		return nil, errInvalidInputParams
	}

	logger := params.Logger
	databaseModels := schema.GetDatabaseSchemas()

	conn, err := connectToDatabase(ctx, params.ConnectionParams, params.Logger, databaseModels...)

	if err != nil {
		return nil, err
	}

	logger.Info("Successfully connected to the database")

	return &Db{
		Conn:                   conn,
		Logger:                 logger,
		MaxConnectionAttempts:  params.MaxConnectionAttempts,
		MaxRetriesPerOperation: params.MaxRetriesPerOperation,
		RetryTimeOut:           params.RetryTimeOut,
		OperationSleepInterval: params.RetrySleepInterval,
		Instrumentation:        params.Instrumentation,
		QueryOperator:          dal.Use(conn.Engine),
	}, nil
}
