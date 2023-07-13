package grpc

import (
	"context"
	"fmt"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	encryptdecrypt "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/encrypt_decrypt"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) PlaidInitiateTokenUpdate(ctx context.Context, req *proto.PlaidInitiateTokenUpdateRequest) (*proto.PlaidInitiateTokenUpdateResponse, error) {
	if req == nil {
		return nil, errors.New("invalid argument. request object cannot be nil")
	}

	if req.GetUserId() == 0 {
		return nil, errors.New("invalid request object. user id cannot be 0")
	}

	if req.GetLinkId() == 0 {
		return nil, errors.New("invalid request object. link id cannot be 0")
	}

	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// ensure operation finished in time
	ctx, cancel := context.WithTimeout(ctx, s.config.RpcTimeout)
	defer cancel()

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-plaid-token-update")
		defer span.End()
	}

	// ensure user exists
	link, err := s.conn.GetLink(ctx, req.UserId, req.LinkId, true)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if link == nil || link.Token == nil {
		return nil, status.Errorf(codes.Internal, "Link token credentials are required but not provided")
	}

	res, err := encryptdecrypt.EncryptAccessToken(ctx, link.Token.AccessToken, s.kms, s.logger)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	linkTokenResponse, err := s.plaidClient.UpdateLinkToken(ctx, &plaidhandler.LinkTokenOptions{
		ClientUserID: fmt.Sprintf("%d", req.UserId),
	}, &res.AccessToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get the link by id
	return &proto.PlaidInitiateTokenUpdateResponse{
		LinkToken:  linkTokenResponse.Token(),
		Expiration: linkTokenResponse.Expiration().String(),
	}, nil
}
