package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// DeleteSmartGoal implements apiv1.FinancialServiceServer
func (s *Server) DeleteSmartGoal(ctx context.Context, req *proto.DeleteSmartGoalRequest) (*proto.DeleteSmartGoalResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-delete-smart-goal")
		defer span.End()
	}

	// delete the required smart goal
	if err := s.conn.DeleteGoal(ctx, req.SmartGoalId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteSmartGoalResponse{
		Deleted: true,
	}, nil
}
