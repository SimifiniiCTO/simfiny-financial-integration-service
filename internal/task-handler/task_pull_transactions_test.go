package taskhandler

import (
	"context"
	"testing"
	"time"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandler_RunPullTransactionsTask(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		args         args
		th           *TaskHandler
		precondition func(t *testing.T, arg *args) *asynq.Task
		wantErr      bool
	}{
		{
			name: "runPullTransactionsTask",
			th:   testTaskHandler,
			precondition: func(t *testing.T, arg *args) *asynq.Task {
				// create a user and associate a link to said user
				profile := helper.GenereateRandomUserProfileForTest()
				profile.Link[0].PlaidLink.ItemId = testItemId
				profile.Link[0].Token.ItemId = testItemId
				userId := profile.UserId

				plaidBankAccounts, err := testTaskHandler.plaidClient.GetAccounts(arg.ctx, testAccessToken, userId)
				assert.NoError(t, err)

				profile.Link[0].BankAccounts = plaidBankAccounts

				user, err := testTaskHandler.postgresDb.CreateUserProfile(arg.ctx, profile)
				assert.NoError(t, err)

				linkId := user.Link[0].Id
				task, err := NewPullTransactionsTask(userId, linkId, testAccessToken, time.Now().AddDate(0, 0, -30), time.Now())
				assert.NoError(t, err)

				return task
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsk := tt.precondition(t, &tt.args)
			if err := tt.th.RunPullTransactionsTask(tt.args.ctx, tsk); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunPullTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
