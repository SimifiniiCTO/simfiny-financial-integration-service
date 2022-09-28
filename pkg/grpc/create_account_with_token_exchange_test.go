package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/helper"
	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/plaidhandler"
	proto "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

type CreateAccountWithTokenExchangeRpcScenario struct {
	scenarioName              string
	shouldGeneratePublicToken bool
	shouldErrorOccur          bool
	expectedError             error
	userID                    uint64
	publicToken               string
}

func setupTestPreconditions() (*plaid.SandboxPublicTokenCreateResponse, error) {
	ctx := context.Background()
	pw, err := plaidhandler.GetPlaidWrapperForTest()
	if err != nil {
		return nil, err
	}

	resp, err := pw.GetPlublicTokenForSandboxAcct(ctx)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func CreateAccountWithTokenExchangeRpcScenarios() ([]CreateAccountWithTokenExchangeRpcScenario, error) {
	resp, err := setupTestPreconditions()
	if err != nil {
		return nil, err
	}

	return []CreateAccountWithTokenExchangeRpcScenario{
		{
			// success condition - create an account with token exchange
			scenarioName:              "[success scenario] - create a new account via token exchange",
			shouldGeneratePublicToken: true,
			shouldErrorOccur:          false,
			expectedError:             nil,
			userID:                    uint64(helper.GenerateRandomId(10000, 20000)),
			publicToken:               resp.PublicToken,
		},
		{
			// failure condition - attempt to create an account with no user id specified
			scenarioName:              "[failure scenario] - no userID provided",
			shouldGeneratePublicToken: true,
			shouldErrorOccur:          true,
			expectedError:             errors.New("invalid request object. user id cannot be 0"),
			userID:                    0,
			publicToken:               resp.PublicToken,
		},
		{
			// failure condition - attempt to create an account with no public token provided
			scenarioName:              "[failure scenario] - no userID provided",
			shouldGeneratePublicToken: true,
			shouldErrorOccur:          true,
			expectedError:             errors.New("invalid request object. public token cannot be empty"),
			userID:                    uint64(helper.GenerateRandomId(10000, 20000)),
			publicToken:               "",
		},
	}, nil
}

func TestCreateAccountWithTokenExchangeRpcOperation(t *testing.T) {
	ctx := context.Background()
	conn := MockGRPCService(ctx)
	defer conn.Close()

	scenarios, err := CreateAccountWithTokenExchangeRpcScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		c := proto.NewFinancialIntegrationServiceClient(conn)
		req := &proto.CreateAccountTokenExchangeRequest{
			PublicToken: scenario.publicToken,
			UserID:      scenario.userID,
		}

		resp, err := c.CreateAccountWithTokenExchange(ctx, req)
		if err != nil {
			if scenario.shouldErrorOccur {
				assert.Contains(t, err.Error(), scenario.expectedError.Error())
			} else {
				t.Errorf("obtained error but not expected - %s", err.Error())
			}
		} else {
			if scenario.shouldErrorOccur {
				t.Errorf("expected error %s but none occured", scenario.expectedError.Error())
			}
		}

		if !scenario.shouldErrorOccur {
			assert.NotNil(t, resp.GetVirtualAccountID())
		}
	}
}
