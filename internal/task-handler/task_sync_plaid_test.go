package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestSyncPlaidTaskPayload_String(t *testing.T) {
	tests := []struct {
		name string
		tr   *SyncPlaidTaskPayload
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SyncPlaidTaskPayload.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSyncPlaidTask(t *testing.T) {
	type args struct {
		userId          uint64
		accessToken     string
		plaidLinkItemId uint64
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
			got, err := NewSyncPlaidTask(tt.args.userId, tt.args.accessToken, tt.args.plaidLinkItemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSyncPlaidTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSyncPlaidTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_RunSyncPlaidTransactionsTask(t *testing.T) {
	type args struct {
		ctx  context.Context
		task *asynq.Task
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
			if err := tt.th.RunSyncPlaidTransactionsTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunSyncPlaidTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskHandler_pullReOcurringTransactions(t *testing.T) {
	type args struct {
		ctx         context.Context
		userId      *uint64
		linkId      *uint64
		accessToken *string
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
			if err := tt.th.pullReOcurringTransactions(tt.args.ctx, tt.args.userId, tt.args.linkId, tt.args.accessToken); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.pullReOcurringTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskHandler_processSyncOperation(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *SyncPlaidTaskPayload
	}
	tests := []struct {
		name    string
		th      *TaskHandler
		args    args
		want    []*apiv1.BankAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.th.processSyncOperation(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.processSyncOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskHandler.processSyncOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
