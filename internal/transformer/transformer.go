package transformer

import (
	"github.com/plaid/plaid-go/v12/plaid"
)

// The above type defines the metadata of an account, including its subtype, type, name, ID, available
// and current funds, currency code, balance limit, and account number.
// @property {string} accountSubtype - The subtype of the account, which provides additional
// information about the type of account. For example, a checking account may have subtypes such as
// personal checking, business checking, or student checking.
// @property {string} accountType - The type of the account, such as checking, savings, credit card,
// etc.
// @property {string} accountName - The name of the account.
// @property {string} accountID - The unique identifier for the account. It is used to distinguish one
// account from another.
// @property {float64} availableFunds - The available funds property represents the amount of money
// that is currently available for use in the account. This may be different from the current funds
// property, which represents the total amount of money in the account, including any pending
// transactions or holds. The available funds property is important for determining how much money a
// user
// @property {float64} currentFunds - currentFunds is a property of the accountMetadata struct that
// represents the current balance of the account. It is a float64 data type, which means it can store
// decimal values.
// @property {string} currencyCode - The currency code is a three-letter code that represents the
// currency used in the account. For example, USD for US dollars, EUR for euros, GBP for British
// pounds, etc.
// @property {uint64} balanceLimit - The balanceLimit property represents the maximum amount of funds
// that can be held in the account at any given time. It is usually set by the financial institution or
// the account holder.
// @property {string} accountNumber - The account number is a unique identifier assigned to a bank
// account. It is used to identify the account when making transactions or inquiries.
type accountMetadata struct {
	accountSubtype string
	accountType    string
	accountName    string
	accountID      string
	availableFunds float64
	currentFunds   float64
	currencyCode   string
	balanceLimit   uint64
	accountNumber  string
}

// GetAccountMetadata takes in a list of Plaid account objects and returns a map of account metadata,
// including account type, balance, and account number.
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
			accountNumber:  el.GetMask(),
		}
	}
	return acctIDtoAcctObj
}
