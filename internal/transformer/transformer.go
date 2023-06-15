package transformer

import (
	"github.com/plaid/plaid-go/v12/plaid"
)

type accountMetadata struct {
	accountSubtype string
	accountType    string
	accountName    string
	accountID      string
	availableFunds float64
	currentFunds   float64
	currencyCode   string
	balanceLimit   uint64
}

func GetAccountMetadata(accounts *[]plaid.AccountBase) map[string]*accountMetadata {
	acctIDtoAcctObj := make(map[string]*accountMetadata)
	// iterate over account object and store account id as key and account subtype + type as value
	for _, el := range *accounts {
		balance := el.GetBalances()
		acctIDtoAcctObj[el.AccountId] = &accountMetadata{
			accountSubtype: string(el.GetSubtype()),
			accountType:    string(el.GetType()),
			accountID:      el.GetAccountId(),
			accountName:    el.GetName(),
			availableFunds: float64(balance.GetAvailable()),
			currentFunds:   float64(balance.GetCurrent()),
			currencyCode:   balance.GetIsoCurrencyCode(),
			balanceLimit:   uint64(balance.GetLimit()),
		}
	}
	return acctIDtoAcctObj
}
