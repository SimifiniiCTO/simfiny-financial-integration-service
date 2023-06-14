package plaidhandler

import (
	"errors"
	"reflect"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

type PopulateAccountTestScenario struct {
	scenarioName     string
	shouldErrorOccur bool
	accessToken      *string
	expectedError    error
}

// findAccountByIdScenarios returns a set of scenarios to test the account's existence based on provided email
func getPopulateAccountTestScenarios() ([]PopulateAccountTestScenario, *PlaidWrapper, error) {
	accessToken, err := plaidTestClient.getAccessTokenForSandboxAcct()
	if err != nil {
		return nil, nil, err
	}

	return []PopulateAccountTestScenario{
		{
			// success condition: valid access token
			scenarioName:     "[success condition]: get an account deposits with valid access token",
			shouldErrorOccur: false,
			accessToken:      &accessToken,
		},
		{
			// failure condition: access token is invalid
			scenarioName:     "[failure condition]: get an account deposits with invalid access token",
			shouldErrorOccur: true,
			expectedError:    errors.New("invalid input argument. access token cannot be empty"),
			accessToken:      nil,
		},
	}, plaidTestClient, nil
}

func processError(err error, scenario PopulateAccountTestScenario, t *testing.T) {
	if err != nil {
		if scenario.shouldErrorOccur {
			assert.Equal(t, scenario.expectedError, err)
		} else {
			t.Errorf("obtained error but not expected - %s", err.Error())
		}
	}
}

func Test_filterAccounts(t *testing.T) {
	type args struct {
		accts    []plaid.AccountBase
		acctType string
	}
	tests := []struct {
		name string
		args args
		want []plaid.AccountBase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterAccounts(tt.args.accts, tt.args.acctType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterAccountsOfTwoType(t *testing.T) {
	type args struct {
		accts     []plaid.AccountBase
		acctType1 string
		acctType2 string
	}
	tests := []struct {
		name string
		args args
		want []plaid.AccountBase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterAccountsOfTwoType(tt.args.accts, tt.args.acctType1, tt.args.acctType2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterAccountsOfTwoType() = %v, want %v", got, tt.want)
			}
		})
	}
}
