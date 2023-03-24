package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreateBudget implements apiv1.FinancialServiceServer
func (s *Server) CreateBudget(ctx context.Context, req *proto.CreateBudgetRequest) (*proto.CreateBudgetResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-create-budget")
		defer span.End()
	}

	// create the required budget
	budget, err := s.conn.CreateBudget(ctx, req.MilestroneId, req.Budget)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateBudgetResponse{
		BudgetId: budget.Id,
	}, nil
}
