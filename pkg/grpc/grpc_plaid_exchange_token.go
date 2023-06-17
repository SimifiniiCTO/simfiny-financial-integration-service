package grpc

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	taskhandler "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/task-handler"
	encryptdecrypt "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/encrypt_decrypt"
)

type AccessTokenMeta struct {
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
		span := s.instrumentation.StartSegment(txn, "grpc-plaid-exchange-token")
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

	// ensure the itemID is not already in use by the user
	// (meaning the user does not have duplicate login events with the same account registered on our backend)
	if _, err := s.conn.LinkExistsForItem(ctx, req.UserId, token.ItemId); err == nil {
		return nil, status.Error(codes.AlreadyExists, "item already exists for user")
	}

	// store a reference to the link in the database
	if _, err := s.createAndStoreLink(ctx, req.UserId, &tokenExchCallbackMetadata{
		item_token:       token,
		item_id:          &token.ItemId,
		institution_id:   &req.InstitutionId,
		institution_name: &req.InstitutionName,
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.PlaidExchangeTokenResponse{
		Success: true,
	}, nil
}

type tokenExchCallbackMetadata struct {
	item_token       *plaidhandler.ItemToken
	item_id          *string
	institution_id   *string
	institution_name *string
}

func (t *tokenExchCallbackMetadata) validate() error {
	if t.item_token == nil || t.item_id == nil || t.institution_id == nil || t.institution_name == nil {
		return status.Error(codes.InvalidArgument, "missing itemToken or itemID or institution_id or institution_name")
	}

	return nil
}

// createAndStoreLink creates a new link and stores it in the database and pull connected accounts as well as trigger
// background operations to sync the account's transactions
func (s *Server) createAndStoreLink(ctx context.Context, userID uint64, meta *tokenExchCallbackMetadata) (*uint64, error) {
	var (
		handler = s.plaidClient
	)

	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartSegment(txn, "create-and-store-link")
		defer span.End()
	}

	if userID == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing userID")
	}

	if err := meta.validate(); err != nil {
		return nil, err
	}

	s.logger.Info("creating and storing link for user", zap.Any("user_id", userID))

	accessToken := meta.item_token.AccessToken
	products := s.plaidClient.EnabledProductsToString()
	webhookUrl := handler.GetWebhooksURL()

	s.logger.Info("about to cryptographically encrypt token", zap.Any("user_id", userID))

	// cryptographically hash the access token before storing it
	res, err := encryptdecrypt.EncryptAccessToken(ctx, accessToken, s.kms, s.logger)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token := &proto.Token{
		ItemId:      *meta.item_id,
		KeyId:       res.KeyID,
		AccessToken: res.AccessToken,
		Version:     res.Version,
	}

	plaidLink := &proto.PlaidLink{
		Products:        products,
		WebhookUrl:      webhookUrl,
		InstitutionId:   *meta.institution_id,
		InstitutionName: *meta.institution_name,
		UsePlaidSync:    true,
		ItemId:          *meta.item_id,
	}

	// create a reference to a new link and store it in the database
	link := &proto.Link{
		LinkStatus:                proto.LinkStatus_LINK_STATUS_PENDING, // successfully exchanged public token for access token
		PlaidLink:                 plaidLink,                            // plaid link object
		PlaidNewAccountsAvailable: false,                                // identifier wether new accounts are available
		InstitutionName:           *meta.institution_name,               // institution name
		LastSuccessfulUpdate:      time.Now().String(),                  // last successful update
		Token:                     token,                                // access token
		PlaidInstitutionId:        *meta.institution_id,                 // plaid institution id
		LinkType:                  proto.LinkType_LINK_TYPE_PLAID,       // link type
	}

	// retrieve links bank accounts
	link.BankAccounts, err = handler.GetAccounts(ctx, accessToken, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := s.conn.CreateLink(ctx, userID, link, false)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: kick off a background job to fetch the account transactions
	if plaidLink.UsePlaidSync {
		// kick off a background job to pull transactions by use of plaid sync
		if err := taskhandler.DispatchPlaidSyncTask(ctx, s.Taskprocessor, s.instrumentation, s.logger, userID, result.Id, accessToken); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		linkId := result.Id
		start := time.Now().Add(-30 * 24 * time.Hour) // Last 30 days.
		end := time.Now()
		// kick off background job to pull transactions
		if err := taskhandler.DispatchPullTransactionsTask(
			ctx,
			s.Taskprocessor,
			s.instrumentation,
			s.logger,
			userID,
			linkId,
			accessToken,
			start,
			end,
		); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &result.Id, nil
}
