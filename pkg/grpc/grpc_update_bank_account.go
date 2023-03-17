package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// UpdateBankAccount implements apiv1.FinancialServiceServer
func (s *Server) UpdateBankAccount(ctx context.Context, req *proto.UpdateBankAccountRequest) (*proto.UpdateBankAccountResponse, error) {
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

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-update-bank-account")
		defer span.End()
	}

	if err := s.conn.UpdateBankAccount(ctx, req.BankAccount); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateBankAccountResponse{
		Updated:     true,
		BankAccount: req.BankAccount,
	}, nil
}
