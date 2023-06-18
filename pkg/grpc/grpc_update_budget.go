package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// UpdateBudget implements apiv1.FinancialServiceServer
func (s *Server) UpdateBudget(ctx context.Context, req *proto.UpdateBudgetRequest) (*proto.UpdateBudgetResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-update-budget")
		defer span.End()
	}

	if err := s.conn.UpdateBudget(ctx, req.Budget); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateBudgetResponse{
		Budget: req.Budget,
	}, nil
}
