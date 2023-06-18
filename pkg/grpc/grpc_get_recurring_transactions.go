package grpc

import (
	"context"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetReOcurringTransactions(ctx context.Context, req *proto.GetReCurringTransactionsRequest) (*proto.GetReCurringTransactionsResponse, error) {
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

	// instrument operation
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(ctx)
		span := s.instrumentation.StartDatastoreSegment(txn, "grpc-get-recurring-transactions")
		defer span.End()
	}

	// query db for user profile
	reCurringTransactions, err := s.clickhouseConn.GetUserReOccurringTransactions(ctx, &req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	txnIdSet := make([]string, 0)
	txnIdToRecurringTxnIdMap := make(map[string]uint64, 0)
	for _, txn := range reCurringTransactions {
		for _, txid := range helper.CommaSeparatedStringToStringSlice(txn.TransactionIds) {
			txnIdSet = append(txnIdSet, txid)
			txnIdToRecurringTxnIdMap[txid] = txn.Id
		}
	}

	// get all transactions
	transactions, err := s.clickhouseConn.GetTransactionsByPlaidTransactionIds(ctx, txnIdSet)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	participantRecurringTxnSet := make([]*proto.GetReCurringTransactionsResponse_ParticipantReCurringTransactions, 0)
	for _, txn := range transactions {
		if recurringTxnId, ok := txnIdToRecurringTxnIdMap[txn.TransactionId]; ok {
			participantRecurringTxnSet = append(participantRecurringTxnSet, &proto.GetReCurringTransactionsResponse_ParticipantReCurringTransactions{
				ReocurringTransactionId: recurringTxnId,
				Transactions:            transactions,
			})
		}
	}

	return &proto.GetReCurringTransactionsResponse{
		ReCcuringTransactions:            reCurringTransactions,
		ParticipantReCcuringTransactions: participantRecurringTxnSet,
	}, nil
}
