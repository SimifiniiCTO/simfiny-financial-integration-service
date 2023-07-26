package grpc

import (
	"context"

	clickhousedatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMonthlyTransactionCount(ctx context.Context, req *proto.GetMonthlyTransactionCountRequest) (*proto.GetMonthlyTransactionCountResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-monthly-transaction-count")
		defer span.End()
	}

	res, nextPageNumber, err := s.
		clickhouseConn.
		GetMonthlyTransactionCount(
			ctx,
			&req.UserId,
			&clickhousedatabase.BaseParams{
				Month:      req.Month,
				PageSize:   uint64(req.PageSize),
				PageNumber: uint64(req.PageNumber),
			})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetMonthlyTransactionCountResponse{
		MonthlyTransactionCounts: res,
		NextPageNumber:           nextPageNumber,
	}, nil
}
