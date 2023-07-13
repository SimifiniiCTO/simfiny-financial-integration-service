package plaidhandler

import (
	"github.com/plaid/plaid-go/v14/plaid"
)

// TODO: below should be read from env vars
var (
	PlaidClientName = "simfiny"
	PlaidLanguage   = "en"
	PlaidCountries  = []plaid.CountryCode{
		plaid.COUNTRYCODE_US,
	}
)
