package database

import (
	"context"
	"reflect"
	"testing"
	"time"

	postgresdb "github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/dal"
	"go.uber.org/zap"
)

func TestDb_GetCreditAccount(t *testing.T) {
	type fields struct {
		Conn                   *postgresdb.Client
		Logger                 *zap.Logger
		MaxConnectionAttempts  int
		MaxRetriesPerOperation int
		RetryTimeOut           time.Duration
		OperationSleepInterval time.Duration
		Instrumentation        *instrumentation.Client
		QueryOperator          *dal.Query
		Kms                    secrets.KeyManagement
	}
	type args struct {
		ctx             context.Context
		creditAccountId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *schema.CreditAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Db{
				Conn:                   tt.fields.Conn,
				Logger:                 tt.fields.Logger,
				MaxConnectionAttempts:  tt.fields.MaxConnectionAttempts,
				MaxRetriesPerOperation: tt.fields.MaxRetriesPerOperation,
				RetryTimeOut:           tt.fields.RetryTimeOut,
				OperationSleepInterval: tt.fields.OperationSleepInterval,
				Instrumentation:        tt.fields.Instrumentation,
				QueryOperator:          tt.fields.QueryOperator,
				Kms:                    tt.fields.Kms,
			}
			got, err := db.GetCreditAccount(tt.args.ctx, tt.args.creditAccountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetCreditAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetCreditAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetSudentLoanAccount(t *testing.T) {
	type fields struct {
		Conn                   *postgresdb.Client
		Logger                 *zap.Logger
		MaxConnectionAttempts  int
		MaxRetriesPerOperation int
		RetryTimeOut           time.Duration
		OperationSleepInterval time.Duration
		Instrumentation        *instrumentation.Client
		QueryOperator          *dal.Query
		Kms                    secrets.KeyManagement
	}
	type args struct {
		ctx                  context.Context
		studentLoanAccountId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *schema.StudentLoanAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Db{
				Conn:                   tt.fields.Conn,
				Logger:                 tt.fields.Logger,
				MaxConnectionAttempts:  tt.fields.MaxConnectionAttempts,
				MaxRetriesPerOperation: tt.fields.MaxRetriesPerOperation,
				RetryTimeOut:           tt.fields.RetryTimeOut,
				OperationSleepInterval: tt.fields.OperationSleepInterval,
				Instrumentation:        tt.fields.Instrumentation,
				QueryOperator:          tt.fields.QueryOperator,
				Kms:                    tt.fields.Kms,
			}
			got, err := db.GetSudentLoanAccount(tt.args.ctx, tt.args.studentLoanAccountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetSudentLoanAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetSudentLoanAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
