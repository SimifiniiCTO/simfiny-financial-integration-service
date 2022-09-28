package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/service_errors"
	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ConnectionParameters struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         int
	SslMode      string
}

// connectToDatabase establish and connects to a database instance
func connectToDatabase(ctx context.Context, params *ConnectionParameters, log *zap.Logger, models ...interface{}) (*core_database.DatabaseConn,
	error) {
	var dbConn *core_database.DatabaseConn
	connectionString := configureConnectionString(ctx, params.Host, params.User, params.Password, params.DatabaseName, params.Port, params.SslMode)

	log.Info(connectionString)

	if dbConn = core_database.NewDatabaseConn(connectionString, "postgres"); dbConn == nil {
		return nil, errors.New("failed to connect to merchant component database")
	}

	if err := pingDatabase(ctx, dbConn); err != nil {
		return nil, err
	}

	if err := migrateSchemas(ctx, dbConn, log, models...); err != nil {
		return nil, err
	}

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
func configureConnectionString(ctx context.Context, host, user, password, dbname string, port int, sslMode string) string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)
	return connectionString
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
