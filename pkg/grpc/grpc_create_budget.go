package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

// CreateBudget implements apiv1.FinancialServiceServer
func (*Server) CreateBudget(context.Context, *proto.CreateBudgetRequest) (*proto.CreateBudgetResponse, error) {
	panic("unimplemented")
}
