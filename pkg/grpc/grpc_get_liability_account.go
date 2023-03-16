package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetLiabilityAccount implements apiv1.FinancialServiceServer
func (*Server) GetLiabilityAccount(context.Context, *proto.GetLiabilityAccountRequest) (*proto.GetLiabilityAccountResponse, error) {
	panic("unimplemented")
}
