package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
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

	// perform deletion as a distributed transaction
	wf, err := s.ExecuteWorkflow(ctx, s.TransactionManager.DeleteProfileWorkflow, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// execute the workflow and wait for completion
	if err := wf.Get(ctx, nil); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if err := s.conn.DeleteUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteUserProfileResponse{
		ProfileDeleted: true,
	}, nil
}
