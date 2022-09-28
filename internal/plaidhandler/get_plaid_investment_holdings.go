package plaidhandler

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/plaid"
)

type Investments struct {
	Securities []plaid.Security
	Holdings   []plaid.Holding
	Accounts   []plaid.AccountBase
}

func (p *PlaidWrapper) getPlaidInvestmentHoldings(ctx context.Context, accessToken *string) (*Investments, error) {
	// TODO: emit metrics
	txn := p.NewRelicClient.StartTransaction("GET_PLAID_INVESTMENTS")
	defer txn.End()

	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	segment := &newrelic.Segment{
		StartTime: txn.StartSegmentNow(),
		Name:      "plaid_investment_holdings_outbound_request",
	}
	defer segment.End()

	request := plaid.NewInvestmentsHoldingsGetRequest(*accessToken)
	options := plaid.NewInvestmentHoldingsGetRequestOptions()
	request.SetOptions(*options)

	resp, _, err := p.Client.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(*request).Execute()
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
