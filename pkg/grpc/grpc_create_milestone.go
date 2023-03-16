package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreateMilestone implements apiv1.FinancialServiceServer
func (*Server) CreateMilestone(context.Context, *proto.CreateMilestoneRequest) (*proto.CreateMilestoneResponse, error) {
	panic("unimplemented")
}
