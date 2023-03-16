package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// GetMortageAccount implements apiv1.FinancialServiceServer
func (*Server) GetMortageAccount(context.Context, *proto.GetMortageAccountRequest) (*proto.GetMortageAccountResponse, error) {
	panic("unimplemented")
}
