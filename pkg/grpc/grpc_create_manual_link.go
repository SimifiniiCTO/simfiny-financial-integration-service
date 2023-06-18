package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

// CreateManualLink implements apiv1.FinancialServiceServer
// This endpoint is used to create a manual link for a user. A manual link is used to associate a manually created bank account.
// All links created via plaid should not be created via this endpoint.
func (s *Server) CreateManualLink(ctx context.Context, req *proto.CreateManualLinkRequest) (*proto.CreateManualLinkResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-create-manual-link")
		defer span.End()
	}

	// this endpoint is only used to create manual links
	// manual links are used to associate an abstraction of a manual bank account for the user
	manualLink := req.GetManualAccountLink()
	manualLink.LinkType = proto.LinkType_LINK_TYPE_MANUAL
	manualLink.LinkStatus = proto.LinkStatus_LINK_STATUS_SUCCESS

	// create a default bank account object for the user creating a manual link
	manualAcct, err := s.autoGenerateManualBankAccount(ctx, &req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// create a manual bank account for the user with a default set of associated pockets
	manualLink.BankAccounts = []*proto.BankAccount{manualAcct}

	// create the required manual link
	link, err := s.conn.CreateLink(ctx, req.UserId, req.ManualAccountLink, true)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateManualLinkResponse{
		LinkId: link.GetId(),
	}, nil
}

// autoGenerateManualBankAccount is used to generate a manual bank account for a user
func (s *Server) autoGenerateManualBankAccount(ctx context.Context, userID *uint64) (*proto.BankAccount, error) {
	if s.instrumentation != nil {
		txn := s.instrumentation.GetTraceFromContext(context.Background())
		span := s.instrumentation.StartSegment(txn, "auto-generate-manual-bank-account")
		defer span.End()
	}

	if userID == nil {
		return nil, errors.New("invalid user id")
	}

	bankAccountName := fmt.Sprintf("manual-account-%d-%s", *userID, uuid.New().String())

	return &proto.BankAccount{
		UserId:         *userID,
		Name:           bankAccountName,
		Number:         uuid.New().String(),
		Type:           proto.BankAccountType_BANK_ACCOUNT_TYPE_MANUAL,
		Balance:        0,
		Currency:       "USD",
		CurrentFunds:   0,
		BalanceLimit:   10000,
		Pockets:        s.DefaultPockets(),
		PlaidAccountId: "",
		Subtype:        "deposit",
	}, nil
}
