package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// DeleteBankAccount implements apiv1.FinancialServiceServer
func (s *Server) DeleteBankAccount(ctx context.Context, req *proto.DeleteBankAccountRequest) (*proto.DeleteBankAccountResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-delete-profile")
		defer span.End()
	}

	// ensure the user exists
	if _, err := s.conn.GetUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// ensure the bank account exists
	bankAcct, err := s.conn.GetBankAccount(ctx, req.BankAccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if bankAcct.Type == proto.BankAccountType_BANK_ACCOUNT_TYPE_MANUAL {
		// Given people will have a ton of data in their accounts, we
		// do a soft delete here, and then have a background job that will clean up all resources
		//
		// Deletion should only be allowed in the main branch of execution if the bank account being deleted
		// is a manual one
		if err := s.conn.DeleteBankAccount(ctx, req.BankAccountId); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		// TODO: perfrom async soft deletion as a background task
		s.logger.Warn("bank account is not a manual one, performing deletion in the background")
	}

	return &proto.DeleteBankAccountResponse{
		Deleted: true,
	}, nil
}
