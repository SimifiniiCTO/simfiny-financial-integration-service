package transformer

import (
	"fmt"

	"github.com/plaid/plaid-go/v14/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// NewCreditAccount transforms plaid credit account objects to internal plaid credit account representation
func NewCreditAccount(userID uint64, input *plaid.CreditCardLiability, acct *plaid.AccountBase) (*schema.CreditAccount, error) {
	if input == nil {
		return nil, fmt.Errorf("invalid input argument. credit card liability cannot be nil")
	}

	if acct == nil {
		return nil, fmt.Errorf("invalid input argument. account base cannot be nil")
	}

	if acct.Type != plaid.ACCOUNTTYPE_CREDIT {
		return nil, fmt.Errorf("invalid input argument. account base must be of type credit")
	}

	aprs, err := NewAPR(&input.Aprs)
	if err != nil {
		return nil, err
	}

	return &schema.CreditAccount{
		Id:                     0,
		UserId:                 userID,
		Name:                   acct.Name,
		Number:                 *acct.Mask.Get(),
		Type:                   string(acct.Type),
		Balance:                float32(acct.Balances.GetAvailable()),
		CurrentFunds:           float64(acct.Balances.GetCurrent()),
		BalanceLimit:           uint64(acct.Balances.GetLimit()),
		PlaidAccountId:         acct.AccountId,
		Subtype:                string(acct.GetSubtype()),
		IsOverdue:              input.GetIsOverdue(),
		LastPaymentAmount:      float64(*input.LastPaymentAmount.Get()),
		LastPaymentDate:        input.GetLastPaymentDate(),
		LastStatementIssueDate: *input.LastStatementIssueDate.Get(),
		MinimumAmountDueDate:   0,
		NextPaymentDate:        *input.NextPaymentDueDate.Get(),
		Aprs:                   aprs,
		LastStatementBalance:   float64(*input.LastStatementBalance.Get()),
		MinimumPaymentAmount:   float64(input.GetMinimumPaymentAmount()),
		NextPaymentDueDate:     input.GetNextPaymentDueDate(),
	}, nil
}

// NewAPR transforms plaid APR objects to internal plaid APR representation
func NewAPR(plaidAPR *[]plaid.APR) ([]*schema.Apr, error) {
	if plaidAPR == nil {
		return nil, fmt.Errorf("invalid input argument. apr cannot be nil")
	}

	aprs := make([]*schema.Apr, 0)
	for _, element := range *plaidAPR {
		transformedApr := &schema.Apr{
			Id:                   0,
			Percentage:           element.AprPercentage,
			Type:                 string(element.AprType),
			BalanceSubjectToApr:  float64(element.GetBalanceSubjectToApr()),
			InterestChargeAmount: float64(element.GetInterestChargeAmount()),
		}

		aprs = append(aprs, transformedApr)
	}

	return aprs, nil
}
