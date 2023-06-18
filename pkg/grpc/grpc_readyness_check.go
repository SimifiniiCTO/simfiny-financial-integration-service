package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// ReadynessCheck implements apiv1.FinancialServiceServer
func (s *Server) ReadynessCheck(ctx context.Context, req *proto.ReadynessCheckRequest) (*proto.ReadynessCheckResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-readiness-check")
		defer span.End()
	}

	// return a response
	return &proto.ReadynessCheckResponse{
		Healthy: true,
	}, nil
}
