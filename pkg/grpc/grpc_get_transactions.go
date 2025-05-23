package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetTransactions Gets transactions for a given user in a paginated manner.
func (s *Server) GetTransactions(ctx context.Context, req *proto.GetTransactionsRequest) (*proto.GetTransactionsResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-get-transactions")
		defer span.End()
	}

	transactions, nextPageNumber, err := s.clickhouseConn.GetTransactions(ctx, &req.UserId, int64(req.PageNumber), int64(req.PageSize))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetTransactionsResponse{
		Transactions:   transactions,
		NextPageNumber: uint64(nextPageNumber),
	}, nil
}

func (s *Server) GetTransactionsForBankAccount(ctx context.Context, req *proto.GetTransactionsForBankAccountRequest) (*proto.GetTransactionsForBankAccountResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-get-transactions-for-bank-account")
		defer span.End()
	}

	transactions, nextPageNumber, err := s.clickhouseConn.GetTransactionsForAccount(
		ctx, &req.UserId,
		int64(req.PageNumber),
		int64(req.PageSize),
		req.PlaidAccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetTransactionsForBankAccountResponse{
		Transactions:   transactions,
		NextPageNumber: uint64(nextPageNumber),
	}, nil
}
