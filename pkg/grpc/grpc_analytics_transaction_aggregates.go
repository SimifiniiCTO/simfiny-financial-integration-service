package grpc

import (
	"context"

	clickhousedatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTransactionAggregates(ctx context.Context, req *proto.GetTransactionAggregatesRequest) (*proto.GetTransactionAggregatesResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-transaction-aggregates")
		defer span.End()
	}

	// perform operation
	res, nextPageNumber, err := s.
		clickhouseConn.
		GetTransactionAggregates(
			ctx,
			&req.UserId,
			&clickhousedatabase.TransactionAggregatedParams{
				Month:                   req.Month,
				PersonalFinanceCategory: req.PersonalFinanceCategoryPrimary,
				LocationCity:            req.LocationCity,
				PaymentChannel:          req.PaymentChannel,
				MerchantName:            req.MerchantName,
				PageSize:                uint64(req.PageSize),
				PageNumber:              uint64(req.PageNumber),
			})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetTransactionAggregatesResponse{
		TransactionAggregates: res,
		NextPageNumber:        nextPageNumber,
	}, nil
}
