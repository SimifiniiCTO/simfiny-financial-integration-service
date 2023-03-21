package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
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
	if _, err := s.conn.GetUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// create a link token
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
