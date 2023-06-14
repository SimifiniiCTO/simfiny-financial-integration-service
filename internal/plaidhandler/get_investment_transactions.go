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

	investmentAccountIds := make([]string, 0, len(link.InvestmentAccounts))
	for _, acct := range link.InvestmentAccounts {
		investmentAccountIds = append(investmentAccountIds, acct.PlaidAccountId)
	}

	// get transactions for specific investment account ids
	opts.AccountIds = &investmentAccountIds

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

	// transform transactions
	return toInvestmentTransactions(resp.GetInvestmentTransactions(), link.Id, *userId), nil
}

func toInvestmentTransactions(tx []plaid.InvestmentTransaction, linkId, userId uint64) []*schema.InvestmentTransaction {
	res := make([]*schema.InvestmentTransaction, 0, len(tx))
	for _, t := range tx {
		res = append(res, &schema.InvestmentTransaction{
			AccountId:               t.AccountId,
			Ammount:                 fmt.Sprintf("%f", t.GetAmount()),
			InvestmentTransactionId: t.InvestmentTransactionId,
			SecurityId:              *t.SecurityId.Get(),
			Date:                    t.Date,
			Name:                    t.Name,
			Quantity:                t.Quantity,
			Amount:                  t.Amount,
			Price:                   t.Price,
			Fees:                    *t.Fees.Get(),
			Type:                    string(t.GetType()),
			Subtype:                 string(t.Subtype),
			IsoCurrencyCode:         t.GetIsoCurrencyCode(),
			UnofficialCurrencyCode:  *t.UnofficialCurrencyCode.Get(),
			LinkId:                  linkId,
			Id:                      0,
			UserId:                  userId,
			CreatedAt:               time.Now().String(),
		})
	}

	return res
}
