package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// GetMilestonesBySmartGoalId implements apiv1.FinancialServiceServer
func (s *Server) GetMilestonesBySmartGoalId(ctx context.Context, req *proto.GetMilestonesBySmartGoalIdRequest) (*proto.GetMilestonesBySmartGoalIdResponse, error) {
	// perform validations
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing request")
	}

	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-get-milestones-by-smart-goal-id")
		defer span.End()
	}

	// get the required milestones
	milestones, err := s.conn.GetMilestonesByGoalID(ctx, req.SmartGoalId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetMilestonesBySmartGoalIdResponse{
		Milestones: milestones,
	}, nil
}
