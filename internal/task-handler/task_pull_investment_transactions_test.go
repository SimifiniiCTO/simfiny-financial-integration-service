package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

func TestNewPullInvestmentTransactionsTask(t *testing.T) {
	type args struct {
		userId      uint64
		linkId      uint64
		accessToken string
		accountIds  []string
	}
	tests := []struct {
		name    string
		args    args
		want    *asynq.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPullInvestmentTransactionsTask(tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.accountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPullInvestmentTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullInvestmentTransactionsTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_RunPullInvestmentTransactionsTask(t *testing.T) {
	type args struct {
		ctx context.Context
		t   *asynq.Task
	}
	tests := []struct {
		name    string
		th      *TaskHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.th.RunPullInvestmentTransactionsTask(tt.args.ctx, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunPullInvestmentTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isTransactionInSlice(t *testing.T) {
	type args struct {
		transaction  *schema.InvestmentTransaction
		transactions []*schema.InvestmentTransaction
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTransactionInSlice(tt.args.transaction, tt.args.transactions); got != tt.want {
				t.Errorf("isTransactionInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
