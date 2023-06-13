package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/v12/plaid"
)

// transformStudentloanObject transforms a plaid student loan liabilities object to the internal mortgage account object
func transformStudentloanObject(studentLoans *[]plaid.StudentLoan, acctIDToTypeMap map[string]*accountMetadata) ([]*schema.
	StudentLoanAccount, error) {
	if studentLoans == nil {
		return nil, errors.New("invalid input argument. student loan liability cannot be nil")
	}

	if acctIDToTypeMap == nil {
		return nil, errors.New("invalid input argument. account id to type map is nil")
	}

	accts := make([]*schema.StudentLoanAccount, 0)
	for _, element := range *studentLoans {
		metadata := acctIDToTypeMap[element.GetAccountId()]

		repaymentPlan := element.GetRepaymentPlan()
		servisorAddress := element.GetServicerAddress()
		loanStatus := element.GetLoanStatus()

		studentloanAcct := &schema.StudentLoanAccount{
			PlaidAccountID:             element.GetAccountId(),
			AccountSubtype:             metadata.accountSubtype,
			AccountType:                metadata.accountType,
			AccountNumber:              element.GetAccountNumber(),
			DisbursementDates:          element.GetDisbursementDates(),
			ExpectedPayoffDate:         element.GetExpectedPayoffDate(),
			Guarantor:                  element.GetGuarantor(),
			InterestRatePercentage:     float64(element.GetInterestRatePercentage()),
			IsOverdue:                  element.GetIsOverdue(),
			LastPaymentAmount:          float64(element.GetLastPaymentAmount()),
			LastPaymentDate:            element.GetLastPaymentDate(),
			LastStatementIssueDate:     element.GetLastStatementIssueDate(),
			LoanStatementIssueDate:     loanStatus.GetEndDate(),
			LoanName:                   element.GetLoanName(),
			LoanRepaymentEndDate:       element.GetExpectedPayoffDate(),
			MinimumPaymentAmount:       float64(element.GetMinimumPaymentAmount()),
			NextPaymentDueDate:         element.GetNextPaymentDueDate(),
			OriginationDate:            element.GetOriginationDate(),
			OriginationPrincipalAmount: float64(element.GetOriginationPrincipalAmount()),
			OutstandingInterestAmount:  float64(element.GetOutstandingInterestAmount()),
			PaymentReferenceNumber:     element.GetPaymentReferenceNumber(),
			SequenceNumber:             element.GetSequenceNumber(),
			PslfID:                     nil,
			RepaymentPlan:              repaymentPlan.GetDescription(),
			ServisorAddressID: &schema.Address{
				City:    servisorAddress.GetCity(),
				Country: servisorAddress.GetCountry(),
				Region:  servisorAddress.GetRegion(),
				Street:  servisorAddress.GetStreet(),
			},
			InterestPaidYTD:  float64(element.GetYtdInterestPaid()),
			PrincipalPaidYTD: float64(element.GetYtdPrincipalPaid()),
			BalanceID:        metadata.balance,
			AccountName:      metadata.accountName,
		}

		accts = append(accts, studentloanAcct)
	}

	return accts, nil
}
