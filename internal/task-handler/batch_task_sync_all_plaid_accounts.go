package taskhandler

import (
	"context"
	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
	encryptdecrypt "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/encrypt_decrypt"
	"go.uber.org/zap"
)

type SyncAllPlatformConnectedPlaidAccountsPayload struct{}

func (t *SyncAllPlatformConnectedPlaidAccountsPayload) String() *string {
	str := "SyncAllPlatformConnectedPlaidAccountsPayload"
	return &str
}

func NewSyncAllPlatformConnectedPlaidAccounts() (*asynq.Task, error) {
	payload, err := json.Marshal(&SyncPlaidTaskPayload{})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskSyncAllAccounts.String(), payload), nil
}

// RunSyncAllPlatformConnectedPlaidAccounts is a method on the TaskHandler struct that takes a context and a task
// as parameters. It is used to sync all accounts for all platform users. It does this by querying the database for all
// connected accounts, decrypting the access token, and calling the sync operation for each linked account a given user has.
func (th *TaskHandler) RunSyncAllPlatformConnectedPlaidAccounts(ctx context.Context, task *asynq.Task) error {
	var (
		payload SyncPlaidTaskPayload
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

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
				trigger := payload.String()

				// decrypt the access token
				accessToken, err := encryptdecrypt.DecryptUserAccessToken(ctx, link.Token, th.postgresDb.Kms, th.logger)
				if err != nil {
					th.logger.Error("failed to decrypt access token", zap.Error(err))
					continue
				}

				// call the sync operation for the link
				accts, err := th.processSyncOperation(ctx, profile.UserId, link.Id, *accessToken, *trigger)
				if err != nil {
					return err
				}

				// log a warning if no accounts were found
				if len(accts) == 0 {
					th.logger.Warn("no accounts found for user", zap.Uint64("user_id", payload.UserId))
				}
			}
		}
	}

	return nil
}
