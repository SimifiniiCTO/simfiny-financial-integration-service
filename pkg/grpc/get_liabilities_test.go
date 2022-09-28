package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
	"github.com/stretchr/testify/assert"
)

type GetLiabilitiesRpcScenario struct {
	scenarioName        string
	shouldErrorOccur    bool
	expectedError       error
	userID              uint64
	publicToken         string
	virtualAccountID    uint64
	shouldCreateAccount bool
}

func GetLiabilitiesRpcScenarios() ([]GetLiabilitiesRpcScenario, error) {
	resp, err := setupTestPreconditions()
	if err != nil {
		return nil, err
	}

	return []GetLiabilitiesRpcScenario{
		{
			scenarioName:        "[success scenario] - get an account that already exists",
			shouldErrorOccur:    false,
			expectedError:       nil,
			userID:              uint64(helper.GenerateRandomId(10000, 20000)),
			publicToken:         resp.PublicToken,
			shouldCreateAccount: true,
			virtualAccountID:    0,
		},
		{
			scenarioName:        "[failure scenario] - get an account that doesn't already exist",
			shouldErrorOccur:    true,
			expectedError:       errors.New("record not found"),
			userID:              uint64(helper.GenerateRandomId(10000, 20000)),
			publicToken:         resp.PublicToken,
			shouldCreateAccount: false,
			virtualAccountID:    uint64(helper.GenerateRandomId(10000, 20000)),
		},
	}, nil
}

func TestGetLiabilitiesRpcOperation(t *testing.T) {
	ctx := context.Background()
	conn := MockGRPCService(ctx)
	defer conn.Close()

	scenarios, err := GetLiabilitiesRpcScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		c := proto.NewFinancialIntegrationServiceClient(conn)

		if scenario.shouldCreateAccount {
			req := &proto.CreateAccountTokenExchangeRequest{
				PublicToken: scenario.publicToken,
				UserID:      scenario.userID,
			}

			resp, err := c.CreateAccountWithTokenExchange(ctx, req)
			assert.Nil(t, err)

			scenario.virtualAccountID = resp.GetVirtualAccountID()
		}

		req := &proto.GetLiabilitiesRequest{
			VirtualAccountID: scenario.virtualAccountID,
			UserID:           scenario.userID,
		}

		resp, err := c.GetLiabilities(ctx, req)
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
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Liabilities)
		}
	}
}
