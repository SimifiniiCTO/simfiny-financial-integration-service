package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetHistoricalAccountBalance(ctx context.Context, req *proto.GetHistoricalAccountBalanceRequest) (*proto.GetHistoricalAccountBalanceResponse, error) {
	// perform validations
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing request")
	}

	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-get-historical-account-balance")
		defer span.End()
	}

	// get the required budget
	accountBalance, err := s.clickhouseConn.GetBalanceHistoryByAccountID(ctx, &req.PlaidAccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetHistoricalAccountBalanceResponse{
		HistoricalAccountBalance: accountBalance,
	}, nil
}
