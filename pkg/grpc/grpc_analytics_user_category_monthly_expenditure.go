package grpc

import (
	"context"

	clickhousedatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserCategoryMonthlyExpenditure(ctx context.Context, req *proto.GetUserCategoryMonthlyExpenditureRequest) (*proto.GetUserCategoryMonthlyExpenditureResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-user-category-monthly-expenditure")
		defer span.End()
	}

	// perform operation
	res, nextPageNumber, err := s.
		clickhouseConn.
		GetCategoryMonthlyExpenditures(
			ctx,
			&req.UserId,
			&clickhousedatabase.CategoryMonthlyExpenditureParams{
				Month:                          req.Month,
				PersonalFinanceCategoryPrimary: req.PersonalFinanceCategoryPrimary,
				PageSize:                       uint64(req.PageSize),
				PageNumber:                     uint64(req.PageNumber),
			})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetUserCategoryMonthlyExpenditureResponse{
		CategoryMonthlyExpenditure: res,
		NextPageNumber:             nextPageNumber,
	}, nil
}
