package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// DeleteMilestone implements apiv1.FinancialServiceServer
func (*Server) DeleteMilestone(context.Context, *proto.DeleteMilestoneRequest) (*proto.DeleteMilestoneResponse, error) {
	panic("unimplemented")
}
