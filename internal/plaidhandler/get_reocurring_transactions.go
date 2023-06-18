package plaidhandler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

// GetRecurringTransactionsForAccounts is used to retrieve recurring transactions from the Plaid API for a given set of account IDs
func (p *PlaidWrapper) GetRecurringTransactionsForAccounts(ctx context.Context, accessToken *string, userId *uint64, linkId *uint64, plaidAccountIds []string) ([]*schema.ReOccuringTransaction, error) {
	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	if userId == nil {
		return nil, errors.New("invalid input argument. user id cannot be empty")
	}

	if linkId == nil {
		return nil, errors.New("invalid input argument. link id cannot be empty")
	}

	if len(plaidAccountIds) == 0 {
		return nil, errors.New("invalid input argument. plaid account ids cannot be empty")
	}

	includePersonalFinanceCategory := true
	request := plaid.TransactionsRecurringGetRequest{
		AccessToken: *accessToken,
		Secret:      &p.SecretKey,
		ClientId:    &p.ClientID,
		AccountIds:  plaidAccountIds,
		Options: &plaid.TransactionsRecurringGetRequestOptions{
			IncludePersonalFinanceCategory: &includePersonalFinanceCategory,
		},
	}

	return p.getRecurringTransactions(ctx, &request, userId, linkId)

}

// GetRecurringTransactions is used to retrieve recurring transactions from the Plaid API for a userId and linkId
func (p *PlaidWrapper) GetRecurringTransactions(ctx context.Context, accessToken *string, userId *uint64, linkId *uint64) ([]*schema.ReOccuringTransaction, error) {
	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	if userId == nil {
		return nil, errors.New("invalid input argument. user id cannot be empty")
	}

	if linkId == nil {
		return nil, errors.New("invalid input argument. link id cannot be empty")
	}

	includePersonalFinanceCategory := true
	request := plaid.TransactionsRecurringGetRequest{
		AccessToken: *accessToken,
		Secret:      &p.SecretKey,
		ClientId:    &p.ClientID,
		Options: &plaid.TransactionsRecurringGetRequestOptions{
			IncludePersonalFinanceCategory: &includePersonalFinanceCategory,
		},
	}

	return p.getRecurringTransactions(ctx, &request, userId, linkId)
}

func (p *PlaidWrapper) getRecurringTransactions(ctx context.Context, req *plaid.TransactionsRecurringGetRequest, userId *uint64, linkId *uint64) ([]*schema.ReOccuringTransaction, error) {
	resp, _, err := p.client.PlaidApi.
		TransactionsRecurringGet(ctx).
		TransactionsRecurringGetRequest(*req).Execute()
	if err != nil {
		return nil, err
	}

	inflowStreams := transactionStreamToRecurringTransactions(userId, linkId, resp.GetInflowStreams(), schema.ReCurringFlow_RE_CURRING_FLOW_INFLOW)
	outflowStream := transactionStreamToRecurringTransactions(userId, linkId, resp.GetOutflowStreams(), schema.ReCurringFlow_RE_CURRING_FLOW_OUTFLOW)
	transactions := make([]*schema.ReOccuringTransaction, 0, len(inflowStreams)+len(outflowStream))
	transactions = append(transactions, inflowStreams...)
	transactions = append(transactions, outflowStream...)

	return transactions, nil
}

func transactionStreamToRecurringTransactions(userId, linkId *uint64, streams []plaid.TransactionStream, flow schema.ReCurringFlow) []*schema.ReOccuringTransaction {
	recurringTransactions := make([]*schema.ReOccuringTransaction, 0, len(streams))
	for _, stream := range streams {
		recurringTransactions = append(recurringTransactions, &schema.ReOccuringTransaction{
			AccountId:                       stream.GetAccountId(),
			StreamId:                        stream.GetStreamId(),
			CategoryId:                      stream.GetCategoryId(),
			Description:                     stream.GetDescription(),
			MerchantName:                    stream.GetMerchantName(),
			PersonalFinanceCategoryPrimary:  stream.GetPersonalFinanceCategory().Primary,
			PersonalFinanceCategoryDetailed: stream.GetPersonalFinanceCategory().Detailed,
			FirstDate:                       stream.GetFirstDate(),
			LastDate:                        stream.GetLastDate(),
			Frequency:                       getFrequency(stream.GetFrequency()),
			TransactionIds:                  helper.SliceToCommaSeparatedString(stream.GetTransactionIds()),
			AverageAmount:                   fmt.Sprintf("%f", *stream.GetAverageAmount().Amount),
			AverageAmountIsoCurrencyCode:    *stream.GetAverageAmount().IsoCurrencyCode.Get(),
			LastAmount:                      fmt.Sprintf("%f", *stream.LastAmount.Amount),
			LastAmountIsoCurrencyCode:       *stream.GetLastAmount().IsoCurrencyCode.Get(),
			IsActive:                        stream.GetIsActive(),
			Status:                          getStatus(stream.GetStatus()),
			UpdatedTime:                     time.Now().String(),
			UserId:                          *userId,
			LinkId:                          *linkId,
			Id:                              0,
			Flow:                            flow,
		})
	}

	return recurringTransactions
}

func getFrequency(freq plaid.RecurringTransactionFrequency) schema.ReOccuringTransactionsFrequency {
	switch freq {
	case plaid.RECURRINGTRANSACTIONFREQUENCY_ANNUALLY:
		return schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_ANNUALLY
	case plaid.RECURRINGTRANSACTIONFREQUENCY_BIWEEKLY:
		return schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_BIWEEKLY
	case plaid.RECURRINGTRANSACTIONFREQUENCY_UNKNOWN:
		return schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED
	case plaid.RECURRINGTRANSACTIONFREQUENCY_MONTHLY:
		return schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_MONTHLY
	case plaid.RECURRINGTRANSACTIONFREQUENCY_SEMI_MONTHLY:
		return schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_SEMI_MONTHLY
	case plaid.RECURRINGTRANSACTIONFREQUENCY_WEEKLY:
		return schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_WEEKLY
	default:
		return schema.ReOccuringTransactionsFrequency_RE_OCCURING_TRANSACTIONS_FREQUENCY_UNSPECIFIED
	}
}

func getStatus(status plaid.TransactionStreamStatus) schema.ReOccuringTransactionsStatus {
	switch status {
	case plaid.TRANSACTIONSTREAMSTATUS_EARLY_DETECTION:
		return schema.ReOccuringTransactionsStatus_RE_OCCURING_TRANSACTIONS_STATUS_EARLY_DETECTION
	case plaid.TRANSACTIONSTREAMSTATUS_TOMBSTONED:
		return schema.ReOccuringTransactionsStatus_RE_OCCURING_TRANSACTIONS_STATUS_TOMBSTONED
	case plaid.TRANSACTIONSTREAMSTATUS_MATURE:
		return schema.ReOccuringTransactionsStatus_RE_OCCURING_TRANSACTIONS_STATUS_MATURE
	default:
		return schema.ReOccuringTransactionsStatus_RE_OCCURING_TRANSACTIONS_STATUS_UNSPECIFIED
	}
}
