package taskhandler

import (
	"context"
	"testing"
	"time"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandler_RunDeleteTransactionsTask(t *testing.T) {
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
		// TODO: Add test cases.
		{
			name: "runDeleteTransactionsTask",
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

				// pull transactions
				bankAccountIds := make([]string, 0)
				for _, acct := range user.Link[0].BankAccounts {
					bankAccountIds = append(bankAccountIds, acct.PlaidAccountId)
				}

				txs, err := testTaskHandler.plaidClient.GetTransactions(arg.ctx, testAccessToken, time.Now().AddDate(0, 0, -5), time.Now(), 10, 1, bankAccountIds)
				assert.NoError(t, err)

				// save the transactions to the database
				err = testTaskHandler.clickhouseDb.AddTransactions(arg.ctx, &userId, txs)
				assert.NoError(t, err)

				// get the transaction ids
				createdTxs, _, err := testTaskHandler.clickhouseDb.GetTransactions(arg.ctx, &userId, 1, 10)
				assert.NoError(t, err)

				txsIds := make([]string, 0)
				for _, tx := range createdTxs {
					txsIds = append(txsIds, tx.Id)
				}

				task, err := NewDeleteTransactionsTask(userId, linkId, txsIds)
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
			if err := tt.th.RunDeleteTransactionsTask(tt.args.ctx, tsk); (err != nil) != tt.wantErr {
				t.Errorf("TaskHandler.RunDeleteTransactionsTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
