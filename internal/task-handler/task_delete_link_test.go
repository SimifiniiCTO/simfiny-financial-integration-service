package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestTaskDeleteLinkPayload_String(t *testing.T) {
	tests := []struct {
		name    string
		tr      *TaskDeleteLinkPayload
		wantErr bool
	}{
		{
			name: "delete link payload",
			tr: &TaskDeleteLinkPayload{
				UserId:          uint64(helper.GenerateRandomId(1000, 10000)),
				PlaidLinkItemId: helper.GenerateRandomString(10),
				AccessToken:     helper.GenerateRandomString(10),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.String()
			if !tt.wantErr {
				assert.NotNil(t, got)
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
