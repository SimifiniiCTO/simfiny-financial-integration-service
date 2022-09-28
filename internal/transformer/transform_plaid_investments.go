package transformer

import (
	"errors"

	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/plaid/plaid-go/plaid"
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
				PlaidAccountID: metadata.accountID,
				AccountSubtype: metadata.accountSubtype,
				AccountType:    metadata.accountType,
				AccountName:    metadata.accountName,
				BalanceID:      metadata.balance,
				SecurityID:     securities,
			}

			accts = append(accts, investmentAcct)
		}
	}

	return accts, nil
}

func transformAndMapSecuritiesToAcctType(securities *[]plaid.Security) map[string][]*schema.Security {
	// iterate and transform securities set
	var acctTypeToSecuritiesMap map[string][]*schema.Security
	for _, security := range *securities {
		key := security.GetType()

		if val, ok := acctTypeToSecuritiesMap[key]; ok {
			sec := &schema.Security{
				ClosePrice:       uint64(security.GetClosePrice()),
				IsCashEquivalent: security.GetIsCashEquivalent(),
				CurrencyCode:     security.GetIsoCurrencyCode(),
				SecurityID:       security.GetSecurityId(),
				TickerSymbol:     security.GetTickerSymbol(),
				SecurityType:     security.GetType(),
				SecurityName:     security.GetName(),
			}

			val = append(val, sec)
			acctTypeToSecuritiesMap[key] = val
		}
	}

	return acctTypeToSecuritiesMap
}
