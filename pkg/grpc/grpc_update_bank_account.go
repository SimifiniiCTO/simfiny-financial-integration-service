package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// UpdateBankAccount implements apiv1.FinancialServiceServer
func (*Server) UpdateBankAccount(context.Context, *proto.UpdateBankAccountRequest) (*proto.UpdateBankAccountResponse, error) {
	panic("unimplemented")
}
