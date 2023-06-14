package transformer

import (
	"errors"

	"github.com/plaid/plaid-go/v12/plaid"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// transformMortgageObject transforms a plaid mortgage loan liabilities object to the internal mortgage account object
func TransformMortgageAccount(mortgageLoan *[]plaid.MortgageLiability, acctIDToTypeMap map[string]*accountMetadata) ([]*schema.
	MortgageAccount, error) {
	if mortgageLoan == nil {
		return nil, errors.New("invalid input argument. mortgage loan liability cannot be nil")
	}

	if acctIDToTypeMap == nil {
		return nil, errors.New("invalid input argument. account id to type map is nil")
	}

	accts := make([]*schema.MortgageAccount, 0)
	for _, element := range *mortgageLoan {
		//metadata := acctIDToTypeMap[element.AccountId]
		interest := element.GetInterestRate()
		mortgageAcct := &schema.MortgageAccount{
			Id:                          0,
			PlaidAccountId:              element.GetAccountId(),
			AccountNumber:               element.GetAccountNumber(),
			CurrentLateFee:              float64(element.GetCurrentLateFee()),
			EscrowBalance:               float64(element.GetEscrowBalance()),
			HasPmi:                      element.GetHasPmi(),
			HasPrepaymentPenalty:        element.GetHasPrepaymentPenalty(),
			LastPaymentAmount:           float64(element.GetLastPaymentAmount()),
			LastPaymentDate:             element.GetLastPaymentDate(),
			LoanTerm:                    element.GetLoanTerm(),
			LoanTypeDescription:         element.GetLoanTypeDescription(),
			MaturityDate:                element.GetMaturityDate(),
			NextMonthlyPayment:          float64(element.GetNextMonthlyPayment()),
			NextPaymentDueDate:          element.GetNextPaymentDueDate(),
			OriginalPrincipalBalance:    float64(element.GetOriginationPrincipalAmount()),
			OriginalPropertyValue:       0,
			OutstandingPrincipalBalance: 0,
			PaymentAmount:               element.GetLastPaymentAmount(),
			PaymentDate:                 element.GetLastPaymentDate(),
			OriginationDate:             element.GetOriginationDate(),
			OriginationPrincipalAmount:  float64(element.GetOriginationPrincipalAmount()),
			PastDueAmount:               float64(element.GetPastDueAmount()),
			YtdInterestPaid:             *element.YtdInterestPaid.Get(),
			YtdPrincipalPaid:            *element.YtdPrincipalPaid.Get(),
			PropertyAddressCity:         element.PropertyAddress.GetCity(),
			PropertyAddressState:        "",
			PropertyAddressStreet:       element.PropertyAddress.GetStreet(),
			PropertyAddressPostalCode:   element.PropertyAddress.GetPostalCode(),
			PropertyRegion:              element.PropertyAddress.GetRegion(),
			PropertyCountry:             *element.PropertyAddress.Country.Get(),
			InterestRatePercentage:      interest.GetPercentage(),
			InterestRateType:            interest.GetType(),
		}

		accts = append(accts, mortgageAcct)
	}

	return accts, nil
}
