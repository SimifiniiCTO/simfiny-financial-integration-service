package plaidhandler

import (
	"github.com/plaid/plaid-go/v12/plaid"
)

const (
	DEPOSITORY  = string(plaid.ACCOUNTTYPE_DEPOSITORY)
	CREDIT      = string(plaid.ACCOUNTTYPE_CREDIT)
	LOAN        = string(plaid.ACCOUNTTYPE_LOAN)
	INVESTMENTS = string(plaid.ACCOUNTTYPE_INVESTMENT)
)

// filterAccounts filters accounts of a type
func filterAccounts(accts []plaid.AccountBase, acctType string) []plaid.AccountBase {
	result := make([]plaid.AccountBase, 0)

	for _, acct := range accts {
		if acct.Type == plaid.AccountType(acctType) {
			result = append(result, acct)
		}
	}

	return result
}

// filterAccountsOfTwoType filters accounts of two types
func filterAccountsOfTwoType(accts []plaid.AccountBase, acctType1, acctType2 string) []plaid.AccountBase {
	result := make([]plaid.AccountBase, 0)

	for _, acct := range accts {
		if acct.Type == plaid.AccountType(acctType1) || acct.Type == plaid.AccountType(acctType2) {
			result = append(result, acct)
		}
	}

	return result
}
