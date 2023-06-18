package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

// TransformStudentloanAccount transforms a list of Plaid student loan objects and a map of account metadata into a
// list of schema StudentLoanAccount objects.
func TransformStudentloanAccount(studentLoans *[]plaid.StudentLoan, acctIDToTypeMap map[string]*accountMetadata) ([]*schema.
	StudentLoanAccount, error) {
	if studentLoans == nil {
		return nil, errors.New("invalid input argument. student loan liability cannot be nil")
	}

	if acctIDToTypeMap == nil {
		return nil, errors.New("invalid input argument. account id to type map is nil")
	}

	accts := make([]*schema.StudentLoanAccount, 0)
	for _, element := range *studentLoans {
		metadata, ok := acctIDToTypeMap[element.GetAccountId()]
		if !ok {
			return nil, errors.New("invalid input argument. account id to type map does not contain account id")
		}

		repaymentPlan := element.GetRepaymentPlan()
		servisorAddress := element.GetServicerAddress()

		studentloanAcct := &schema.StudentLoanAccount{
			Id:                                 0,
			PlaidAccountId:                     element.GetAccountId(),
			DisbursementDates:                  element.GetDisbursementDates(),
			ExpectedPayoffDate:                 element.GetExpectedPayoffDate(),
			Guarantor:                          element.GetGuarantor(),
			InterestRatePercentage:             float64(element.GetInterestRatePercentage()),
			IsOverdue:                          element.GetIsOverdue(),
			LastPaymentAmount:                  float64(element.GetLastPaymentAmount()),
			LastPaymentDate:                    element.GetLastPaymentDate(),
			LastStatementIssueDate:             element.GetLastStatementIssueDate(),
			LoanName:                           element.GetLoanName(),
			LoanEndDate:                        element.LoanStatus.GetEndDate(),
			MinimumPaymentAmount:               float64(element.GetMinimumPaymentAmount()),
			NextPaymentDueDate:                 element.GetNextPaymentDueDate(),
			OriginationDate:                    element.GetOriginationDate(),
			OriginationPrincipalAmount:         float64(element.GetOriginationPrincipalAmount()),
			OutstandingInterestAmount:          float64(element.GetOutstandingInterestAmount()),
			PaymentReferenceNumber:             element.GetPaymentReferenceNumber(),
			SequenceNumber:                     element.GetSequenceNumber(),
			YtdInterestPaid:                    *element.YtdInterestPaid.Get(),
			YtdPrincipalPaid:                   *element.YtdPrincipalPaid.Get(),
			LoanType:                           *element.LoanStatus.Type.Get(),
			PslfStatusEstimatedEligibilityDate: element.PslfStatus.GetEstimatedEligibilityDate(),
			PslfStatusPaymentsMade:             element.PslfStatus.GetPaymentsMade(),
			PslfStatusPaymentsRemaining:        element.PslfStatus.GetPaymentsRemaining(),
			RepaymentPlanType:                  repaymentPlan.GetType(),
			RepaymentPlanDescription:           repaymentPlan.GetDescription(),
			ServicerAddressCity:                servisorAddress.GetCity(),
			ServicerAddressPostalCode:          servisorAddress.GetPostalCode(),
			ServicerAddressState:               servisorAddress.GetCity(),
			ServicerAddressStreet:              servisorAddress.GetStreet(),
			ServicerAddressRegion:              servisorAddress.GetRegion(),
			ServicerAddressCountry:             servisorAddress.GetCountry(),
			UserId:                             0,
			Name:                               metadata.accountName,
		}

		accts = append(accts, studentloanAcct)
	}

	return accts, nil
}
