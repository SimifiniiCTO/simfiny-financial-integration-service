package grpc

import (
	"context"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

// populateBalanceObjWithMortgageLoanAcct populates a balance slice with balance objects tied to mortgage loan accounts
func (s *Server) populateBalanceObjWithMortgageLoanAcct(account *proto.AccountsResponseMetadata, balanceObjects []*proto.AccountBalance) []*proto.AccountBalance {
	// iterate over all mortgage accts
	for _, mortgageAcct := range account.GetMortgage() {
		studentLoanBalance := &proto.AccountBalance{
			PlaidAccountID: mortgageAcct.GetPlaidAccountID(),
			AccountName:    mortgageAcct.GetAccountName(),
			Balance:        mortgageAcct.GetBalanceID(),
		}

		balanceObjects = append(balanceObjects, studentLoanBalance)
	}
	return balanceObjects
}

// populateBalanceObjWithStudentLoanAcct populates a balance slice with balance objects tied to student loan accounts
func (s *Server) populateBalanceObjWithStudentLoanAcct(account *proto.AccountsResponseMetadata,
	balanceObjects []*proto.AccountBalance) []*proto.AccountBalance {
	for _, studentLoanAcct := range account.GetStudentLoan() {
		studentLoanBalance := &proto.AccountBalance{
			PlaidAccountID: studentLoanAcct.GetPlaidAccountID(),
			AccountName:    studentLoanAcct.GetAccountName(),
			Balance:        studentLoanAcct.GetBalanceID(),
		}

		balanceObjects = append(balanceObjects, studentLoanBalance)
	}
	return balanceObjects
}

// populateBalanceObjWithInvestmentAcct populates a balance slice with balance objects tied to investment accounts
func (s *Server) populateBalanceObjWithInvestmentAcct(account *proto.AccountsResponseMetadata, balanceObjects []*proto.AccountBalance) []*proto.AccountBalance {
	for _, investmentAcct := range account.GetInvestments() {
		investmentBalance := &proto.AccountBalance{
			PlaidAccountID: investmentAcct.GetPlaidAccountID(),
			AccountName:    investmentAcct.GetAccountName(),
			Balance:        investmentAcct.GetBalanceID(),
		}

		balanceObjects = append(balanceObjects, investmentBalance)
	}
	return balanceObjects
}

// populateBalanceObjWithCreditAcc populates a balance slice with balance objects tied to a credit account
func (s *Server) populateBalanceObjWithCreditAcc(account *proto.AccountsResponseMetadata, balanceObjects []*proto.AccountBalance) []*proto.AccountBalance {
	// iterate over all credit acct
	for _, creditAcct := range account.GetCredit() {
		creditBalance := &proto.AccountBalance{
			PlaidAccountID: creditAcct.GetPlaidAccountID(),
			AccountName:    creditAcct.GetAccountName(),
			Balance:        creditAcct.GetBalanceID(),
		}

		balanceObjects = append(balanceObjects, creditBalance)
	}
	return balanceObjects
}

// populateBalanceObjWithDepositAcct populates a balance slice with balance objects tied to a deposit account
func (s *Server) populateBalanceObjWithDepositAcct(account *proto.AccountsResponseMetadata, balanceObjects []*proto.AccountBalance) []*proto.AccountBalance {
	for _, depositAcct := range account.GetDeposit() {
		depositBalance := &proto.AccountBalance{
			PlaidAccountID: depositAcct.GetPlaidAccountID(),
			AccountName:    depositAcct.GetAccountName(),
			Balance:        depositAcct.GetBalanceID(),
		}

		balanceObjects = append(balanceObjects, depositBalance)
	}
	return balanceObjects
}

// getAccountBalances gets account balances across all account types
func (s *Server) getAccountBalances(balanceObjects []*proto.AccountBalance, account *proto.AccountsResponseMetadata) []*proto.AccountBalance {
	balanceObjects = s.populateBalanceObjWithDepositAcct(account, balanceObjects)
	balanceObjects = s.populateBalanceObjWithCreditAcc(account, balanceObjects)
	balanceObjects = s.populateBalanceObjWithInvestmentAcct(account, balanceObjects)
	balanceObjects = s.populateBalanceObjWithStudentLoanAcct(account, balanceObjects)
	balanceObjects = s.populateBalanceObjWithMortgageLoanAcct(account, balanceObjects)
	return balanceObjects
}

func (s *Server) getVirtualAcct(ctx context.Context, vAcctID uint64) (*proto.AccountsResponseMetadata, error) {
	vAcct, err := s.getVirtualAcctFromDatastore(ctx, vAcctID)
	if err != nil {
		return nil, err
	}

	account := &proto.AccountsResponseMetadata{
		Deposit:     vAcct.DepositAccountID,
		Credit:      vAcct.CreditAccountID,
		Mortgage:    vAcct.MortgageLoanAccountID,
		StudentLoan: vAcct.StudentLoanAccountID,
		Investments: vAcct.InvestmentAccountID,
	}
	return account, nil
}

func (s *Server) getVirtualAcctFromDatastore(ctx context.Context, vAcctID uint64) (*proto.VirtualAccount, error) {
	return s.conn.FindVirtualAcctID(ctx, vAcctID)
}
