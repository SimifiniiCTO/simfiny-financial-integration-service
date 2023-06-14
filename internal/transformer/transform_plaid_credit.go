package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

// transformPlaidCreditObject transforms a plaid credit card liabilities object to the internal credit account object
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
		metadata := acctIDToTypeMap[creditAcctId]
		creditAcct := &schema.CreditAccount{
			Id:                     0,
			UserId:                 0,
			Name:                   metadata.accountName,
			Number:                 "",
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
			Aprs:                   transformPlaidApr(liability.Aprs),
			LastStatementBalance:   liability.GetLastStatementBalance(),
			MinimumPaymentAmount:   liability.GetMinimumPaymentAmount(),
			NextPaymentDueDate:     liability.GetNextPaymentDueDate(),
		}

		accts = append(accts, creditAcct)
	}

	return accts, nil
}

// transformPlaidApr transforms plaid APR objects to internal plaid APR representation
func transformPlaidApr(plaidAPR []plaid.APR) []*schema.Apr {
	if plaidAPR == nil {
		return []*schema.Apr{}
	}

	aprs := make([]*schema.Apr, 0)
	for _, apr := range plaidAPR {
		transformedApr := &schema.Apr{
			Id:                   0,
			Percentage:           apr.AprPercentage,
			Type:                 apr.AprType,
			BalanceSubjectToApr:  *apr.BalanceSubjectToApr.Get(),
			InterestChargeAmount: float64(apr.GetInterestChargeAmount()),
		}

		aprs = append(aprs, transformedApr)
	}

	return aprs
}
