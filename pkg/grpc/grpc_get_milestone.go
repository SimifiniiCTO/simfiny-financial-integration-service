package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetMilestone implements apiv1.FinancialServiceServer
func (s *Server) GetMilestone(ctx context.Context, req *proto.GetMilestoneRequest) (*proto.GetMilestoneResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-get-milestone")
		defer span.End()
	}

	// get the required milestone
	milestone, err := s.conn.GetMilestone(ctx, req.MilestoneId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetMilestoneResponse{
		Milestone: milestone,
	}, nil
}
