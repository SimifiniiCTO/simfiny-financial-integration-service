package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreateSmartGoal implements apiv1.FinancialServiceServer
func (*Server) CreateSmartGoal(context.Context, *proto.CreateSmartGoalRequest) (*proto.CreateSmartGoalResponse, error) {
	panic("unimplemented")
}
