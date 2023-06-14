package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
)

func TestTaskDeleteTransactionsPayload_String(t *testing.T) {
	tests := []struct {
		name string
		tr   *TaskDeleteTransactionsPayload
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskDeleteTransactionsPayload.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeleteTransactionsTask(t *testing.T) {
	type args struct {
		userId         uint64
		linkId         uint64
		transactionIds []uint64
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
			got, err := NewDeleteTransactionsTask(tt.args.userId, tt.args.linkId, tt.args.transactionIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDeleteTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteTransactionsTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_RunDeleteTransactionsTask(t *testing.T) {
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
			if err := tt.th.RunDeleteTransactionsTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunDeleteTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
