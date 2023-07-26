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

func (s *Server) GetMonthlyBalance(context.Context, *proto.GetMonthlyBalanceRequest) (*proto.GetMonthlyBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyBalance not implemented")
}

func (s *Server) GetMonthlyExpenditure(ctx context.Context, req *proto.GetMonthlyExpenditureRequest) (*proto.GetMonthlyExpenditureResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-monthly-expenditure")
		defer span.End()
	}

	res, nextPageNumber, err := s.
		clickhouseConn.
		GetMonthlyExpenditure(
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

	return &proto.GetMonthlyExpenditureResponse{
		MonthlyExpenditures: res,
		NextPageNumber:      nextPageNumber,
	}, nil
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
