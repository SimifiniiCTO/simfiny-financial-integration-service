package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/v12/plaid"
)

// The type Deposits contains an array of Plaid account objects.
// @property {[]plaid.AccountBase} Accounts - The `Accounts` property is a slice of `plaid.AccountBase`
// structs. It represents a list of bank accounts that have been linked to a Plaid access token. Each
// `plaid.AccountBase` struct contains basic information about a single bank account, such as the
// account ID, account type
type Deposits struct {
	Accounts []plaid.AccountBase
}

// getPlaidDeposit is used to get deposit accounts from the Plaid API for a given access token
func (p *PlaidWrapper) getPlaidDeposit(ctx context.Context, accessToken *string) (*Deposits, error) {
	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	request := plaid.AccountsBalanceGetRequest{
		AccessToken: *accessToken,
		Secret:      &p.SecretKey,
		ClientId:    &p.ClientID,
	}
	resp, _, err := p.client.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(request).Execute()
	if err != nil {
		return nil, err
	}

	return &Deposits{Accounts: filterAccounts(resp.GetAccounts(), DEPOSITORY)}, nil
}
