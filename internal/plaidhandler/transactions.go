package plaidhandler

import (
	"context"
	"time"

	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
)

// GetAllTransactions implements PlaidWrapperImpl
func (p *PlaidWrapper) GetAllTransactions(ctx context.Context, accessToken string, start time.Time, end time.Time, accountIds []string) ([]*schema.Transaction, error) {
	var perPage int32 = 500
	var offset int32 = 0
	transactions := make([]*schema.Transaction, 0)

	for {
		someTransactions, err := p.GetTransactions(ctx, accessToken, start, end, perPage, offset, accountIds)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, someTransactions...)
		if retrieved := int32(len(someTransactions)); retrieved == perPage {
			offset += retrieved
			continue
		}

		break
	}

	return transactions, nil
}

func (p *PlaidWrapper) GetTransactions(ctx context.Context, accessToken string, start, end time.Time, count, offset int32, bankAccountIds []string) ([]*schema.Transaction, error) {
	request := p.client.PlaidApi.
		TransactionsGet(ctx).
		TransactionsGetRequest(plaid.TransactionsGetRequest{
			Options: &plaid.TransactionsGetRequestOptions{
				AccountIds:                 &bankAccountIds,
				Count:                      &count,
				Offset:                     &offset,
				IncludeOriginalDescription: plaid.NullableBool{},
			},
			AccessToken: accessToken,
			StartDate:   start.Format("2006-01-02"),
			EndDate:     end.Format("2006-01-02"),
		})

	// Send the request.
	result, _, err := request.Execute()
	if err != nil {
		p.Logger.Error("failed to retrieve transactions from plaid", zap.Error(err))
		return nil, err
	}

	transactions := make([]*schema.Transaction, len(result.Transactions))
	for i, transaction := range result.Transactions {
		transactions[i], err = transformer.NewTransactionFromPlaid(transaction)
		if err != nil {
			return nil, err
		}
	}

	return transactions, nil
}
