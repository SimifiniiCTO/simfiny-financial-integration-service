package taskhandler

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/SimifiniiCTO/asynq"
)

func TestNewPullTransactionsTask(t *testing.T) {
	type args struct {
		userId      uint64
		linkId      uint64
		accessToken string
		startTime   time.Time
		endTime     time.Time
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
			got, err := NewPullTransactionsTask(tt.args.userId, tt.args.linkId, tt.args.accessToken, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPullTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPullTransactionsTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_RunPullTransactionsTask(t *testing.T) {
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
			if err := tt.th.RunPullTransactionsTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunPullTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
