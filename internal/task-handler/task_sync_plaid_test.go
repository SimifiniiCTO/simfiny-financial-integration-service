package taskhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	apiv1 "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandler_RunSyncPlaidTransactionsTask(t *testing.T) {
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
			name: "runSyncPlaidTransactionsTask",
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
				task, err := NewSyncPlaidTask(userId, testAccessToken, linkId)
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
			if err := tt.th.RunSyncPlaidTransactionsTask(tt.args.ctx, tsk); (err != nil) != tt.wantErr {
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
