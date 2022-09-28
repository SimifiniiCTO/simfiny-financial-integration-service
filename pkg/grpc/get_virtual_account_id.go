package grpc

import (
	"context"
	"errors"

	"github.com/SimifiniiCTO/simfinii/src/backend/services/financial-integration-service/proto"
)

func (s *Server) GetVirtualAccountId(ctx context.Context, request *proto.GetVirtualAccountIdRequest) (*proto.GetVirtualAccountIdResponse, error) {
	txn := s.InstrumentIncomingRequestAndStartTxn(ctx, "get virtual account id")
	defer txn.End()

	// TODO: time operation and trace distributed tx
	if request == nil {
		return nil, errors.New("invalid argument. request object cannot be nil")
	}

	if request.GetUserID() == 0 {
		return nil, errors.New("invalid request object. user id field cannot be 0")
	}

	acct, err := s.conn.FindVirtualAcctByUserID(ctx, request.GetUserID())
	if err != nil {
		return nil, err
	}

	return &proto.GetVirtualAccountIdResponse{VirtualAccountID: acct.GetId()}, nil
}
