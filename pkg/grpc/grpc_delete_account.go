package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// DeleteBankAccount implements apiv1.FinancialServiceServer
func (*Server) DeleteBankAccount(context.Context, *proto.DeleteBankAccountRequest) (*proto.DeleteBankAccountResponse, error) {
	panic("unimplemented")
}
