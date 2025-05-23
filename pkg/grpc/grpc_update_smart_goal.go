package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// UpdateSmartGoal implements apiv1.FinancialServiceServer
func (s *Server) UpdateSmartGoal(ctx context.Context, req *proto.UpdateSmartGoalRequest) (*proto.UpdateSmartGoalResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-update-smart-goal")
		defer span.End()
	}

	if err := s.conn.UpdateGoal(ctx, req.SmartGoal); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateSmartGoalResponse{
		SmartGoalId: req.SmartGoal.Id,
	}, nil
}
