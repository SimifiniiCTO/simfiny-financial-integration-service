package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// GetAllBudgets implements apiv1.FinancialServiceServer
func (s *Server) GetAllBudgets(ctx context.Context, req *proto.GetAllBudgetsRequest) (*proto.GetAllBudgetsResponse, error) {
	// perform validations
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing request")
	}

	// validate the request
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-get-all-budgets")
		defer span.End()
	}

	// get all budgets
	pocket, err := s.conn.GetPocket(ctx, req.PocketId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get all the goals and milestones for a given pocket
	budgets := s.getPocketBudgets(pocket)
	return &proto.GetAllBudgetsResponse{
		Budgets: budgets,
	}, nil
}

func (s *Server) getPocketBudgets(pocket *proto.Pocket) []*proto.Budget {
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(context.Background())
		span := s.instrumentation.StartSegment(txn, "get-pocket-budgets")
		defer span.End()
	}

	goals := pocket.Goals
	milestones := make([]*proto.Milestone, 0)
	budgets := make([]*proto.Budget, 0)
	for _, goal := range goals {
		milestones = append(milestones, goal.Milestones...)
	}

	for _, milestone := range milestones {
		budgets = append(budgets, milestone.Budget)
	}
	return budgets
}
