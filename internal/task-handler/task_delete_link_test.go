package taskhandler

import (
	"context"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandler_RunDeleteLinkTask(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		th           *TaskHandler
		precondition func(t *testing.T, arg *args) *asynq.Task
		args         args
		wantErr      bool
	}{
		{
			name: "runDeleteLinkTask",
			th:   testTaskHandler,
			precondition: func(t *testing.T, arg *args) *asynq.Task {
				// create a user and associate a link to said user
				profile := helper.GenereateRandomUserProfileForTest()
				profile.Link[0].PlaidLink.ItemId = testItemId
				profile.Link[0].Token.ItemId = testItemId
				userId := profile.UserId

				user, err := testTaskHandler.postgresDb.CreateUserProfile(arg.ctx, profile)
				assert.NoError(t, err)

				link := user.Link[0]
				task, err := NewDeleteLinkTask(userId, link.Token.ItemId, testAccessToken)
				assert.NoError(t, err)

				return task
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsk := tt.precondition(t, &tt.args)

			if err := tt.th.RunDeleteLinkTask(tt.args.ctx, tsk); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunDeleteLinkTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
