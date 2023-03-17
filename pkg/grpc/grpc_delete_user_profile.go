package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

// DeleteUserProfile implements apiv1.FinancialServiceServer
func (s *Server) DeleteUserProfile(ctx context.Context, req *proto.DeleteUserProfileRequest) (*proto.DeleteUserProfileResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-delete-profile")
		defer span.End()
	}

	// NOTE: we need to perform this asynchronously given we have to delete all user data or we can opt to
	// soft delete and utilize a worker job to delete all user data in the background
	// TODO: we need to seriously consider this
	// TODO:  perform this entire transaction as a distributed transaction that is atomic and consistent.
	// via stripe
	// ref: https://github.com/hibiken/asynq (redis based task queue)
	// ref: https://github.com/temporalio/temporal-ecommerce (temporal deletion workflow)
	if err := s.conn.DeleteUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteUserProfileResponse{
		ProfileDeleted: true,
	}, nil
}
