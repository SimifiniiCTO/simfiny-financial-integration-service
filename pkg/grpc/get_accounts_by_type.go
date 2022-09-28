package grpc

import (
	"context"
	"errors"

	proto "github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
)

func (s *Server) GetAccountsByType(ctx context.Context, request *proto.GetAccountsByTypeRequest) (*proto.GetAccountsByTypeResponse, error) {
	txn := s.InstrumentIncomingRequestAndStartTxn(ctx, "get accounts by type")
	defer txn.End()

	// TODO: time operation and trace distributed tx
	if request == nil {
		return nil, errors.New("invalid argument. request object cannot be nil")
	}

	if request.GetVirtualAccountID() == 0 {
		return nil, errors.New("invalid request object. virtual account id field cannot be 0")
	}

	if request.GetUserID() == 0 {
		return nil, errors.New("invalid request object. user id cannot be 0")
	}

	var account = &proto.AccountsResponseMetadata{}
	response, err := s.getVirtualAcct(ctx, request.GetVirtualAccountID())
	if err != nil {
		return nil, err
	}

	switch request.GetAccountType() {
	case proto.AccountType_INVESTMENT:
		account.Investments = response.Investments
	case proto.AccountType_CREDIT:
		account.Credit = response.Credit
	case proto.AccountType_DEPOSITORY:
		account.Deposit = response.Deposit
	case proto.AccountType_LOAN:
		account.StudentLoan = response.StudentLoan
		account.Mortgage = response.Mortgage
	default:
		account = response
	}

	return &proto.GetAccountsByTypeResponse{Accounts: account}, nil
}
