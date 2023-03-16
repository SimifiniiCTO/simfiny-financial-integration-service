package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// UpdateBudget implements apiv1.FinancialServiceServer
func (*Server) UpdateBudget(context.Context, *proto.UpdateBudgetRequest) (*proto.UpdateBudgetResponse, error) {
	panic("unimplemented")
}
