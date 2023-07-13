package plaidhandler

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/v14/plaid"
)

// The Investments type contains arrays of securities, holdings, and account bases from the Plaid API.
// @property {[]plaid.Security} Securities - This is a property of the Investments struct that
// represents a slice of plaid.Security objects. Securities are financial instruments that can be
// bought and sold, such as stocks, bonds, and mutual funds. The Securities property likely contains
// information about the securities that the user has invested in through their linked investment
// accounts.
// @property {[]plaid.Holding} Holdings - Holdings is a property of the Investments struct that
// represents a slice of plaid.Holding objects. A holding is a financial instrument that an investor
// holds in their portfolio, such as stocks, bonds, or mutual funds. The plaid.Holding object contains
// information about the holding, such as its name
// @property {[]plaid.AccountBase} Accounts - The `Accounts` property is a slice of `plaid.AccountBase`
// structs, which represent the bank accounts linked to a Plaid Item. Each `AccountBase` struct
// contains information such as the account ID, account type (e.g. checking, savings), current balance,
// and available balance.
type Investments struct {
	Securities []plaid.Security
	Holdings   []plaid.Holding
	Accounts   []plaid.AccountBase
}

// GetPlaidInvestmentHoldings is used to retrieve investment holdings from the Plaid API for a given set of account IDs
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
