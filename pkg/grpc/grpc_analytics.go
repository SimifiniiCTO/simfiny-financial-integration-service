package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetDebtToIncomeRatio(context.Context, *proto.GetDebtToIncomeRatioRequest) (*proto.GetDebtToIncomeRatioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDebtToIncomeRatio not implemented")
}
func (s *Server) ListDebtToIncomeRatio(context.Context, *proto.ListDebtToIncomeRatioRequest) (*proto.ListDebtToIncomeRatioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDebtToIncomeRatio not implemented")
}
func (s *Server) GetExpenseMetrics(context.Context, *proto.GetExpenseMetricsRequest) (*proto.GetExpenseMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpenseMetrics not implemented")
}
func (s *Server) ListExpenseMetrics(context.Context, *proto.ListExpenseMetricsRequest) (*proto.ListExpenseMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExpenseMetrics not implemented")
}
func (s *Server) GetFinancialProfile(context.Context, *proto.GetFinancialProfileRequest) (*proto.GetFinancialProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialProfile not implemented")
}
func (s *Server) ListFinancialProfile(context.Context, *proto.ListFinancialProfileRequest) (*proto.ListFinancialProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFinancialProfile not implemented")
}
func (s *Server) GetIncomeExpenseRatio(context.Context, *proto.GetIncomeExpenseRatioRequest) (*proto.GetIncomeExpenseRatioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncomeExpenseRatio not implemented")
}
func (s *Server) ListIncomeExpenseRatio(context.Context, *proto.ListIncomeExpenseRatioRequest) (*proto.ListIncomeExpenseRatioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncomeExpenseRatio not implemented")
}
func (s *Server) GetIncomeMetrics(context.Context, *proto.GetIncomeMetricsRequest) (*proto.GetIncomeMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncomeMetrics not implemented")
}
func (s *Server) ListIncomeMetrics(context.Context, *proto.ListIncomeMetricsRequest) (*proto.ListIncomeMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIncomeMetrics not implemented")
}
func (s *Server) GetMerchantMonthlyExpenditure(context.Context, *proto.GetMerchantMonthlyExpenditureRequest) (*proto.GetMerchantMonthlyExpenditureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantMonthlyExpenditure not implemented")
}
func (s *Server) ListMerchantMonthlyExpenditure(context.Context, *proto.ListMerchantMonthlyExpenditureRequest) (*proto.ListMerchantMonthlyExpenditureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchantMonthlyExpenditure not implemented")
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
