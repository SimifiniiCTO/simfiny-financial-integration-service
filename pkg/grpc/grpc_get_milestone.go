package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetMilestone implements apiv1.FinancialServiceServer
func (*Server) GetMilestone(context.Context, *proto.GetMilestoneRequest) (*proto.GetMilestoneResponse, error) {
	panic("unimplemented")
}
