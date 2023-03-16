package database

import (
	"fmt"
	"sync"
	"time"

	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/dal"
)

const dbName = "gen_test.db"

var testdb *gorm.DB
var once sync.Once
var conn *Db

func init() {
	conn = NewTestDatabase()
}

func NewTestDatabase() *Db {
	InitializeDB()
	testdb.AutoMigrate(schema.GetDatabaseSchemas()...)
	return &Db{
		Conn: &core_database.DatabaseConn{
			Engine: testdb,
		},
		QueryOperator:          dal.Use(testdb),
		Logger:                 zap.L(),
		MaxConnectionAttempts:  3,
		RetryTimeOut:           1 * time.Minute,
		OperationSleepInterval: 500 * time.Millisecond,
	}
}

func InitializeDB() {
	once.Do(func() {
		var err error
		testdb, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
	})
}
