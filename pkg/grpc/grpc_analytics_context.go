package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMelodyFinancialContext(ctx context.Context, req *proto.GetMelodyFinancialContextRequest) (*proto.GetMelodyFinancialContextResponse, error) {

	// perform validations
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing request")
	}

	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-get-financial-context")
		defer span.End()
	}

	res, err := s.
		clickhouseConn.
		GetFinancialContextForCurrentMonth(
			ctx,
			&req.UserId, 4)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetMelodyFinancialContextResponse{
		MelodyFinancialContext: res,
	}, nil
}
