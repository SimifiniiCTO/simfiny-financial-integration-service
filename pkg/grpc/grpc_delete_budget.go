package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// DeleteBudget implements apiv1.FinancialServiceServer
func (s *Server) DeleteBudget(ctx context.Context, req *proto.DeleteBudgetRequest) (*proto.DeleteBudgetResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-delete-budget")
		defer span.End()
	}

	// delete the required budget
	if err := s.conn.DeleteBudget(ctx, req.BudgetId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteBudgetResponse{
		Deleted: true,
	}, nil
}
