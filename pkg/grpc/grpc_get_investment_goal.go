package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetInvestmentAcccount implements apiv1.FinancialServiceServer
func (*Server) GetInvestmentAcccount(context.Context, *proto.GetInvestmentAcccountRequest) (*proto.GetInvestmentAcccountResponse, error) {
	panic("unimplemented")
}
