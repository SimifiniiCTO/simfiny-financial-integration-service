package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/plaid/plaid-go/v14/plaid"
)

// TransformPlaidCredit transforms Plaid credit card liabilities into a list of credit accounts with associated
// metadata.
func TransformPlaidCredit(creditCardliabilities *[]plaid.CreditCardLiability, acctIDToTypeMap map[string]*accountMetadata) ([]*schema.
	CreditAccount, error) {
	if creditCardliabilities == nil {
		return nil, errors.New("invalid input argument. credit card liability cannot be nil")
	}

	if acctIDToTypeMap == nil {
		return nil, errors.New("invalid input argument. account id to type map is nil")
	}

	accts := make([]*schema.CreditAccount, 0)
	for _, liability := range *creditCardliabilities {
		creditAcctId := *liability.AccountId.Get()

		metadata, ok := acctIDToTypeMap[creditAcctId]
		if !ok {
			return nil, errors.New("invalid input argument. account id to type map does not contain liability account id")
		}

		aprs, err := NewAPR(&liability.Aprs)
		if err != nil {
			return nil, err
		}

		creditAcct := &schema.CreditAccount{
			Id:                     0,
			UserId:                 0,
			Name:                   metadata.accountName,
			Number:                 metadata.accountNumber,
			Type:                   metadata.accountType,
			Balance:                float32(metadata.balanceLimit),
			CurrentFunds:           metadata.currentFunds,
			BalanceLimit:           metadata.balanceLimit,
			PlaidAccountId:         liability.GetAccountId(),
			Subtype:                metadata.accountSubtype,
			IsOverdue:              liability.GetIsOverdue(),
			LastPaymentAmount:      liability.GetLastPaymentAmount(),
			LastPaymentDate:        liability.GetLastPaymentDate(),
			LastStatementIssueDate: liability.GetLastStatementIssueDate(),
			MinimumAmountDueDate:   0,
			NextPaymentDate:        liability.GetNextPaymentDueDate(),
			Aprs:                   aprs,
			LastStatementBalance:   liability.GetLastStatementBalance(),
			MinimumPaymentAmount:   liability.GetMinimumPaymentAmount(),
			NextPaymentDueDate:     liability.GetNextPaymentDueDate(),
		}

		accts = append(accts, creditAcct)
	}

	return accts, nil
}
