package database

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
)

var Conn *Db

type NewDBConnectionTestMetadata struct {
	shouldErrorOccur bool
	connectionParams *ConnectionInitializationParams
}

// setup sets up a database connection to the test db node
func setup() {
	ctx := context.Background()
	telemetry, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("test-service"),
		newrelic.ConfigLicense("62fd721c712d5863a4e75b8f547b7c1ea884NRAL"),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = true
		},
		// Use nrzap to register the logger with the agent:
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(false),
	)
	DefaultConnInitializationParams.Telemetry = telemetry
	// setup database connection before tests
	Conn, _ = New(ctx, &DefaultConnInitializationParams)
}

// teardown tears down a connection object to the test db node
func teardown() {
	db, err := Conn.Conn.Engine.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

// reset a database connection
func reset() {
	teardown()
	setup()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// TestNewDbConnection tests connections attempts to database node
func TestNewDbConnection(t *testing.T) {
	ctx := context.Background()
	var scenarios = getNewDbConnectionTestScenarios()

	for _, scenario := range scenarios {
		conn, err := New(ctx, scenario.connectionParams)
		if err != nil && !scenario.shouldErrorOccur {
			t.Errorf("obtained error but not expected - %s", err.Error())
		}

		if err == nil && scenario.shouldErrorOccur {
			t.Errorf("failed to obtain expected error - %s", err.Error())
		}

		if !scenario.shouldErrorOccur {
			assert.NotNil(t, conn)
		}
	}
}

// getNewDbConnectionTestScenarios returns a set of test scenarios to test creation of new database connections
func getNewDbConnectionTestScenarios() []NewDBConnectionTestMetadata {
	return []NewDBConnectionTestMetadata{
		{
			shouldErrorOccur: false,
			connectionParams: &DefaultConnInitializationParams,
		},
		{
			true,
			nil,
		},
	}
}
