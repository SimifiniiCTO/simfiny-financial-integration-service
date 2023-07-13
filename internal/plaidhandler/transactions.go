package plaidhandler

import (
	"context"
	"time"

	"github.com/plaid/plaid-go/v14/plaid"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// GetAllTransactions is used to retrieve all transactions for
// a given set of account IDs within a specified date range. It takes in the `accessToken` for the
// Plaid API, the start and end dates for the transaction search, and an array of account IDs.
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
		if retrieved := int32(len(someTransactions)); retrieved >= perPage {
			offset += retrieved
			continue
		}

		break
	}

	return transactions, nil
}

// GetTransactions is used to retrieve transactions from the Plaid API for a given set of account IDs
// within a specified date range.
//
// It takes in the `accessToken` for the Plaid API, the start and end
// dates for the transaction search, an optional count and offset for pagination, and an array of
// account IDs. It constructs a `TransactionsGetRequest` using the Plaid API client and sends the
// request to retrieve the transactions.
//
// It then transforms the retrieved transactions from the Plaid
// format to the schema format using the `transformer.NewTransactionFromPlaid` function and returns
// them as an array of `schema.Transaction` objects.
func (p *PlaidWrapper) GetTransactions(ctx context.Context, accessToken string, start, end time.Time, count, offset int32, bankAccountIds []string) ([]*schema.Transaction, error) {
	opts := &plaid.TransactionsGetRequestOptions{
		Count:                              &count,
		Offset:                             &offset,
		IncludeOriginalDescription:         *plaid.NewNullableBool(pointer.BoolP(true)),
		IncludePersonalFinanceCategoryBeta: pointer.BoolP(true),
		IncludePersonalFinanceCategory:     pointer.BoolP(true),
		IncludeLogoAndCounterpartyBeta:     pointer.BoolP(true),
	}

	if len(bankAccountIds) > 0 {
		opts.AccountIds = &bankAccountIds
	}

	request := p.client.PlaidApi.
		TransactionsGet(ctx).
		TransactionsGetRequest(plaid.TransactionsGetRequest{
			Options:     opts,
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

	transactions := make([]*schema.Transaction, 0, len(result.Transactions))
	for _, transaction := range result.Transactions {
		transaction, err := transformer.NewTransactionFromPlaid(&transaction)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
