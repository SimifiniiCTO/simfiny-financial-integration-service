package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

// GetUserProfile implements apiv1.FinancialServiceServer
func (s *Server) GetUserProfile(ctx context.Context, req *proto.GetUserProfileRequest) (*proto.GetUserProfileResponse, error) {
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
		span := s.instrumentation.StartSegment(txn, "grpc-get-profile")
		defer span.End()
	}

	// query db for user profile
	record, err := s.conn.GetUserProfileByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// remove the access token from the response
	for _, link := range record.Link {
		NullifyAccessToken(link)
	}

	// query the financial context from the given user
	financialContext, err := s.clickhouseConn.
		GetFinancialContextForCurrentMonth(ctx, &req.UserId, 10)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// enrich the financial context with the accounts
	// from the user profile
	enrichFinancialContext(record, financialContext)

	return &proto.GetUserProfileResponse{
		Profile:          record,
		FinancialContext: financialContext,
	}, nil
}

func enrichFinancialContext(user *proto.UserProfile, ctx *proto.MelodyFinancialContext) {
	// extract all financial accounst accross all connected links
	// and add them to the response
	bankAccts := make([]*proto.BankAccount, 0)
	creditAccts := make([]*proto.CreditAccount, 0)
	loanAccts := make([]*proto.StudentLoanAccount, 0)
	investmentAccts := make([]*proto.InvestmentAccount, 0)
	mortgageAccts := make([]*proto.MortgageAccount, 0)

	for _, link := range user.Link {
		if link.BankAccounts != nil {
			bankAccts = append(bankAccts, link.BankAccounts...)
		}

		if link.CreditAccounts != nil {
			creditAccts = append(creditAccts, link.CreditAccounts...)
		}

		if link.StudentLoanAccounts != nil {
			loanAccts = append(loanAccts, link.StudentLoanAccounts...)
		}

		if link.InvestmentAccounts != nil {
			investmentAccts = append(investmentAccts, link.InvestmentAccounts...)
		}

		if link.MortgageAccounts != nil {
			mortgageAccts = append(mortgageAccts, link.MortgageAccounts...)
		}
	}

	// we only take 2 accounts of each type
	if len(bankAccts) > 2 {
		bankAccts = bankAccts[:2]
	}

	if len(creditAccts) > 2 {
		creditAccts = creditAccts[:2]
	}

	if len(loanAccts) > 2 {
		loanAccts = loanAccts[:2]
	}

	if len(investmentAccts) > 2 {
		investmentAccts = investmentAccts[:2]
	}

	if len(mortgageAccts) > 2 {
		mortgageAccts = mortgageAccts[:2]
	}

	// sanitize the bank accounts
	// we do not want the pockets
	for _, acct := range bankAccts {
		acct.Pockets = nil
	}

	// sanitize the investment account s
	// we do not want the positions nor holdings
	for _, acct := range investmentAccts {
		acct.Securities = nil
		acct.Holdings = nil
	}

	ctx.BankAccounts = bankAccts
	ctx.CreditAccounts = creditAccts
	ctx.StudentLoanAccounts = loanAccts
	ctx.InvestmentAccounts = investmentAccts
	ctx.MortgageLoanAccounts = mortgageAccts
}
