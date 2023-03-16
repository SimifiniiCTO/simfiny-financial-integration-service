package plaidhandler

import (
	"context"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

func (p *PlaidWrapper) CallPlaidAPIAndPopulateVirtualAccount(ctx context.Context, accessToken *string) (*proto.VirtualAccount, error) {
	deposits, err := p.getPlaidDeposit(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	liabilities, err := p.getPlaidLiabilities(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	investments, err := p.getPlaidInvestmentHoldings(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	return p.populateVirtualAccount(ctx, liabilities, investments, deposits)
}
