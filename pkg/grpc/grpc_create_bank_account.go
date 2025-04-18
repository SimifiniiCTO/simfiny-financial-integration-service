package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// CreateBankAccount creates a new bank account for the given user
// The bank account must be of type manual as this is the only type that can be created via the API
func (s *Server) CreateBankAccount(ctx context.Context, req *proto.CreateBankAccountRequest) (*proto.CreateBankAccountResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-create-bank-acct")
		defer span.End()
	}

	// attempt to create the bank account record but we must first set the account type to manual
	// because only manual accounts can be created via the API
	req.BankAccount.Type = proto.BankAccountType_BANK_ACCOUNT_TYPE_MANUAL

	// if a set of pockets are not already present for the given account, we must create them
	// pockets are crucial given they fascilitate the abstraction of sub-bank accounts better enabling proper goal management
	if len(req.BankAccount.Pockets) == 0 {
		req.BankAccount.Pockets = helper.DefaultPockets()
	}

	createdBankAcct, err := s.conn.CreateBankAccount(ctx, req.UserId, req.BankAccount)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateBankAccountResponse{
		BankAccountId: createdBankAcct.Id,
	}, nil
}
