package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/plaid"
)

// transformPlaidCreditObject transforms a plaid credit card liabilities object to the internal credit account object
func transformPlaidCreditObject(creditCardliabilities *[]plaid.CreditCardLiability, acctIDToTypeMap map[string]*accountMetadata) ([]*schema.
	CreditAccount, error) {
	if creditCardliabilities == nil {
		return nil, errors.New("invalid input argument. credit card liability cannot be nil")
	}

	if acctIDToTypeMap == nil {
		return nil, errors.New("invalid input argument. account id to type map is nil")
	}

	accts := make([]*schema.CreditAccount, 0)
	for _, element := range *creditCardliabilities {
		aprs, err := transformPlaidApr(&element.Aprs)
		if err != nil {
			return nil, err
		}

		creditAcctId := *element.AccountId.Get()
		metadata := acctIDToTypeMap[creditAcctId]
		creditAcct := &schema.CreditAccount{
			PlaidAccountID:         element.GetAccountId(),
			AccountSubtype:         metadata.accountSubtype,
			AccountType:            metadata.accountType,
			IsOverdue:              element.GetIsOverdue(),
			LastPaymentAmount:      float64(element.LastPaymentAmount),
			LastPaymentDate:        element.GetLastPaymentDate(),
			LastStatementIssueDate: element.LastStatementIssueDate,
			MinimumPaymentAmount:   float64(element.GetMinimumPaymentAmount()),
			NextPaymentDueDate:     element.GetNextPaymentDueDate(),
			Aprs:                   aprs,
			BalanceID:              metadata.balance,
			AccountName:            metadata.accountName,
		}

		accts = append(accts, creditAcct)
	}

	return accts, nil
}

// transformPlaidApr transforms plaid APR objects to internal plaid APR representation
func transformPlaidApr(plaidAPR *[]plaid.APR) ([]*schema.APR, error) {
	if plaidAPR == nil {
		return nil, errors.New("invalid input argument. apr cannot be nil")
	}

	aprs := make([]*schema.APR, 0)
	for _, element := range *plaidAPR {
		transformedApr := &schema.APR{
			APRPercentage:        float64(element.GetAprPercentage()),
			APRType:              element.GetAprType(),
			BalanceSubjectToAPR:  float64(element.GetBalanceSubjectToApr()),
			InterestChargeAmount: float64(element.GetInterestChargeAmount()),
		}

		aprs = append(aprs, transformedApr)
	}

	return aprs, nil
}
