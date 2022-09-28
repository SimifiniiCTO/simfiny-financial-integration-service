package plaidhandler

import (
	"context"
	"errors"
	"testing"

	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

type plaidDepositScenarios struct {
	scenarioName     string
	shouldErrorOccur bool
	accessToken      *string
	expectedError    error
}

// findAccountByIdScenarios returns a set of scenarios to test the account's existence based on provided email
func getPlaidDepositScenarios() ([]plaidDepositScenarios, *PlaidWrapper, error) {
	accessToken, err := plaidTestClient.GetAccessTokenForSandboxAcct()
	if err != nil {
		return nil, nil, err
	}

	return []plaidDepositScenarios{
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

func TestGetPlaidDepositoryAccountsOperation(t *testing.T) {
	scenarios, p, err := getPlaidDepositScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		deposits, err := p.getPlaidDeposit(context.Background(), scenario.accessToken)
		if err != nil {
			if scenario.shouldErrorOccur {
				assert.Equal(t, err, scenario.expectedError)
			} else {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}
		}

		if scenario.shouldErrorOccur && err == nil {
			t.Errorf("expected error to occur but none did")
		}

		if !scenario.shouldErrorOccur {
			assertAccountsOfDepositoryType(t, deposits)
		}
	}
}

func assertAccountsOfDepositoryType(t *testing.T, d *Deposits) {
	assert.NotNil(t, d)
	for _, acct := range d.Accounts {
		assert.NotNil(t, acct)
		assert.Equal(t, plaid.AccountType(DEPOSITORY), acct.Type)
	}
}
