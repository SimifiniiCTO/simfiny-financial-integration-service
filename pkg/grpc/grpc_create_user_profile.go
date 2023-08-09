package grpc

import (
	"context"

	"github.com/stripe/stripe-go/v74"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// CreateUserProfile implements apiv1.FinancialServiceServer
func (s *Server) CreateUserProfile(ctx context.Context, req *proto.CreateUserProfileRequest) (*proto.CreateUserProfileResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-create-profile")
		defer span.End()
	}

	// TODO: create the account as a distributed transaction with stripe
	customer, err := s.stripeClient.Customers.New(&stripe.CustomerParams{
		Email: &req.Email,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	req.Profile.StripeCustomerId = customer.ID

	// store email address
	if req.Profile.Email == "" && req.Email != "" {
		req.Profile.Email = req.Email
	}

	// save the user profile
	res, err := s.conn.CreateUserProfile(ctx, req.Profile)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateUserProfileResponse{
		UserId: res.UserId,
	}, nil
}
