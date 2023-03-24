package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetForecast implements apiv1.FinancialServiceServer
func (s *Server) GetForecast(ctx context.Context, req *proto.GetForecastRequest) (*proto.GetForecastResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-forecast")
		defer span.End()
	}

	// get the required forecast
	goal, err := s.conn.GetGoal(ctx, req.SmartGoalId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetForecastResponse{
		Forecast: goal.Forecasts,
	}, nil
}
