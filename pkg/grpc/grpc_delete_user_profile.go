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

	// get the user profie
	profile, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// delete the customer from stripe
	if _, err := s.stripeClient.Customers.Del(profile.StripeCustomerId, nil); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// perform deletion operation as a DTX (distributed transaction)
	//
	// TODO: given there will be a lot of data to delete, we should probably
	// do this in a background job. We should also perform this entire transaction as a
	// distributed transaction that is atomic and consistent.
	//
	// DTX:
	// 1. delete the user profile record and all account associations
	// 2. delete the user profile from the context of stripe (background job)
	// 3. delete the user profile from the context of plaid (background job)
	// 4. delete the tx records (background job)
	//
	// ref: https://github.com/hibiken/asynq (redis based task queue)
	// ref: https://github.com/temporalio/temporal-ecommerce (temporal deletion workflow)
	if err := s.conn.DeleteUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteUserProfileResponse{
		ProfileDeleted: true,
	}, nil
}
