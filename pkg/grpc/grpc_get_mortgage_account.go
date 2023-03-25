package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetMortgageAccount implements apiv1.FinancialServiceServer
func (*Server) GetMortgageAccount(context.Context, *proto.GetMortgageAccountRequest) (*proto.GetMortgageAccountResponse, error) {
	panic("unimplemented")
}
