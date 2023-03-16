package grpc

import (
	"context"
	"errors"
	"strconv"

	"github.com/plaid/plaid-go/plaid"
	"go.uber.org/zap"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

func (s *Server) InitiateAccountSetupAndGetLinkToken(ctx context.Context, request *proto.InitiateAccountSetupRequest) (*proto.InitiateAccountSetupResponse, error) {
	if request == nil {
		return nil, errors.New("invalid argument. request object cannot be nil")
	}

	if request.GetUserID() == 0 {
		return nil, errors.New("invalid request object. user id cannot be 0")
	}

	if s.config.Environment == "dev" {
		s.logger.Info("calling dev plaid account")
		resp, err := s.performnOperationAgainstPlaidSandboxEnv(ctx)
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}

		return &proto.InitiateAccountSetupResponse{
			LinkToken:      resp.GetPublicToken(),
			Expiration:     "",
			PlaidRequestID: resp.GetRequestId()}, nil
	}

	// TODO: make this a retryable operation and associate this with a retryable error
	resp, err := s.performnOperationAgainstPlaidProdEnv(ctx, request.GetUserID())
	if err != nil {
		s.logger.Error("failed to get linktoken", zap.Error(err))
		return nil, err
	}

	return &proto.InitiateAccountSetupResponse{
		LinkToken:      resp.GetLinkToken(),
		Expiration:     resp.GetExpiration().String(),
		PlaidRequestID: resp.GetRequestId()}, nil
}

func (s *Server) performnOperationAgainstPlaidSandboxEnv(ctx context.Context) (*plaid.SandboxPublicTokenCreateResponse, error) {
	sandboxInstitution := "ins_109508"
	testProducts := []plaid.Products{plaid.PRODUCTS_AUTH, plaid.PRODUCTS_INVESTMENTS, plaid.PRODUCTS_LIABILITIES}
	resp, _, err := s.plaidClient.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		*plaid.NewSandboxPublicTokenCreateRequest(
			sandboxInstitution,
			testProducts,
		),
	).Execute()

	if err != nil {
		s.logger.Error("error occured", zap.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (s *Server) performnOperationAgainstPlaidProdEnv(ctx context.Context, userID uint64) (plaid.LinkTokenCreateResponse, error) {
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: strconv.Itoa(int(userID)),
	}

	plaidRequest := plaid.NewLinkTokenCreateRequest(
		s.config.ServiceName,
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	plaidRequest.SetProducts([]plaid.Products{
		plaid.PRODUCTS_AUTH,
		plaid.PRODUCTS_LIABILITIES,
		plaid.PRODUCTS_INVESTMENTS,
		plaid.PRODUCTS_TRANSACTIONS,
		plaid.PRODUCTS_BALANCE,
	})
	plaidRequest.SetLinkCustomizationName(s.config.ServiceName)
	plaidRequest.SetWebhook(s.config.PlaidWebhookURI)
	plaidRequest.SetRedirectUri(s.config.PlaidRedirectURI)

	resp, _, err := s.
		plaidClient.
		PlaidApi.
		LinkTokenCreate(ctx).
		LinkTokenCreateRequest(*plaidRequest).
		Execute()
	if err != nil {
		s.logger.Error(err.Error())
		return plaid.LinkTokenCreateResponse{}, err
	}
	return resp, nil
}
