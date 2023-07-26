package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserAccountBalanceHistory(context.Context, *proto.GetUserAccountBalanceHistoryRequest) (*proto.GetUserAccountBalanceHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAccountBalanceHistory not implemented")
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
