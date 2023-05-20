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

	clickhouseDatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	"github.com/alicebob/miniredis/v2"

	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/dal"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	database "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"

	redisDatabase "github.com/SimifiniiCTO/simfiny-core-lib/database/redis"
)

var (
	DbConnHandler   *database.Db
	Port            int    = 5555
	Host            string = "localhost"
	User            string = "service_db"
	Password        string = "service_db"
	Dbname          string = "service_db"
	MockServer      *Server
	redisTestServer *miniredis.Miniredis
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

func MockRedis() *redisDatabase.Client {
	redisTestServer = mockRedis()
	// Set up test stop channel
	stopCh := make(chan struct{})

	// Set up test options
	logger := zap.NewNop()
	opts := []redisDatabase.Option{
		redisDatabase.WithLogger(logger),
		redisDatabase.WithURI(getRedisConnectionString()),
		redisDatabase.WithCacheTTLInSeconds(60),
		redisDatabase.WithServiceName("test-service"),
		redisDatabase.WithTelemetrySdk(&instrumentation.Client{}),
	}

	// Call New() to create a new Redis client
	client, err := redisDatabase.New(stopCh, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// NewMockServer creates a new mock server instance
func NewMockServer(db *database.Db) {
	config := &Config{
		Port:                     9897,
		GatewayPort:              Port,
		ServiceName:              "FinancialIntegrationService",
		NewRelicLicense:          "62fd721c712d5863a4e75b8f547b7c1ea884NRAL",
		Environment:              "dev",
		RpcTimeout:               30 * time.Minute,
		StripeApiKey:             "sk_test_51M1F1pBV97V9M33e3Ki1k5OqkdhfdDUBNTwDFzUtRmsSYbHf7qE3d1kkFCYRxfS70bJKBOKR5Zbv103sqvNd0gnm00lMyRDWEh",
		PlaidClientID:            "",
		PlaidSecretKey:           "",
		PlaidEnv:                 "",
		PlaidOauthDomain:         "",
		PlaidWebhooksEnabled:     false,
		PlaidWebhookOauthDomain:  "",
		AwsAccessKeyID:           "",
		AwsRegion:                "",
		AwsSecretAccessKey:       "",
		AwsKmsKeyID:              "",
		MaxPlaidLinks:            0,
		BillingEnabled:           false,
		WorkflowExecutionTimeout: 0,
		WorkflowTaskTimeout:      0,
		WorkflowRunTimeout:       0,
		TaskProcessorWorkers:     3,
	}

	handler, err := plaidhandler.GetPlaidWrapperForTest()
	if err != nil {
		log.Fatal(err)
	}

	clickhouseDb := clickhouseDatabase.NewMockInMemoryClickhouseDB()

	redisDb := MockRedis()

	kms, err := secrets.NewMockAwsKms()
	if err != nil {
		log.Fatal(err)
	}

	MockServer, err = NewServer(&Params{
		Config:             config,
		Logger:             zap.L(),
		Instrumentation:    &instrumentation.Client{},
		Db:                 DbConnHandler,
		PlaidWrapper:       handler,
		TransactionManager: nil,
		ClickhouseDb:       clickhouseDb,
		RedisDb:            redisDb,
		KeyManagement:      kms,
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

	testdb, err := postgresdb.NewInMemoryTestDbClient(schema.GetDatabaseSchemas()...)
	if err != nil {
		panic(fmt.Errorf("failed to create in memory test db client: %w", err))
	}

	testdb.Engine.AutoMigrate(schema.GetDatabaseSchemas()...)
	DbConnHandler = &database.Db{
		Conn:                   testdb,
		QueryOperator:          dal.Use(testdb.Engine),
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

// mockRedis returns a mock redis server.
func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	return s
}

func getRedisConnectionString() string {
	return fmt.Sprintf("redis://:@%s", redisTestServer.Addr())
}
