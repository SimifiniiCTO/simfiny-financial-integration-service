package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// UpdateMilestone implements apiv1.FinancialServiceServer
func (*Server) UpdateMilestone(context.Context, *proto.UpdateMilestoneRequest) (*proto.UpdateMilestoneResponse, error) {
	panic("unimplemented")
}
