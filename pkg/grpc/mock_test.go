package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"testing"
	"time"

	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
)

var (
	DbConnHandler *database.Db
	Port          int    = 5555
	Host          string = "localhost"
	User          string = "service_db"
	Password      string = "service_db"
	Dbname        string = "service_db"
	MockServer    *Server
)

type MockDialOption func(context.Context, string) (net.Conn, error)

// dialer creates an in memory full duplex connection
func dialer() func() MockDialOption {
	return func() MockDialOption {
		listener := bufconn.Listen(1024 * 1024)

		server := grpc.NewServer()
		schema.RegisterFinancialServiceServer(server, MockServer)

		go func() {
			if err := server.Serve(listener); err != nil {
				log.Fatal(err)
			}
		}()

		return func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}
	}
}

// MockGRPCService creates and returns a mock grpc service connection
func MockGRPCService(ctx context.Context) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// NewMockServer creates a new mock server instance
func NewMockServer(db *database.Db) {
	config := &Config{
		Port:            9897,
		ServiceName:     "FinancialIntegrationService",
		Environment:     "dev",
		NewRelicLicense: "62fd721c712d5863a4e75b8f547b7c1ea884NRAL",
		RpcTimeout:      3 * time.Minute,
	}

	/*plaidHandler, err := plaidhandler.GetPlaidWrapperForTest()
	if err != nil {
		log.Fatal(err)
	}*/

	client := plaidhandler.NewMockPlaidClient()
	var err error
	MockServer, err = NewServer(&Params{
		Config:          config,
		Logger:          zap.L(),
		Instrumentation: nil,
		Db:              DbConnHandler,
		PlaidClient:     client,
	})

	if err != nil {
		log.Fatal(err)
	}
}

// TestMain runs main function
func TestMain(m *testing.M) {
	// os.Exit skips defer calls
	// so we need to call another function
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

// run runs the suite of tests residing within the package
func run(m *testing.M) (code int, err error) {
	// pseudo-code, some implementation excluded:
	//
	// 1. create test.db if it does not exist
	// 2. run our DDL statements to create the required tables if they do not exist
	// 3. run our tests
	// 4. truncate the test db tables
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}

	testdb, err := gorm.Open(sqlite.Open("test-b"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("open sqlite %q fail: %w", "test-db", err))
	}

	testdb.AutoMigrate(schema.GetDatabaseSchemas()...)
	DbConnHandler = &database.Db{
		Conn: &core_database.DatabaseConn{
			Engine: testdb,
		},
		QueryOperator:          dal.Use(testdb),
		Logger:                 logger,
		MaxConnectionAttempts:  3,
		RetryTimeOut:           1 * time.Minute,
		OperationSleepInterval: 500 * time.Millisecond,
	}

	// iniitialize mock server
	NewMockServer(DbConnHandler)

	defer func() {
		// tears down a connection object to the test db node
		db, err := DbConnHandler.Conn.Engine.DB()
		if err != nil {
			os.Exit(1)
		}

		if err := db.Close(); err != nil {
			os.Exit(1)
		}
	}()

	code = m.Run()

	return code, nil
}

// stopConflictingProcesses stops any conflicting process running on a the defined port
func stopConflictingProcesses(port int) {
	if runtime.GOOS == "windows" {
		command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %d).OwningProcess -Force", port)
		exec_cmd(exec.Command("Stop-Process", "-Id", command))
	} else {
		command := fmt.Sprintf("lsof -i tcp:%d | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
		exec_cmd(exec.Command("bash", "-c", command))
	}
}

// Execute command and return exited code.
func exec_cmd(cmd *exec.Cmd) {
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			fmt.Printf("Error during killing (exit code: %s)\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Port successfully killed (exit code: %s)\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}
