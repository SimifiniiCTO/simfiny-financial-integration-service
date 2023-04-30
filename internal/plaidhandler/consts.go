package plaidhandler

import (
	"github.com/plaid/plaid-go/plaid"
)

// TODO: below should be read from env vars
var (
	PlaidClientName = "simfiny"
	PlaidLanguage   = "en"
	PlaidCountries  = []plaid.CountryCode{
		plaid.COUNTRYCODE_US,
	}
	PlaidProducts = []plaid.Products{
		plaid.PRODUCTS_BALANCE,
		plaid.PRODUCTS_INVESTMENTS,
		plaid.PRODUCTS_LIABILITIES,
		plaid.PRODUCTS_TRANSACTIONS,
	}
)

func PlaidProductStrings() []string {
	items := make([]string, len(PlaidProducts))
	for i, product := range PlaidProducts {
		items[i] = string(product)
	}

	return items
}
