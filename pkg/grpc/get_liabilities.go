package grpc

import (
	"context"
	"errors"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

func (s *Server) GetLiabilities(ctx context.Context, request *proto.GetLiabilitiesRequest) (*proto.GetLiabilitiesResponse, error) {
	txn := s.InstrumentIncomingRequestAndStartTxn(ctx, "get liabilities")
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

	account, err := s.getVirtualAcct(ctx, request.GetVirtualAccountID())
	if err != nil {
		return nil, err
	}

	return &proto.GetLiabilitiesResponse{Liabilities: &proto.Liabilities{
		Credit:      account.Credit,
		Mortgage:    account.Mortgage,
		StudentLoan: account.StudentLoan,
	}}, nil
}
