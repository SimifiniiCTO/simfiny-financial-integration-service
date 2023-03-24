package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// UpdateMilestone implements apiv1.FinancialServiceServer
func (s *Server) UpdateMilestone(ctx context.Context, req *proto.UpdateMilestoneRequest) (*proto.UpdateMilestoneResponse, error) {
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

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-update-milestone")
		defer span.End()
	}

	if err := s.conn.UpdateMilestone(ctx, req.Milestone); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateMilestoneResponse{
		Milestone: req.Milestone,
	}, nil
}
