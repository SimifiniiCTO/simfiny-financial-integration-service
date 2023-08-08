package taskhandler

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/SimifiniiCTO/asynq"
)

type SyncPlaidTaskPayload struct {
	UserId      uint64 `json:"user_id"`
	AccessToken string `json:"access_token"`
	LinkId      uint64 `json:"plaid_link_item_id"`
}

func (t *SyncPlaidTaskPayload) String() *string {
	str := fmt.Sprintf("user_id: %d, access_token: %s, plaid_link_item_id: %d", t.UserId, t.AccessToken, t.LinkId)
	return &str
}

// This function creates a new asynchronous task for syncing Plaid  with the provided user
// ID and access token.
func NewSyncPlaidTask(userId uint64, accessToken string, linkId uint64) (*asynq.Task, error) {
	payload, err := json.Marshal(&SyncPlaidTaskPayload{
		UserId:      userId,
		AccessToken: accessToken,
		LinkId:      linkId,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskSyncTransactions.String(), payload), nil
}

// This is a method of the `TaskHandler` struct that handles the execution of a task for syncing Plaid
// transactions. It takes in a context and a task object as parameters and returns an error. Inside the
// method, it unmarshals the payload of the task, queries Plaid for the user's transactions, sanitizes
// and transforms the transactions, inserts them into a database, and sends the data to a data
// warehouse.
func (th *TaskHandler) RunSyncPlaidTransactionsTask(ctx context.Context, task *asynq.Task) error {
	var (
		payload SyncPlaidTaskPayload
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	trigger := payload.String()
	err := th.processSyncOperation(ctx, payload.UserId, payload.LinkId, payload.AccessToken, *trigger)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Refactor this piece of code
func (th *TaskHandler) processSyncOperation(ctx context.Context, userId, linkId uint64, accessToken, trigger string) error {
	link, err := th.postgresDb.GetLink(ctx, userId, linkId, false)
	if err != nil {
		return fmt.Errorf("link with id %d does not exist", linkId)
	}

	// plaidAccountIds := make([]string, 0, len(bankAccounts))
	// for _, plaidBankAccount := range bankAccounts {
	// 	if plaidBankAccount.Subtype == "checking" || plaidBankAccount.Subtype == "savings" || plaidBankAccount.Subtype == "cd" || plaidBankAccount.Subtype == "money market" ||
	// 		plaidBankAccount.Subtype == "gic" || plaidBankAccount.Subtype == "prepaid" || plaidBankAccount.Subtype == "health" || plaidBankAccount.Subtype == "cash management" {
	// 		plaidAccountIds = append(plaidAccountIds, plaidBankAccount.PlaidAccountId)
	// 	}
	// }

	th.syncAllAccounts(ctx, userId, link, &accessToken, &trigger)

	return nil
}
