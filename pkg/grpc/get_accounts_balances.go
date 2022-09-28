package grpc

import (
	"context"
	"errors"

	proto "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
)

func (s *Server) GetAccountsBalances(ctx context.Context, request *proto.GetAccountsBalanceRequest) (*proto.GetAccountsBalanceResponse, error) {
	txn := s.InstrumentIncomingRequestAndStartTxn(ctx, "get accounts balances")
	defer txn.End()

	if request == nil {
		return nil, errors.New("invalid argument. request object cannot be nil")
	}

	if request.GetVirtualAccountID() == 0 {
		return nil, errors.New("invalid request object. virtual account id field cannot be 0")
	}

	if request.GetUserID() == 0 {
		return nil, errors.New("invalid request object. user id cannot be 0")
	}

	account, err := s.getVirtualAcct(ctx, request.GetVirtualAccountID())
	if err != nil {
		return nil, err
	}

	balanceObjects := make([]*proto.AccountBalance, 0)
	balanceObjects = s.getAccountBalances(balanceObjects, account)
	return &proto.GetAccountsBalanceResponse{AccountBalances: balanceObjects}, nil
}
