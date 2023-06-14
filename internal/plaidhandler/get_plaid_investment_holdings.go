package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/v12/plaid"
)

type Investments struct {
	Securities []plaid.Security
	Holdings   []plaid.Holding
	Accounts   []plaid.AccountBase
}

func (p *PlaidWrapper) GetPlaidInvestmentHoldings(ctx context.Context, accessToken *string, accountIds []string) (*Investments, error) {
	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	var opts plaid.InvestmentHoldingsGetRequestOptions
	if len(accountIds) > 0 {
		opts = plaid.InvestmentHoldingsGetRequestOptions{
			AccountIds: &accountIds,
		}
	}

	request := plaid.NewInvestmentsHoldingsGetRequest(*accessToken)
	request.SetClientId(p.ClientID)
	request.SetSecret(p.SecretKey)
	request.SetOptions(opts)

	resp, _, err := p.client.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}
	// Perform transformation from plaid response to
	return &Investments{
		Securities: resp.GetSecurities(),
		Holdings:   resp.GetHoldings(),
		Accounts:   filterAccounts(resp.GetAccounts(), INVESTMENTS),
	}, nil
}
