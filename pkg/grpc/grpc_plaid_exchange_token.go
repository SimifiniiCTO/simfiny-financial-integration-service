package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

type accessTokenMeta struct {
	accessToken string
	keyID       string
	version     string
}

// PlaidExchangeToken implements apiv1.FinancialServiceServer
func (s *Server) PlaidExchangeToken(ctx context.Context, req *proto.PlaidExchangeTokenRequest) (*proto.PlaidExchangeTokenResponse, error) {
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
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-plaid-exchange-token")
		defer span.End()
	}

	// ensure user exists
	if _, err := s.conn.GetUserProfileByUserID(ctx, req.UserId); err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// exchange public token for access token
	token, err := s.plaidClient.ExchangePublicToken(ctx, req.PublicToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// ensure the user profile making the request actually exists
	userProfile, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// cryptographically hash the access token before storing it
	// TODO: refactor this - may not nott to store decryption key
	meta, err := s.EncryptAccessToken(ctx, token.AccessToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// populate the user profile with the access token and necessary decryption keys
	userProfile.PlaidAccessToken = meta.accessToken
	userProfile.DecryptionAccessTokenKey = meta.keyID
	userProfile.DecryptionAccessTokenVersion = meta.version

	if err := s.conn.UpdateUserProfile(ctx, userProfile); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.PlaidExchangeTokenResponse{
		Success: true,
	}, nil
}
