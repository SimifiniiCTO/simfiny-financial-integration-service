package taskhandler

import (
	"context"
	"testing"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandler_RunSyncNewLiabilityAccountsTask(t *testing.T) {
	type args struct {
		// userId      uint64
		// accessToken string
		// linkId      uint64
		// accountIds  []string
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
			name: "runSyncNewLiabilityAccountsTask",
			th:   testTaskHandler,
			precondition: func(t *testing.T, arg *args) *asynq.Task {
				// create a user and associate a link to said user
				profile := helper.GenereateRandomUserProfileForTest()
				profile.Link[0].PlaidLink.ItemId = testItemId
				profile.Link[0].Token.ItemId = testItemId
				userId := profile.UserId

				plaidBankAccounts, err := testTaskHandler.plaidClient.GetAccounts(arg.ctx, testAccessToken, userId)
				assert.NoError(t, err)

				liabilityAcctSet, err := testTaskHandler.plaidClient.GetPlaidLiabilityAccounts(arg.ctx, &testAccessToken)
				assert.NoError(t, err)

				profile.Link[0].BankAccounts = plaidBankAccounts
				profile.Link[0].CreditAccounts = liabilityAcctSet.CrediCardAccounts
				profile.Link[0].StudentLoanAccounts = liabilityAcctSet.StudentLoanAccts
				profile.Link[0].MortgageAccounts = liabilityAcctSet.MortgageLoanAccts

				user, err := testTaskHandler.postgresDb.CreateUserProfile(arg.ctx, profile)
				assert.NoError(t, err)

				acctIdsOfInterest := make([]string, 0)
				for _, acct := range user.Link[0].CreditAccounts {
					acctIdsOfInterest = append(acctIdsOfInterest, acct.PlaidAccountId)
				}

				for _, acct := range user.Link[0].StudentLoanAccounts {
					acctIdsOfInterest = append(acctIdsOfInterest, acct.PlaidAccountId)
				}

				linkId := user.Link[0].Id
				task, err := NewSyncNewLiabilityAccountsTask(userId, testAccessToken, linkId, acctIdsOfInterest)
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
			if err := tt.th.RunSyncNewLiabilityAccountsTask(tt.args.ctx, tsk); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunSyncNewLiabilityAccountsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
