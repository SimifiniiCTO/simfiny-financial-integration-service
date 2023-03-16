package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetMilestonesBySmartGoalId implements apiv1.FinancialServiceServer
func (*Server) GetMilestonesBySmartGoalId(context.Context, *proto.GetMilestonesBySmartGoalIdRequest) (*proto.GetMilestonesBySmartGoalIdResponse, error) {
	panic("unimplemented")
}
