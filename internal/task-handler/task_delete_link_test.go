package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
)

func TestTaskDeleteLinkPayload_String(t *testing.T) {
	tests := []struct {
		name string
		tr   *TaskDeleteLinkPayload
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskDeleteLinkPayload.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeleteLinkTask(t *testing.T) {
	type args struct {
		userId          uint64
		plaidLinkItemId string
		accessToke      string
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
			got, err := NewDeleteLinkTask(tt.args.userId, tt.args.plaidLinkItemId, tt.args.accessToke)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDeleteLinkTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteLinkTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_RunDeleteLinkTask(t *testing.T) {
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
			if err := tt.th.RunDeleteLinkTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunDeleteLinkTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
