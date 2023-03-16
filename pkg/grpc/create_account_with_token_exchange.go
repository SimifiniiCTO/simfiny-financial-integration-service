package grpc

import (
	"context"
	"errors"

	"github.com/plaid/plaid-go/plaid"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

func (s *Server) CreateAccountWithTokenExchange(ctx context.Context, request *proto.CreateAccountTokenExchangeRequest) (*proto.
	CreateAccountTokenExchangeResponse, error) {
	// TODO: time operation and trace distributed tx
	if request == nil {
		return nil, errors.New("invalid argument. request object cannot be nil")
	}

	if request.GetUserID() == 0 {
		return nil, errors.New("invalid request object. user id cannot be 0")
	}

	if request.GetPublicToken() == "" {
		return nil, errors.New("invalid request object. public token cannot be empty")
	}

	var vAcct *proto.VirtualAccount
	var err error
	var resp plaid.ItemPublicTokenExchangeResponse

	plaidTokenExchangeReq := s.plaidClient.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(*plaid.NewItemPublicTokenExchangeRequest(request.PublicToken))
	resp, _, err = plaidTokenExchangeReq.Execute()
	if err != nil {
		return nil, err
	}

	accessToken := resp.GetAccessToken()
	plaidItemID := resp.GetItemId()

	vAcct, err = s.plaidInternalHandler.CallPlaidAPIAndPopulateVirtualAccount(ctx, &accessToken)
	if err != nil {
		return nil, err
	}

	vAcct.PlaidItemID = plaidItemID
	vAcct.AccessToken = accessToken
	vAcct.UserID = request.GetUserID()
	/*
		account, err := s.conn.CreateVirtualAccount(ctx, vAcct, accessToken)
		if err != nil {
			return nil, err
		}
	*/

	return &proto.CreateAccountTokenExchangeResponse{VirtualAccountID: 10}, nil
}
