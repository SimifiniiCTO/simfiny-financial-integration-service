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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-delete-profile")
		defer span.End()
	}

	// perform deletion operation
	// TODO: given there will be a lot of data to delete, we should probably
	// do this in a background job. We should also perform this entire transaction as a
	// distributed transaction that is atomic and consistent.
	// ref: https://github.com/hibiken/asynq (redis based task queue)
	// ref: https://github.com/temporalio/temporal-ecommerce (temporal deletion workflow)
	if err := s.conn.DeleteUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteBankAccountResponse{
		Deleted: true,
	}, nil
}
