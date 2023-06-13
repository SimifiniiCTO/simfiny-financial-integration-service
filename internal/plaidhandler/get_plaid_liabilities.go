package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/v12/plaid"
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
	request.SetClientId(p.ClientID)
	request.SetSecret(p.SecretKey)

	resp, _, err := p.client.PlaidApi.LiabilitiesGet(ctx).LiabilitiesGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &Liabilities{
		Liabilities: resp.GetLiabilities(),
		Accounts:    filterAccountsOfTwoType(resp.GetAccounts(), LOAN, CREDIT),
	}, nil
}
