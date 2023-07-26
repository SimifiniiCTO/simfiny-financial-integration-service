package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMonthlyTotalQuantityBySecurityAndUser(ctx context.Context, req *proto.GetMonthlyTotalQuantityBySecurityAndUserRequest) (*proto.GetMonthlyTotalQuantityBySecurityAndUserResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-monthly-income")
		defer span.End()
	}

	return nil, nil
}

func (s *Server) GetTotalInvestmentBySecurity(ctx context.Context, req *proto.GetTotalInvestmentBySecurityRequest) (*proto.GetTotalInvestmentBySecurityResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-total-investment-by-security")
		defer span.End()
	}

	return nil, nil
}

func (s *Server) GetUserAccountBalanceHistory(ctx context.Context, req *proto.GetUserAccountBalanceHistoryRequest) (*proto.GetUserAccountBalanceHistoryResponse, error) {

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
		span := s.instrumentation.StartSegment(txn, "grpc-get-account-balance-history")
		defer span.End()
	}

	return nil, status.Errorf(codes.Unimplemented, "method GetUserAccountBalanceHistory not implemented")
}

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

	return nil, status.Errorf(codes.Unimplemented, "method GetMelodyFinancialContext not implemented")
}
