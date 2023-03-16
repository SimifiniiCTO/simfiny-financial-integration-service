package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// GetAllBudgets implements apiv1.FinancialServiceServer
func (*Server) GetAllBudgets(context.Context, *proto.GetAllBudgetsRequest) (*proto.GetAllBudgetsResponse, error) {
	panic("unimplemented")
}
