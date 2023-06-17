package plaidhandler

import (
	"context"

	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
)

type SyncResult struct {
	NextCursor string
	HasMore    bool
	New        []*schema.Transaction
	Updated    []*schema.Transaction
	Deleted    []string
}

func (p *PlaidWrapper) Sync(ctx context.Context, cursor, accessToken *string) (*SyncResult, error) {
	request := p.client.PlaidApi.
		TransactionsSync(ctx).
		TransactionsSyncRequest(plaid.TransactionsSyncRequest{
			AccessToken: *accessToken,
			Cursor:      cursor,
			Count:       pointer.Int32P(500),
		})

	result, _, err := request.Execute()
	if err != nil {
		p.Logger.Error("failed to sync data with Plaid", zap.Error(err))
		return nil, err
	}

	added := make([]*schema.Transaction, 0, len(result.Added))
	for _, transaction := range result.Added {
		newTx, err := transformer.NewTransactionFromPlaid(&transaction)
		if err != nil {
			return nil, err
		}

		added = append(added, newTx)
	}

	modified := make([]*schema.Transaction, 0, len(result.Modified))
	for _, transaction := range result.Modified {
		modifiedTx, err := transformer.NewTransactionFromPlaid(&transaction)
		if err != nil {
			return nil, err
		}

		modified = append(modified, modifiedTx)
	}

	removed := make([]string, len(result.Removed))
	for _, transaction := range result.Removed {
		removed = append(removed, transaction.GetTransactionId())
	}

	if len(added)+len(modified)+len(removed) == 0 {
		p.Logger.Info("no changes observed from Plaid via sync")
	} else {
		p.Logger.Info("received changes from Plaid via sync", zap.Any("added", len(added)), zap.Any("modified", len(modified)), zap.Any("removed", len(removed)))
	}

	return &SyncResult{
		NextCursor: result.GetNextCursor(),
		HasMore:    result.GetHasMore(),
		New:        added,
		Updated:    modified,
		Deleted:    removed,
	}, nil
}
