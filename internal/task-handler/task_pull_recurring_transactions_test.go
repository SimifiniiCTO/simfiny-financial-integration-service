package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestNewPullUpdatedReCurringTransactionsTask(t *testing.T) {
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
			got, err := NewPullUpdatedReCurringTransactionsTask(tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.accountIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPullUpdatedReCurringTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullUpdatedReCurringTransactionsTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_RunPullUpdatedReCurringTransactionsTask(t *testing.T) {
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
			if err := tt.th.RunPullUpdatedReCurringTransactionsTask(tt.args.ctx, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunPullUpdatedReCurringTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_copyOldtxnToNewtxn(t *testing.T) {
	type args struct {
		oldtxn *schema.ReOccuringTransaction
		newtxn *schema.ReOccuringTransaction
	}
	tests := []struct {
		name string
		args args
		want *schema.ReOccuringTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyOldtxnToNewtxn(tt.args.oldtxn, tt.args.newtxn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyOldtxnToNewtxn() = %v, want %v", got, tt.want)
			}
		})
	}
}
