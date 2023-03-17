package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	newrelic "github.com/newrelic/go-agent"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	core_database "github.com/SimifiniiCTO/core/core-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/service_errors"
)

type ConnectionParameters struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         int
	SslMode      string
}

// startDatastoreSpan generates a datastore span
func (db *Db) startDatastoreSpan(ctx context.Context, name string) *newrelic.DatastoreSegment {
	if db.Instrumentation != nil {
		txn := db.Instrumentation.GetTraceFromContext(ctx)
		span := db.Instrumentation.StartDatastoreSegment(txn, name)
		return span
	}

	return nil
}

// connectToDatabase establish and connects to a database instance
func connectToDatabase(ctx context.Context, params *ConnectionInitializationParams, log *zap.Logger, models ...interface{}) (*core_database.DatabaseConn, error) {

	var (
		dbConn     *core_database.DatabaseConn
		dbConnStr  *string
		err        error
		connParams = params.ConnectionParams
	)

	if dbConnStr, err = formatDbConnectionString(ctx, connParams); err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("connecting to database: %s", *dbConnStr))

	dbConn = core_database.NewDatabaseConn(
		&core_database.Parameters{
			QueryTimeout:              params.QueryTimeout,           // REFACTOR into env var.
			MaxConnectionRetries:      &params.MaxConnectionAttempts, // REFACTOR into env var.
			MaxConnectionRetryTimeout: &params.RetryTimeOut,          // REFACTOR into env var.
			RetrySleep:                &params.RetrySleepInterval,    // REFACTOR into env var.
			ConnectionString:          dbConnStr,
		})

	if dbConn == nil {
		return nil, errors.New("failed to connect to merchant component database")
	}

	log.Info("database connection established now")

	if err := pingDatabase(ctx, dbConn); err != nil {
		return nil, err
	}

	if err := migrateSchemas(ctx, dbConn, log, models...); err != nil {
		return nil, err
	}

	dbConn.Engine.Preload(clause.Associations)
	db, err := dbConn.Engine.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(5 * time.Minute)

	return dbConn, nil
}

// pingDatabase pings a database node via a connection object
func pingDatabase(ctx context.Context, dbConn *core_database.DatabaseConn) error {
	db, err := dbConn.Engine.DB()
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	setConnectionConfigs(ctx, dbConn)
	return nil
}

// configureConnectionString constructs database connection string from a set of params
func formatDbConnectionString(ctx context.Context, params *ConnectionParameters) (*string, error) {
	if params == nil {
		return nil, fmt.Errorf("invalid input argument. params cannot be nil")
	}

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		params.Host, params.Port, params.User, params.Password, params.DatabaseName, params.SslMode)

	return &connectionStr, nil
}

// configureDatabaseConnection configures a database connection
func setConnectionConfigs(ctx context.Context, dbConn *core_database.DatabaseConn) {
	dbConn.Engine.FullSaveAssociations = true
	dbConn.Engine.SkipDefaultTransaction = false
	dbConn.Engine.PrepareStmt = true
	dbConn.Engine.DisableAutomaticPing = false
	dbConn.Engine = dbConn.Engine.Set("gorm:auto_preload", true)
}

// migrateSchemas creates or updates a given set of model based on a schema
// if it does not exist or migrates the model schemas to the latest version
func migrateSchemas(ctx context.Context, db *core_database.DatabaseConn, log *zap.Logger, models ...interface{}) error {
	var engine *gorm.DB
	if db == nil {
		return service_errors.ErrInvalidAcctParam
	}

	if engine = db.Engine; engine == nil {
		return service_errors.ErrInvalidGormDbOject
	}

	if len(models) > 0 {
		if err := engine.AutoMigrate(models...); err != nil {
			// TODO: emit metric
			log.Error(err.Error())
			return err
		}

		log.Info("successfully migrated database schemas")
	}

	return nil
}

func ConfigureConnectionString(host, user, password, databaseName, sslMode string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, databaseName, sslMode)
}
