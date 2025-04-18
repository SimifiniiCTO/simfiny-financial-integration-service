package plaidhandler

import (
	"context"

	"github.com/plaid/plaid-go/v14/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// The SyncResult type contains information about the result of a synchronization operation, including
// new, updated, and deleted transactions, as well as a cursor for the next page of results and a flag
// indicating whether there are more results available.
// @property {string} NextCursor - NextCursor is a string that represents the cursor for the next page
// of results. It is used for pagination purposes, allowing the client to retrieve the next set of
// results from the server.
// @property {bool} HasMore - HasMore is a boolean property that indicates whether there are more
// transactions available to be fetched or not. If HasMore is true, it means that there are more
// transactions available and the client can make another request with the NextCursor value to fetch
// the next set of transactions. If HasMore is false,
// @property {[]*schema.Transaction} New - New is a slice of pointers to schema.Transaction objects
// that represent newly created transactions.
// @property {[]*schema.Transaction} Updated - The "Updated" property is a slice of pointers to
// "Transaction" objects that represent transactions that have been updated in a data synchronization
// process.
// @property {[]string} Deleted - The `Deleted` property is a slice of strings that contains the IDs of
// the transactions that were deleted in the synchronization process.
type SyncResult struct {
	NextCursor               string
	HasMore                  bool
	New                      []*schema.Transaction
	Updated                  []*schema.Transaction
	Deleted                  []string
	ItemLoginRequiredForLink bool
}

// Sync is used to sync transactions from the Plaid API for a given set of account IDs
func (p *PlaidWrapper) Sync(ctx context.Context, cursor, accessToken *string) (*SyncResult, error) {
	includePersonalFinanceCategory := true

	added := make([]*schema.Transaction, 0)
	modified := make([]*schema.Transaction, 0)
	removed := make([]string, 0)
	linkShouldBeInUpdateMode := false

	for {
		reqCtx := plaid.TransactionsSyncRequest{
			AccessToken: *accessToken,
			// make sure this is the max count ... anything higher will cause the write operations against an account to fail due to timeout
			Count: pointer.Int32P(500),
			Options: &plaid.TransactionsSyncRequestOptions{
				IncludePersonalFinanceCategory: &includePersonalFinanceCategory,
			},
		}

		if cursor != nil {
			reqCtx.Cursor = cursor
		}

		// query the Plaid API for transactions
		request := p.client.PlaidApi.
			TransactionsSync(ctx).
			TransactionsSyncRequest(reqCtx)

		result, _, err := request.Execute()

		if err != nil {
			if plaidErr, currerr := plaid.ToPlaidError(err); currerr == nil {
				p.Logger.Error("failed to sync transaction data with Plaid", zap.Error(err), zap.Any("plaid-error", plaidErr))
				if plaidErr.ErrorCode == "TRANSACTIONS_SYNC_MUTATION_DURING_PAGINATION" {
					// Restart the loop by setting the transactions slice to be empty
					added = []*schema.Transaction{}
					modified = []*schema.Transaction{}
					removed = []string{}
					continue
				}

				if plaidErr.ErrorCode == "ITEM_LOGIN_REQUIRED" {
					// link should be set to be updated
					linkShouldBeInUpdateMode = true
				}
			} else {
				p.Logger.Error("failed to sync data with Plaid", zap.Error(err), zap.Any("request", request))
				return nil, err
			}
		}

		cursor = pointer.StringP(result.GetNextCursor())
		for _, transaction := range result.Added {
			newTx, err := transformer.NewTransactionFromPlaid(&transaction)
			if err != nil {
				return nil, err
			}

			added = append(added, newTx)
		}

		for _, transaction := range result.Modified {
			modifiedTx, err := transformer.NewTransactionFromPlaid(&transaction)
			if err != nil {
				return nil, err
			}

			modified = append(modified, modifiedTx)
		}

		for _, transaction := range result.Removed {
			removed = append(removed, transaction.GetTransactionId())
		}

		if !result.GetHasMore() {
			break
		}
	}

	if len(added)+len(modified)+len(removed) == 0 {
		p.Logger.Info("no changes observed from Plaid via sync")
	} else {
		p.Logger.Info("received changes from Plaid via sync", zap.Any("added", len(added)), zap.Any("modified", len(modified)), zap.Any("removed", len(removed)))
	}

	return &SyncResult{
		NextCursor:               *cursor,
		HasMore:                  false,
		New:                      added,
		Updated:                  modified,
		Deleted:                  removed,
		ItemLoginRequiredForLink: linkShouldBeInUpdateMode,
	}, nil
}
