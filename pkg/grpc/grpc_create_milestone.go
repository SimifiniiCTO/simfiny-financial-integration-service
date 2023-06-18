package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// CreateMilestone implements apiv1.FinancialServiceServer
func (s *Server) CreateMilestone(ctx context.Context, req *proto.CreateMilestoneRequest) (*proto.CreateMilestoneResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-create-milestone")
		defer span.End()
	}

	// create the required milestone
	milestone, err := s.conn.CreateMilestone(ctx, req.SmartGoalId, req.Milestone)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateMilestoneResponse{
		MilestoneId: milestone.Id,
	}, nil
}
