package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

type GetLinkTokenRpcScenario struct {
	scenarioName     string
	shouldErrorOccur bool
	expectedError    error
	userID           uint64
}

func GetLinkTokenRpcScenarios() ([]GetLinkTokenRpcScenario, error) {
	return []GetLinkTokenRpcScenario{
		{
			scenarioName:     "[success scenario] - get an account that already exists",
			shouldErrorOccur: false,
			expectedError:    nil,
			userID:           uint64(helper.GenerateRandomId(10000, 20000)),
		},
		{
			scenarioName:     "[failure scenario] - get an account that doesn't already exist",
			shouldErrorOccur: true,
			expectedError:    errors.New("invalid request object. user id cannot be 0"),
			userID:           0,
		},
	}, nil
}

func TestGetLinkTokenRpcOperation(t *testing.T) {
	ctx := context.Background()
	conn := MockGRPCService(ctx)
	defer conn.Close()

	scenarios, err := GetLinkTokenRpcScenarios()
	assert.Nil(t, err)

	for _, scenario := range scenarios {
		c := proto.NewFinancialIntegrationServiceClient(conn)

		req := &proto.InitiateAccountSetupRequest{
			UserID: scenario.userID,
		}

		linkToken, err := c.InitiateAccountSetupAndGetLinkToken(ctx, req)
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
			assert.NotNil(t, linkToken)
		}
	}
}

func TestServer_InitiateAccountSetupAndGetLinkToken(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *proto.InitiateAccountSetupRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.InitiateAccountSetupResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.InitiateAccountSetupAndGetLinkToken(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.InitiateAccountSetupAndGetLinkToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.InitiateAccountSetupAndGetLinkToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_performnOperationAgainstPlaidSandboxEnv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *plaid.SandboxPublicTokenCreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.performnOperationAgainstPlaidSandboxEnv(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.performnOperationAgainstPlaidSandboxEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.performnOperationAgainstPlaidSandboxEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_performnOperationAgainstPlaidProdEnv(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint64
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    plaid.LinkTokenCreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.performnOperationAgainstPlaidProdEnv(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.performnOperationAgainstPlaidProdEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.performnOperationAgainstPlaidProdEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
