package plaidhandler

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/plaid"
)

type Deposits struct {
	Accounts []plaid.AccountBase
}

func (p *PlaidWrapper) getPlaidDeposit(ctx context.Context, accessToken *string) (*Deposits, error) {
	// TODO: emit metrics
	txn := p.NewRelicClient.StartTransaction("GET_PLAID_DEPOSIT")
	defer txn.End()

	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	segment := &newrelic.Segment{
		StartTime: txn.StartSegmentNow(),
		Name:      "plaid_new_account_balance_outbound_request",
	}
	defer segment.End()

	request := plaid.NewAccountsBalanceGetRequest(*accessToken)
	resp, _, err := p.Client.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &Deposits{Accounts: filterAccounts(resp.GetAccounts(), DEPOSITORY)}, nil
}
