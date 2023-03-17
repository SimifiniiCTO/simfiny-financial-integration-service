package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// CreateBankAccount implements apiv1.FinancialServiceServer
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-create-bank-acct")
		defer span.End()
	}

	// attempt to create the bank account record but we must first set the account type to manual
	// because only manual accounts can be created via the API
	req.BankAccount.Type = proto.BankAccountType_BANK_ACCOUNT_TYPE_MANUAL

	// if a set of pockets are not already present for the given account, we must create them
	// pockets are crucial given they fascilitate the abstraction of sub-bank accounts better enabling proper goal management
	if len(req.BankAccount.Pockets) == 0 {
		req.BankAccount.Pockets = append(req.BankAccount.Pockets, s.defaultPockets()...)
	}

	createdBankAcct, err := s.conn.CreateBankAccount(ctx, req.UserId, req.BankAccount)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateBankAccountResponse{
		BankAccountId: createdBankAcct.Id,
	}, nil
}

func (s *Server) defaultPockets() []*proto.Pocket {
	return []*proto.Pocket{
		{
			Type: proto.PocketType_POCKET_TYPE_DEBT_REDUCTION,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_EMERGENCY_FUND,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_DISCRETIONARY_SPENDING,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_FUN_MONEY,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_LONG_TERM_SAVINGS,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_SHORT_TERM_SAVINGS,
		},
		{
			Type: proto.PocketType_POCKET_TYPE_INVESTMENT,
		},
	}
}
