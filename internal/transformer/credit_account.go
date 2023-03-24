package transformer

import (
	"fmt"

	"github.com/plaid/plaid-go/plaid"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

// NewCreditAccount transforms plaid credit account objects to internal plaid credit account representation
func NewCreditAccount(userID uint64, input *plaid.CreditCardLiability, acct *plaid.AccountBase) (*proto.CreditAccount, error) {
	if input == nil {
		return nil, fmt.Errorf("invalid input argument. credit card liability cannot be nil")
	}

	aprs, err := NewAPR(&input.Aprs)
	if err != nil {
		return nil, err
	}

	return &proto.CreditAccount{
		Id:                     userID,
		PlaidAccountID:         input.GetAccountId(),
		AccountSubtype:         string(acct.GetSubtype()),
		AccountType:            string(acct.GetSubtype()),
		IsOverdue:              input.GetIsOverdue(),
		LastPaymentAmount:      float64(input.LastPaymentAmount),
		LastPaymentDate:        input.GetLastPaymentDate(),
		LastStatementIssueDate: input.LastStatementIssueDate,
		MinimumPaymentAmount:   float64(input.GetMinimumPaymentAmount()),
		NextPaymentDueDate:     input.GetNextPaymentDueDate(),
		Aprs:                   aprs,
		BalanceID: &proto.Balance{
			AvailableFunds: float64(acct.Balances.GetAvailable()),
			CurrentFunds:   float64(acct.Balances.GetCurrent()),
			CurrencyCode:   acct.Balances.GetIsoCurrencyCode(),
			BalanceLimit:   uint64(acct.Balances.GetLimit()),
		},
		AccountName: acct.GetName(),
	}, nil
}

// NewAPR transforms plaid APR objects to internal plaid APR representation
func NewAPR(plaidAPR *[]plaid.APR) ([]*proto.APR, error) {
	if plaidAPR == nil {
		return nil, fmt.Errorf("invalid input argument. apr cannot be nil")
	}

	aprs := make([]*proto.APR, 0)
	for _, element := range *plaidAPR {
		transformedApr := &proto.APR{
			APRPercentage:        float64(element.GetAprPercentage()),
			APRType:              element.GetAprType(),
			BalanceSubjectToAPR:  float64(element.GetBalanceSubjectToApr()),
			InterestChargeAmount: float64(element.GetInterestChargeAmount()),
		}

		aprs = append(aprs, transformedApr)
	}

	return aprs, nil
}
