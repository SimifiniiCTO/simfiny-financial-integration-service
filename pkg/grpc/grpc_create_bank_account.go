package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// CreateBankAccount implements apiv1.FinancialServiceServer
func (*Server) CreateBankAccount(context.Context, *proto.CreateBankAccountRequest) (*proto.CreateBankAccountResponse, error) {
	panic("unimplemented")
}
