package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/plaid"
)

type Deposits struct {
	Accounts []plaid.AccountBase
}

func (p *PlaidWrapper) getPlaidDeposit(ctx context.Context, accessToken *string) (*Deposits, error) {
	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	request := plaid.NewAccountsBalanceGetRequest(*accessToken)
	resp, _, err := p.Client.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &Deposits{Accounts: filterAccounts(resp.GetAccounts(), DEPOSITORY)}, nil
}
