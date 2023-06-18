package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// PlaidInitiateTokenExchange implements apiv1.FinancialServiceServer
func (s *Server) PlaidInitiateTokenExchange(ctx context.Context, req *proto.PlaidInitiateTokenExchangeRequest) (*proto.PlaidInitiateTokenExchangeResponse, error) {
	if req == nil {
		return nil, errors.New("invalid argument. request object cannot be nil")
	}

	if req.GetUserId() == 0 {
		return nil, errors.New("invalid request object. user id cannot be 0")
	}

	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-plaid-token-exchange")
		defer span.End()
	}

	// ensure user exists
	userProfile, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// If there is a configured limit on Plaid links then enforce that limit.
	if s.config.MaxPlaidLinks < len(userProfile.Link) {
		return nil, status.Error(codes.ResourceExhausted, "maximum number of plaid links reached")
	}

	// If billing is enabled and the current account is trialing, then limit them to a single Plaid link until their
	// trial has expired.
	if s.config.BillingEnabled {
		if userProfile.StripeSubscriptions != nil {
			if userProfile.StripeSubscriptions.StripeSubscriptionStatus == proto.StripeSubscriptionStatus_STRIPE_SUBSCRIPTION_STATUS_TRIALING {
				if len(userProfile.Link) > 0 {
					return nil, status.Error(codes.ResourceExhausted, "cannot add more Plaid links during trial")
				}
			}
		}
	}

	// create a link token via plaid
	token, err := s.plaidClient.CreateLinkToken(ctx, &plaidhandler.LinkTokenOptions{
		ClientUserID:             fmt.Sprintf("%d", req.UserId),
		LegalName:                req.FullName,
		PhoneNumber:              &req.PhoneNumber,
		PhoneNumberVerifiedTime:  pointer.TimeP(time.Now()),
		EmailAddress:             req.Email,
		EmailAddressVerifiedTime: pointer.TimeP(time.Now()),
		UpdateMode:               false,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.PlaidInitiateTokenExchangeResponse{
		LinkToken:  token.Token(),
		Expiration: token.Expiration().String(),
	}, nil

}
