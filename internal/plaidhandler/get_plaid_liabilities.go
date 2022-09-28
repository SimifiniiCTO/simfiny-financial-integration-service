package plaidhandler

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/plaid"
)

type Liabilities struct {
	Liabilities plaid.LiabilitiesObject
	Accounts    []plaid.AccountBase
}

func (p *PlaidWrapper) getPlaidLiabilities(ctx context.Context, accessToken *string) (*Liabilities, error) {
	// TODO: emit metrics
	txn := p.NewRelicClient.StartTransaction("GET_PLAID_LIABILITIES")
	defer txn.End()

	if accessToken == nil {
		return nil, errors.New("invalid input argument. access token cannot be empty")
	}

	segment := &newrelic.Segment{
		StartTime: txn.StartSegmentNow(),
		Name:      "plaid_get_liabilities_outbound_request",
	}
	defer segment.End()

	request := plaid.NewLiabilitiesGetRequest(*accessToken)
	resp, _, err := p.Client.PlaidApi.LiabilitiesGet(ctx).LiabilitiesGetRequest(*request).Execute()
	if err != nil {
		return nil, err
	}

	return &Liabilities{
		Liabilities: resp.GetLiabilities(),
		Accounts:    filterAccountsOfTwoType(resp.GetAccounts(), LOAN, CREDIT),
	}, nil
}
