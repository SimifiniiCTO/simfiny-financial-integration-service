package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/plaid"
)

// transformMortgageObject transforms a plaid mortgage loan liabilities object to the internal mortgage account object
func transformMortgageObject(mortgageLoan *[]plaid.MortgageLiability, acctIDToTypeMap map[string]*accountMetadata) ([]*schema.
	MortgageLoanAccount, error) {
	if mortgageLoan == nil {
		return nil, errors.New("invalid input argument. mortgage loan liability cannot be nil")
	}

	if acctIDToTypeMap == nil {
		return nil, errors.New("invalid input argument. account id to type map is nil")
	}

	accts := make([]*schema.MortgageLoanAccount, 0)
	for _, element := range *mortgageLoan {
		metadata := acctIDToTypeMap[element.AccountId]
		interest := element.GetInterestRate()
		address := element.GetPropertyAddress()
		mortgageAcct := &schema.MortgageLoanAccount{
			PlaidAccountID:             element.AccountId,
			AccountSubtype:             metadata.accountSubtype,
			AccountType:                metadata.accountType,
			AccountNumber:              element.GetAccountNumber(),
			CurrentLateFee:             float64(element.GetCurrentLateFee()),
			EscrowBalance:              float64(element.GetEscrowBalance()),
			HasPmi:                     element.GetHasPmi(),
			HasPrepaymentPenalty:       element.GetHasPrepaymentPenalty(),
			LastPaymentAmount:          float64(element.GetLastPaymentAmount()),
			LastPaymentDate:            element.GetLastPaymentDate(),
			LoanTerm:                   element.GetLoanTerm(),
			LoanTypeDescription:        element.GetLoanTypeDescription(),
			MaturityDate:               element.GetMaturityDate(),
			NexMonthlyPayment:          float64(element.GetNextMonthlyPayment()),
			NextPaymentDueDate:         element.GetNextPaymentDueDate(),
			OriginationDate:            element.GetOriginationDate(),
			OriginationPrincipalAmount: float64(element.GetOriginationPrincipalAmount()),
			PastDueAmount:              float64(element.GetPastDueAmount()),
			InterestID: &schema.Interest{
				Percentage: float64(interest.GetPercentage()),
				Type:       interest.GetType(),
			},
			AddressID: &schema.Address{
				City:    address.GetCity(),
				Country: address.GetCountry(),
				Region:  address.GetRegion(),
				Street:  address.GetStreet(),
			},
			InterestPaidYTD:  float64(element.GetYtdInterestPaid()),
			PrincipalPaidYTD: float64(element.GetYtdPrincipalPaid()),
			BalanceID:        metadata.balance,
			AccountName:      metadata.accountName,
		}

		accts = append(accts, mortgageAcct)
	}

	return accts, nil
}
