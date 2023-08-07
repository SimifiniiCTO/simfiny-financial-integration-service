package taskhandler

import (
	"context"
	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
	encryptdecrypt "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/encrypt_decrypt"
	"go.uber.org/zap"
)

type SyncRecurringTransactionsPayload struct{}

func (t *SyncRecurringTransactionsPayload) String() *string {
	str := "SyncRecurringTransactionsPayload"
	return &str
}

func NewSyncRecurringTransactions() (*asynq.Task, error) {
	payload, err := json.Marshal(&SyncPlaidTaskPayload{})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskBatchSyncRecurringTransactions.String(), payload), nil
}

// RunSyncRecurringTransactions is a method on the TaskHandler struct that takes a context and a task
// as parameters. It is used to sync all accounts for all platform users. It does this by querying the database for all
// connected accounts, decrypting the access token, and calling the sync operation for each linked account a given user has.
func (th *TaskHandler) RunBatchSyncRecurringTransactions(ctx context.Context, task *asynq.Task) error {
	// query the database for all connected accounts
	profiles, err := th.postgresDb.GetAllUserProfiles(ctx)
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		for _, link := range profile.Link {
			// TODO: might need to be more intelligent about how we perform this as it will take a long time
			// Just monitor this as much as possible
			if link.PlaidLink != nil {
				bankAccounts := link.BankAccounts
				// get the bank account ids
				bankAccountIds := make([]string, 0, len(bankAccounts))
				for _, bankAccount := range bankAccounts {
					bankAccountIds = append(bankAccountIds, bankAccount.PlaidAccountId)
				}

				if len(bankAccountIds) == 0 {
					continue
				}

				// decrypt the access token
				accessToken, err := encryptdecrypt.DecryptUserAccessToken(ctx, link.Token, th.postgresDb.Kms, th.logger)
				if err != nil {
					th.logger.Error("failed to decrypt access token", zap.Error(err))
					continue
				}

				// call the sync operation for the link
				if err := th.
					queryAndStoreRecurringTransactions(
						ctx,
						*accessToken,
						profile.UserId,
						link.Id,
						bankAccountIds); err != nil {
					th.logger.Error("failed to query recurring transactions", zap.Error(err))
					continue
				}
			}
		}
	}

	return nil
}
