package grpc

import (
	"context"

	clickhousedatabase "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/clickhouse-database"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserAccountBalanceHistory(context.Context, *proto.GetUserAccountBalanceHistoryRequest) (*proto.GetUserAccountBalanceHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAccountBalanceHistory not implemented")
}

func (s *Server) GetDebtToIncomeRatio(ctx context.Context, req *proto.GetDebtToIncomeRatioRequest) (*proto.GetDebtToIncomeRatioResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-debt-to-income-ratio")
		defer span.End()
	}

	// perform operation
	res, nextPageNumber, err := s.
		clickhouseConn.
		GetDebtToIncomeRatio(
			ctx,
			&req.UserId,
			&clickhousedatabase.DebtToIncomeParams{
				Month:      req.Month,
				PageSize:   uint64(req.PageSize),
				PageNumber: uint64(req.PageNumber),
			})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetDebtToIncomeRatioResponse{
		DebtToIncomeRatios: res,
		NextPageNumber:     nextPageNumber,
	}, nil
}

func (s *Server) GetExpenseMetrics(ctx context.Context, req *proto.GetExpenseMetricsRequest) (*proto.GetExpenseMetricsResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-expense-metrics")
		defer span.End()
	}

	// perform operation
	res, nextPageNumber, err := s.
		clickhouseConn.
		GetExpenseMetrics(
			ctx,
			&req.UserId,
			&clickhousedatabase.ExpenseMetricsParams{
				Month:                          req.Month,
				PersonalFinanceCategoryPrimary: req.PersonalFinanceCategoryPrimary,
				PageSize:                       uint64(req.PageSize),
				PageNumber:                     uint64(req.PageNumber),
			})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetExpenseMetricsResponse{
		ExpenseMetrics: res,
		NextPageNumber: nextPageNumber,
	}, nil
}

func (s *Server) GetFinancialProfile(ctx context.Context, req *proto.GetFinancialProfileRequest) (*proto.GetFinancialProfileResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-financial-profile")
		defer span.End()
	}

	// perform operation
	res, nextPageNumber, err := s.
		clickhouseConn.
		GetFinancialProfile(
			ctx,
			&req.UserId,
			&clickhousedatabase.FinancialProfileParams{
				Month:      req.Month,
				PageSize:   uint64(req.PageSize),
				PageNumber: uint64(req.PageNumber),
			})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetFinancialProfileResponse{
		FinancialProfiles: res,
		NextPageNumber:    nextPageNumber,
	}, nil
}

func (s *Server) GetIncomeExpenseRatio(context.Context, *proto.GetIncomeExpenseRatioRequest) (*proto.GetIncomeExpenseRatioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncomeExpenseRatio not implemented")
}

func (s *Server) GetIncomeMetrics(context.Context, *proto.GetIncomeMetricsRequest) (*proto.GetIncomeMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncomeMetrics not implemented")
}

func (s *Server) GetMerchantMonthlyExpenditure(context.Context, *proto.GetMerchantMonthlyExpenditureRequest) (*proto.GetMerchantMonthlyExpenditureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantMonthlyExpenditure not implemented")
}

func (s *Server) GetMonthlyBalance(context.Context, *proto.GetMonthlyBalanceRequest) (*proto.GetMonthlyBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyBalance not implemented")
}
func (s *Server) GetMonthlyExpenditure(context.Context, *proto.GetMonthlyExpenditureRequest) (*proto.GetMonthlyExpenditureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyExpenditure not implemented")
}
func (s *Server) GetMonthlyIncome(context.Context, *proto.GetMonthlyIncomeRequest) (*proto.GetMonthlyIncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyIncome not implemented")
}
func (s *Server) GetMonthlySavings(context.Context, *proto.GetMonthlySavingsRequest) (*proto.GetMonthlySavingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlySavings not implemented")
}
func (s *Server) GetMonthlyTotalQuantityBySecurityAndUser(context.Context, *proto.GetMonthlyTotalQuantityBySecurityAndUserRequest) (*proto.GetMonthlyTotalQuantityBySecurityAndUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyTotalQuantityBySecurityAndUser not implemented")
}
func (s *Server) GetMonthlyTransactionCount(context.Context, *proto.GetMonthlyTransactionCountRequest) (*proto.GetMonthlyTransactionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyTransactionCount not implemented")
}
func (s *Server) GetPaymentChannelMonthlyExpenditure(context.Context, *proto.GetPaymentChannelMonthlyExpenditureRequest) (*proto.GetPaymentChannelMonthlyExpenditureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentChannelMonthlyExpenditure not implemented")
}
func (s *Server) GetTotalInvestmentBySecurity(context.Context, *proto.GetTotalInvestmentBySecurityRequest) (*proto.GetTotalInvestmentBySecurityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalInvestmentBySecurity not implemented")
}
func (s *Server) GetMelodyFinancialContext(context.Context, *proto.GetMelodyFinancialContextRequest) (*proto.GetMelodyFinancialContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMelodyFinancialContext not implemented")
}
func (s *Server) GetTransactionsForBankAccount(context.Context, *proto.GetTransactionsForBankAccountRequest) (*proto.GetTransactionsForBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsForBankAccount not implemented")
}
