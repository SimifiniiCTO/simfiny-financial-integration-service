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
