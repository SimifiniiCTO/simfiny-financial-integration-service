package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// GetSmartGoalsByPocketId implements apiv1.FinancialServiceServer
func (s *Server) GetSmartGoalsByPocketId(ctx context.Context, req *proto.GetSmartGoalsByPocketIdRequest) (*proto.GetSmartGoalsByPocketIdResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-get-smart-goals-by-pocket-id")
		defer span.End()
	}

	// get the required smart goals
	smartGoals, err := s.conn.GetGoalsByPocketID(ctx, req.PocketId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetSmartGoalsByPocketIdResponse{
		SmartGoals: smartGoals,
	}, nil

}
