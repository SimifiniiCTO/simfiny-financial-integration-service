package plaidhandler

import (
	"context"
	"time"

	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
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
		if retrieved := int32(len(someTransactions)); retrieved >= perPage {
			offset += retrieved
			continue
		}

		break
	}

	return transactions, nil
}

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
