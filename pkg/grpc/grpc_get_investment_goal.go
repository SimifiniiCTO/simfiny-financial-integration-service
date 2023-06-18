package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetInvestmentAcccount implements apiv1.FinancialServiceServer
func (s *Server) GetInvestmentAcccount(ctx context.Context, req *proto.GetInvestmentAcccountRequest) (*proto.GetInvestmentAcccountResponse, error) {
	// perform validations
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing request")
	}

	// validate the request
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "grpc-get-investment-account")
		defer span.End()
	}

	// get the required investment account
	profile, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: refactor this and only query investment account table
	for _, link := range profile.GetLink() {
		for _, account := range link.GetInvestmentAccounts() {
			if account.Id == req.InvestmentAccountId {
				return &proto.GetInvestmentAcccountResponse{
					InvestmentAccount: account,
				}, nil
			}
		}
	}

	return nil, status.Error(codes.NotFound, "investment account not found")
}
