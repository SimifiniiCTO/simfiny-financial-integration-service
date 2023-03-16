package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// GetBudget implements apiv1.FinancialServiceServer
func (*Server) GetBudget(context.Context, *proto.GetBudgetRequest) (*proto.GetBudgetResponse, error) {
	panic("unimplemented")
}
