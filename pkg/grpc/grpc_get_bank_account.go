package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// GetBankAccount implements apiv1.FinancialServiceServer
func (s *Server) GetBankAccount(ctx context.Context, req *proto.GetBankAccountRequest) (*proto.GetBankAccountResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-bank-account")
		defer span.End()
	}

	// get the required bank account
	bankAccount, err := s.conn.GetBankAccount(ctx, req.BankAccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetBankAccountResponse{
		BankAccount: bankAccount,
	}, nil

}
