package taskhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/SimifiniiCTO/asynq"
	"go.uber.org/zap"
)

type TaskDeleteTransactionsPayload struct {
	UserId         uint64   `json:"user_id"`
	TransactionIds []uint64 `json:"transaction_ids"`
	LinkId         uint64   `json:"link_id"`
}

func (t *TaskDeleteTransactionsPayload) String() *string {
	str := fmt.Sprintf("userId: %d, TransactionIds: %v", t.UserId, t.TransactionIds)
	return &str
}

func NewDeleteTransactionsTask(userId, linkId uint64, transactionIds []uint64) (*asynq.Task, error) {
	rec := &TaskDeleteTransactionsPayload{
		UserId:         userId,
		TransactionIds: transactionIds,
		LinkId:         linkId,
	}

	payload, err := json.Marshal(rec)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskDeleteTransactions.String(), payload), nil
}

func (th *TaskHandler) RunDeleteTransactionsTask(ctx context.Context, task *asynq.Task) error {
	var (
		payload          TaskDeleteTransactionsPayload
		postgresClient   = th.postgresDb
		clickhouseClient = th.clickhouseDb
	)

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	if len(payload.TransactionIds) == 0 {
		th.logger.Error("no transaction ids provided", zap.Uint64("link_id", payload.LinkId))
		return nil
	}

	// query the link from the database
	link, err := postgresClient.GetLink(ctx, payload.UserId, payload.LinkId)
	if err != nil {
		return err
	}

	if link.PlaidLink == nil {
		th.logger.Error("provided link does not have any plaid credentials", zap.Uint64("link_id", payload.LinkId))
		return nil
	}

	// delete the transactions from the database
	if err := clickhouseClient.DeleteTransactionsByIds(ctx, payload.TransactionIds); err != nil {
		return err
	}

	link.LastSuccessfulUpdate = time.Now().String()
	if err := postgresClient.UpdateLink(ctx, link); err != nil {
		return err
	}

	return nil
}
