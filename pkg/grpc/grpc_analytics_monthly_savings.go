package grpc

import (
	"context"

	clickhousedatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMonthlySavings(ctx context.Context, req *proto.GetMonthlySavingsRequest) (*proto.GetMonthlySavingsResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-monthly-savings")
		defer span.End()
	}

	res, nextPageNumber, err := s.
		clickhouseConn.
		GetMonthlySavings(
			ctx,
			&req.UserId,
			&clickhousedatabase.BaseParams{
				Month:        req.Month,
				MerchantName: "",
				PageSize:     uint64(req.PageSize),
				PageNumber:   uint64(req.PageNumber),
			})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetMonthlySavingsResponse{
		MonthlySavings: res,
		NextPageNumber: nextPageNumber,
	}, nil
}
