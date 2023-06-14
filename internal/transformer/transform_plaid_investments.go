package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func transformPlaidInvestmentObject(securities *[]plaid.Security,
	holdings *[]plaid.Holding, acctIDToTypeMap map[string]*accountMetadata) ([]*schema.InvestmentAccount, error) {
	if securities == nil {
		return nil, errors.New("invalid input argument. securities cannot be nil")
	}

	if holdings == nil {
		return nil, errors.New("invalid input argument. holdings cannot be nil")
	}

	if acctIDToTypeMap == nil {
		return nil, errors.New("invalid input argument. account id to type map is nil")
	}
	accts := make([]*schema.InvestmentAccount, 0)
	acctToSecuritiesMap := transformAndMapSecuritiesToAcctType(securities)

	for _, element := range *holdings {
		metadata := acctIDToTypeMap[element.AccountId]
		if securities, ok := acctToSecuritiesMap[metadata.accountType]; ok {
			investmentAcct := &schema.InvestmentAccount{
				Id:             0,
				UserId:         0,
				Name:           metadata.accountName,
				Number:         "",
				Type:           metadata.accountType,
				Balance:        0,
				CurrentFunds:   0,
				BalanceLimit:   0,
				PlaidAccountId: metadata.accountID,
				Subtype:        metadata.accountType,
				Holdings:       []*schema.InvesmentHolding{},
				Securities:     securities,
			}
			accts = append(accts, investmentAcct)
		}
	}

	return accts, nil
}

func transformAndMapSecuritiesToAcctType(securities *[]plaid.Security) map[string][]*schema.InvestmentSecurity {
	// iterate and transform securities set
	var acctTypeToSecuritiesMap map[string][]*schema.InvestmentSecurity
	for _, security := range *securities {
		key := security.GetType()

		if val, ok := acctTypeToSecuritiesMap[key]; ok {
			sec := TransformSinglePlaidInvestmentSecurity(security)
			val = append(val, sec)
			acctTypeToSecuritiesMap[key] = val
		}
	}

	return acctTypeToSecuritiesMap
}
