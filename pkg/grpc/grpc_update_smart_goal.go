package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// UpdateSmartGoal implements apiv1.FinancialServiceServer
func (*Server) UpdateSmartGoal(context.Context, *proto.UpdateSmartGoalRequest) (*proto.UpdateSmartGoalResponse, error) {
	panic("unimplemented")
}
