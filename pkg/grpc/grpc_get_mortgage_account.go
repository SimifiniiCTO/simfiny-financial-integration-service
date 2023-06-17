package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetMortgageAccount implements apiv1.FinancialServiceServer
func (s *Server) GetMortgageAccount(ctx context.Context, req *proto.GetMortgageAccountRequest) (*proto.GetMortgageAccountResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-mortgage-account")
		defer span.End()
	}

	// get the required credit account
	if _, err := s.conn.GetUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	account, err := s.conn.GetMortgageAccount(ctx, req.MortgageAccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetMortgageAccountResponse{
		MortageAccount: account,
	}, nil
}
