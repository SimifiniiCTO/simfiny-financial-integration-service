package database

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/helper"
	schema "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/stretchr/testify/assert"
)

type deleteAccountScenario struct {
	scenarioName                string
	shouldErrorOccur            bool
	shouldAccountBeCreatedFirst bool
	account                     *schema.VirtualAccount
	expectedError               error
	context                     context.Context
}

// deleteAccountScenarios returns a set of scenarios to test the delete account operation
func deleteAccountScenarios() []deleteAccountScenario {
	acctWithRandomVAcctID := generateRandomizedAccount()
	acctWithRandomVAcctID.Id = uint64(helper.GenerateRandomId(10000, 200000))

	return []deleteAccountScenario{
		{
			// success condition - deletion of an existing account
			scenarioName:                "delete existing account",
			shouldErrorOccur:            false,
			shouldAccountBeCreatedFirst: true,
			expectedError:               nil,
			account:                     generateRandomizedAccount(),
			context:                     context.Background(),
		},
		{
			// failure condition - deletion of an account that does not exist
			scenarioName:                "delete existing account - invalid id (0)",
			shouldErrorOccur:            true,
			shouldAccountBeCreatedFirst: false,
			expectedError:               errors.New("invalid input argument. vAcctID cannot be 0"),
			account:                     generateRandomizedAccount(),
			context:                     context.Background(),
		},
		{
			// failure condition - deletion of an account that does not exist
			scenarioName:                "delete existing account - id does not exist in db",
			shouldErrorOccur:            true,
			shouldAccountBeCreatedFirst: false,
			expectedError:               errors.New("record not found"),
			account:                     acctWithRandomVAcctID,
			context:                     context.Background(),
		},
	}
}

func TestDeleteAccountOperation(t *testing.T) {
	scenarios := deleteAccountScenarios()
	reset()

	for _, scenario := range scenarios {
		var account *schema.VirtualAccount = scenario.account
		var err error
		if scenario.shouldAccountBeCreatedFirst {
			accessToken := helper.GenerateRandomString(10)
			account, err = Conn.CreateVirtualAccount(scenario.context, account, accessToken)
			if err != nil {
				t.Errorf("failed to create test account as precondition - %s", err.Error())
			}

			assert.NotNil(t, account)
		}

		err = Conn.DeactivateVirtualAccount(scenario.context, account.Id)
		if err != nil {
			if scenario.shouldErrorOccur {
				assert.Equal(t, err, scenario.expectedError)
			} else {
				t.Errorf("obtained unexpected error - %s", err.Error())
			}
		}

		if scenario.shouldErrorOccur && err == nil {
			t.Errorf("expected error to occur but none did")
		}

		if !scenario.shouldErrorOccur {
			assert.Nil(t, err)
		}
	}

}
