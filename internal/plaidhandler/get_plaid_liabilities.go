package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/plaid"
)

type Liabilities struct {
	Liabilities plaid.LiabilitiesObject
	Accounts    []plaid.AccountBase
}

func (p *PlaidWrapper) getPlaidLiabilities(ctx context.Context, accessToken *string) (*Liabilities, error) {
	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	request := plaid.NewLiabilitiesGetRequest(*accessToken)
	resp, _, err := p.client.PlaidApi.LiabilitiesGet(ctx).LiabilitiesGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &Liabilities{
		Liabilities: resp.GetLiabilities(),
		Accounts:    filterAccountsOfTwoType(resp.GetAccounts(), LOAN, CREDIT),
	}, nil
}
