package plaidhandler

import (
	"context"
	"errors"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/transformer"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/v12/plaid"
)

const (
	DEPOSITORY  = "depository"
	CREDIT      = "credit"
	LOAN        = "loan"
	INVESTMENTS = "investment"
)

func (p *PlaidWrapper) populateVirtualAccount(ctx context.Context, liabilitiesObj *Liabilities, investmentsObj *Investments,
	depositsObj *Deposits) (*proto.VirtualAccount,
	error) {
	// TODO: trace and emit metrics
	if liabilitiesObj == nil {
		return nil, errors.New("invalid input argument. liabilitiesObj cannot be nil")
	}

	var vAcct proto.VirtualAccount
	if err := transformer.TransformPlaidLiabilitiesResponse(ctx, &vAcct, &liabilitiesObj.Liabilities, &liabilitiesObj.Accounts); err != nil {
		return nil, err
	}

	if err := transformer.TransformPlaidInvestmentResponse(ctx, &vAcct, &investmentsObj.Securities, &investmentsObj.Holdings,
		&investmentsObj.Accounts); err != nil {
		return nil, err
	}

	if err := transformer.TransformPlaidDepositResponse(ctx, &vAcct, &depositsObj.Accounts); err != nil {
		return nil, err
	}

	return &vAcct, nil
}

func filterAccounts(accts []plaid.AccountBase, acctType string) []plaid.AccountBase {
	result := make([]plaid.AccountBase, 0)

	for _, acct := range accts {
		if acct.Type == plaid.AccountType(acctType) {
			result = append(result, acct)
		}
	}

	return result
}

func filterAccountsOfTwoType(accts []plaid.AccountBase, acctType1, acctType2 string) []plaid.AccountBase {
	result := make([]plaid.AccountBase, 0)

	for _, acct := range accts {
		if acct.Type == plaid.AccountType(acctType1) || acct.Type == plaid.AccountType(acctType2) {
			result = append(result, acct)
		}
	}

	return result
}
