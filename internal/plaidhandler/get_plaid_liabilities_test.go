package plaidhandler

import (
	"errors"
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
