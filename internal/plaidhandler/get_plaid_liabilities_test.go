package plaidhandler

import (
	"context"
	"errors"
	"testing"

	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

type plaidLiabilityScenarios struct {
	scenarioName     string
	shouldErrorOccur bool
	accessToken      *string
	expectedError    error
}

func getplaidLiabilityScenarios() ([]plaidLiabilityScenarios, *PlaidWrapper, error) {
	accessToken, err := plaidTestClient.getAccessTokenForSandboxAcct()
	if err != nil {
		return nil, nil, err
	}

	return []plaidLiabilityScenarios{
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

func TestGetPlaidAccountLiabilityOperation(t *testing.T) {
	scenarios, p, err := getplaidLiabilityScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		liabilities, err := p.getPlaidLiabilities(context.Background(), scenario.accessToken)
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
			assertAccountsOfLiabilityType(t, liabilities)
		}
	}
}

func assertAccountsOfLiabilityType(t *testing.T, d *Liabilities) {
	assert.NotNil(t, d)
	assert.NotNil(t, d.Accounts)
	for _, acct := range d.Accounts {
		assert.NotNil(t, acct)
		if acct.Type != plaid.AccountType(CREDIT) && acct.Type != plaid.AccountType(LOAN) {
			t.Errorf("account can only be of credit or loan type")
		}
	}
}
