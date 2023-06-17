package transformer

import (
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestGetAccountMetadata(t *testing.T) {
	// Prepare sample data for testing
	accounts := []plaid.AccountBase{
		{
			AccountId: "1",
			Subtype:   *plaid.NewNullableAccountSubtype(plaid.ACCOUNTSUBTYPE_BUSINESS.Ptr()),
			Type:      plaid.AccountType("depository"),
			Name:      "Savings Account",
			Balances: plaid.AccountBalance{
				Available:       *plaid.NewNullableFloat64(pointer.Float64P(float64(1000))),
				Current:         *plaid.NewNullableFloat64(pointer.Float64P(float64(2000))),
				IsoCurrencyCode: *plaid.NewNullableString(pointer.StringP("USD")),
				Limit:           *plaid.NewNullableFloat64(pointer.Float64P(float64(5000))),
			},
			Mask: *plaid.NewNullableString(pointer.StringP("1234")),
		},
		{
			AccountId: "2",
			Subtype:   *plaid.NewNullableAccountSubtype(plaid.ACCOUNTSUBTYPE_CREDIT_CARD.Ptr()),
			Type:      plaid.AccountType("credit"),
			Name:      "Credit Card",
			Balances: plaid.AccountBalance{
				Available:       *plaid.NewNullableFloat64(pointer.Float64P(float64(500))),
				Current:         *plaid.NewNullableFloat64(pointer.Float64P(float64(-1000))),
				IsoCurrencyCode: *plaid.NewNullableString(pointer.StringP("USD")),
				Limit:           *plaid.NewNullableFloat64(pointer.Float64P(float64(2000))),
			},
			Mask: *plaid.NewNullableString(pointer.StringP("5678")),
		},
	}

	expectedResult := map[string]*accountMetadata{
		"1": {
			accountSubtype: "business",
			accountType:    "depository",
			accountID:      "1",
			accountName:    "Savings Account",
			availableFunds: 1000,
			currentFunds:   2000,
			currencyCode:   "USD",
			balanceLimit:   5000,
			accountNumber:  "1234",
		},
		"2": {
			accountSubtype: "credit card",
			accountType:    "credit",
			accountID:      "2",
			accountName:    "Credit Card",
			availableFunds: 500,
			currentFunds:   -1000,
			currencyCode:   "USD",
			balanceLimit:   2000,
			accountNumber:  "5678",
		},
	}

	// Call the function being tested
	result := GetAccountMetadata(&accounts)

	// Compare the result with the expected value
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Unexpected result. Got: %v, want: %v", result, expectedResult)
	}
}
