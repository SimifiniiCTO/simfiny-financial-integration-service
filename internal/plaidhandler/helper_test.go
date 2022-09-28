package plaidhandler

import (
	"context"
	"errors"
	"testing"

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
	accessToken, err := plaidTestClient.GetAccessTokenForSandboxAcct()
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

func TestPopulateAcountFromPlaidOperation(t *testing.T) {
	scenarios, p, err := getPopulateAccountTestScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		deposits, err := p.getPlaidDeposit(context.Background(), scenario.accessToken)
		processError(err, scenario, t)

		liabilities, err := p.getPlaidLiabilities(context.Background(), scenario.accessToken)
		processError(err, scenario, t)

		investments, err := p.getPlaidInvestmentHoldings(context.Background(), scenario.accessToken)
		processError(err, scenario, t)

		vAcct, err := p.populateVirtualAccount(context.Background(), liabilities, investments, deposits)
		if scenario.shouldErrorOccur && err == nil {
			t.Errorf("expected error to occur but none did")
		}

		if !scenario.shouldErrorOccur {
			assert.NotNil(t, vAcct)
		}
	}
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
