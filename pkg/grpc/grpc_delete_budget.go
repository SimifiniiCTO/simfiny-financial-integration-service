package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// DeleteBudget implements apiv1.FinancialServiceServer
func (*Server) DeleteBudget(context.Context, *proto.DeleteBudgetRequest) (*proto.DeleteBudgetResponse, error) {
	panic("unimplemented")
}
