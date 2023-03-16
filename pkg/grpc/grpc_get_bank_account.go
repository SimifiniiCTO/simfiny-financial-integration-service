package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// GetBankAccount implements apiv1.FinancialServiceServer
func (*Server) GetBankAccount(context.Context, *proto.GetBankAccountRequest) (*proto.GetBankAccountResponse, error) {
	panic("unimplemented")
}
