package plaidhandler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/plaid/plaid-go/v12/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func (p *PlaidWrapper) GetInvestmentTransactions(ctx context.Context, accessToken *string, userId *uint64, link *schema.Link, startDate, endDate string, accountIds []string, numTransactionsToFetch int32, offset int32) ([]*schema.InvestmentTransaction, error) {
	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	opts := plaid.InvestmentsTransactionsGetRequestOptions{}

	if accountIds != nil {
		opts.AccountIds = &accountIds
	}

	if numTransactionsToFetch != 0 {
		opts.Count = &numTransactionsToFetch
	}

	if offset != 0 {
		opts.Offset = &offset
	}

	if link == nil {
		return nil, errors.New("invalid input argument. link cannot be empty")
	}

	investmentAccountIds := make([]string, 0, len(link.InvestmentAccounts))
	for _, acct := range link.InvestmentAccounts {
		investmentAccountIds = append(investmentAccountIds, acct.PlaidAccountId)
	}

	if len(investmentAccountIds) > 0 {
		if opts.AccountIds == nil {
			opts.AccountIds = &investmentAccountIds
		} else {
			ids := append(*opts.AccountIds, investmentAccountIds...)
			opts.AccountIds = &ids
		}
	}

	request := plaid.InvestmentsTransactionsGetRequest{
		AccessToken: *accessToken,
		Secret:      &p.SecretKey,
		ClientId:    &p.ClientID,
		StartDate:   startDate,
		EndDate:     endDate,
		Options:     &opts,
	}

	return p.getInvestmentTransactions(ctx, &request, userId, link)
}

func (p *PlaidWrapper) getInvestmentTransactions(ctx context.Context, req *plaid.InvestmentsTransactionsGetRequest, userId *uint64, link *schema.Link) ([]*schema.InvestmentTransaction, error) {
	resp, _, err := p.client.PlaidApi.
		InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(*req).Execute()
	if err != nil {
		return nil, err
	}

	results := make([]*schema.InvestmentTransaction, 0, len(resp.GetInvestmentTransactions()))
	for _, t := range resp.GetInvestmentTransactions() {

		newTx := &schema.InvestmentTransaction{
			AccountId:               t.GetAccountId(),
			Ammount:                 fmt.Sprintf("%f", t.GetAmount()),
			InvestmentTransactionId: t.GetInvestmentTransactionId(),
			SecurityId:              t.GetSecurityId(),
			Date:                    t.GetDate(),
			Name:                    t.GetName(),
			Quantity:                t.GetQuantity(),
			Amount:                  t.GetAmount(),
			Price:                   t.GetPrice(),
			Fees:                    t.GetFees(),
			Type:                    string(t.GetType()),
			Subtype:                 string(t.GetSubtype()),
			IsoCurrencyCode:         t.GetIsoCurrencyCode(),
			LinkId:                  link.Id,
			Id:                      0,
			UserId:                  *userId,
			CreatedAt:               time.Now().String(),
			UnofficialCurrencyCode:  t.GetUnofficialCurrencyCode(),
		}

		results = append(results, newTx)
	}

	// transform transactions
	return results, nil
}
