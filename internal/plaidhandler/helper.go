package plaidhandler

import (
	"github.com/plaid/plaid-go/v12/plaid"
)

const (
	DEPOSITORY  = "depository"
	CREDIT      = "credit"
	LOAN        = "loan"
	INVESTMENTS = "investment"
)

func filterAccounts(accts []plaid.AccountBase, acctType string) []plaid.AccountBase {
	result := make([]plaid.AccountBase, 0)

	for _, acct := range accts {
		if acct.Type == plaid.AccountType(acctType) {
			result = append(result, acct)
		}
	}

	return result
}

func filterAccountsOfTwoType(accts []plaid.AccountBase, acctType1, acctType2 string) []plaid.AccountBase {
	result := make([]plaid.AccountBase, 0)

	for _, acct := range accts {
		if acct.Type == plaid.AccountType(acctType1) || acct.Type == plaid.AccountType(acctType2) {
			result = append(result, acct)
		}
	}

	return result
}
