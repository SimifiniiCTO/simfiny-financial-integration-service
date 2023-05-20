package taskhandler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SimifiniiCTO/asynq"
	"go.uber.org/zap"
)

type TaskDeleteLinkPayload struct {
	UserId          uint64 `json:"user_id"`
	PlaidLinkItemId string `json:"plaid_link_item_id"`
	AccessToken     string `json:"access_token"`
}

func (t *TaskDeleteLinkPayload) String() *string {
	str := fmt.Sprintf("userId: %d, PlaidLinkItemId: %s", t.UserId, t.PlaidLinkItemId)
	return &str
}

func NewDeleteLinkTask(userId uint64, plaidLinkItemId, accessToke string) (*asynq.Task, error) {
	rec := &TaskDeleteLinkPayload{
		UserId:          userId,
		PlaidLinkItemId: plaidLinkItemId,
		AccessToken:     accessToke,
	}

	payload, err := json.Marshal(rec)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskDeleteLink.String(), payload), nil
}

func (th *TaskHandler) RunDeleteLinkTask(ctx context.Context, task *asynq.Task) error {
	var (
		payload          TaskDeleteLinkPayload
		postgresClient   = th.postgresDb
		clickhouseClient = th.clickhouseDb
		plaidClient      = th.plaidClient
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	// query the link from the database
	link, err := postgresClient.GetLinkByItemId(ctx, payload.PlaidLinkItemId)
	if err != nil {
		return err
	}

	if link.PlaidLink == nil {
		th.logger.Error("provided link does not have any plaid credentials", zap.String("plaid_link_item_id", payload.PlaidLinkItemId))
		return nil
	}

	// Delete link and all associated references
	if err := postgresClient.DeleteLink(ctx, payload.UserId, link.Id); err != nil {
		return err
	}

	// now deactivate the plaid item
	if err := plaidClient.DeleteItem(ctx, &payload.AccessToken); err != nil {
		return err
	}

	// delete all transactions associated with the link
	if err := clickhouseClient.DeleteTransactionsByLinkId(ctx, &link.Id); err != nil {
		return err
	}

	return nil
}
