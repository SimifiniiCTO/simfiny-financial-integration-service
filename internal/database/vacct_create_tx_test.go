package database

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/stretchr/testify/assert"
)

type createAccountScenario struct {
	scenarioName     string
	shouldErrorOccur bool
	account          *schema.VirtualAccount
	expectedError    error
	createDuplicate  bool
	context          context.Context
}

// createAccountTestScenarios returns a set of scenarios to test the create account operation
func createAccountTestScenarios() []createAccountScenario {
	testAcct := generateRandomizedAccount()
	invalidAccMissingUserID := generateRandomizedAccount()
	invalidAccMissingUserID.UserID = 0

	invalidAccMissingAccessToken := generateRandomizedAccount()
	invalidAccMissingAccessToken.AccessToken = ""

	return []createAccountScenario{
		{
			// success condition - create a new virtual account
			scenarioName:     "create new virtual account",
			shouldErrorOccur: false,
			account:          testAcct,
			expectedError:    nil,
			context:          context.Background(),
		},
		{
			// failure condition - create a an invalid virtual account (nil)
			scenarioName:     "create an invalid virtual account of type nil",
			shouldErrorOccur: true,
			account:          nil,
			expectedError:    errors.New("invalid input argument. vAcct cannot be nil"),
			context:          context.Background(),
		},
		{
			// failure condition - create a an invalid virtual account (missing necessary fields)
			scenarioName:     "create an invalid virtual account with missing userID field",
			shouldErrorOccur: true,
			account:          invalidAccMissingUserID,
			expectedError:    errors.New("userID cannot be 0"),
			context:          context.Background(),
		},
		{
			// failure condition - create a an invalid virtual account (missing necessary fields)
			scenarioName:     "create an invalid virtual account with missing accestoken field",
			shouldErrorOccur: true,
			account:          invalidAccMissingAccessToken,
			expectedError:    errors.New("accesstoken cannot be empty"),
			context:          context.Background(),
		},
		{
			// failure condition - create a duplicate account
			scenarioName:     "create a duplicate virtual account",
			shouldErrorOccur: true,
			account:          generateRandomizedAccount(),
			expectedError:    errors.New("account already exists"),
			createDuplicate:  true,
			context:          context.Background(),
		},
	}
}

func TestCreateAccountOperation(t *testing.T) {
	scenarios := createAccountTestScenarios()
	reset()

	for _, scenario := range scenarios {
		accessToken := helper.GenerateRandomString(10)
		var err error

		if scenario.createDuplicate {
			account, err := Conn.CreateVirtualAccount(scenario.context, scenario.account, accessToken)
			assert.Nil(t, err)
			scenario.account = account
		}

		account, err := Conn.CreateVirtualAccount(scenario.context, scenario.account, accessToken)
		if err != nil {
			if scenario.shouldErrorOccur {
				assert.Equal(t, err, scenario.expectedError)
			} else {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}
		} else {
			if scenario.shouldErrorOccur {
				t.Errorf("expected error %s but none occured", scenario.expectedError.Error())
			}
		}

		if !scenario.shouldErrorOccur {
			assert.NotNil(t, account)
		}
	}
}
