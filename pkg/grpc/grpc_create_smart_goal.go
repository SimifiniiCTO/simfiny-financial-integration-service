package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreateSmartGoal implements apiv1.FinancialServiceServer
func (s *Server) CreateSmartGoal(ctx context.Context, req *proto.CreateSmartGoalRequest) (*proto.CreateSmartGoalResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-create-smartgoal")
		defer span.End()
	}

	// create the required smart goal
	smartGoal, err := s.conn.CreateGoal(ctx, req.PocketId, req.SmartGoal)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateSmartGoalResponse{
		SmartGoalId: smartGoal.Id,
	}, nil
}
