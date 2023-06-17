package taskhandler

import (
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	clickhousedb "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	postgresdb "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"
	"go.uber.org/zap"
)

func TestTaskType_String(t *testing.T) {
	tests := []struct {
		name string
		tr   TaskType
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("TaskType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLogger(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPostgresDb(t *testing.T) {
	type args struct {
		postgresDb *postgresdb.Db
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPostgresDb(tt.args.postgresDb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPostgresDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithClickhouseDb(t *testing.T) {
	type args struct {
		clickhouseDb *clickhousedb.Db
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithClickhouseDb(tt.args.clickhouseDb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithClickhouseDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithInstrumentationClient(t *testing.T) {
	type args struct {
		instrumentationClient *instrumentation.Client
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithInstrumentationClient(tt.args.instrumentationClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithInstrumentationClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPlaidClient(t *testing.T) {
	type args struct {
		plaidClient *plaidhandler.PlaidWrapper
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPlaidClient(tt.args.plaidClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPlaidClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTaskHandler(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *TaskHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskHandler(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_Validate(t *testing.T) {
	tests := []struct {
		name    string
		th      *TaskHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.th.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskHandler_RegisterTaskHandler(t *testing.T) {
	tests := []struct {
		name string
		th   *TaskHandler
		want *asynq.ServeMux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.th.RegisterTaskHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskHandler.RegisterTaskHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_instrumentationMiddleware(t *testing.T) {
	type args struct {
		next asynq.Handler
	}
	tests := []struct {
		name string
		th   *TaskHandler
		args args
		want asynq.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.th.instrumentationMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskHandler.instrumentationMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
