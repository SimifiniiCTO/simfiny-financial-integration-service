package grpc

import (
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

func TestServer_populateBalanceObjWithMortgageLoanAcct(t *testing.T) {
	type args struct {
		account        *proto.AccountsResponseMetadata
		balanceObjects []*proto.AccountBalance
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want []*proto.AccountBalance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.populateBalanceObjWithMortgageLoanAcct(tt.args.account, tt.args.balanceObjects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.populateBalanceObjWithMortgageLoanAcct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_populateBalanceObjWithStudentLoanAcct(t *testing.T) {
	type args struct {
		account        *proto.AccountsResponseMetadata
		balanceObjects []*proto.AccountBalance
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want []*proto.AccountBalance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.populateBalanceObjWithStudentLoanAcct(tt.args.account, tt.args.balanceObjects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.populateBalanceObjWithStudentLoanAcct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_populateBalanceObjWithInvestmentAcct(t *testing.T) {
	type args struct {
		account        *proto.AccountsResponseMetadata
		balanceObjects []*proto.AccountBalance
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want []*proto.AccountBalance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.populateBalanceObjWithInvestmentAcct(tt.args.account, tt.args.balanceObjects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.populateBalanceObjWithInvestmentAcct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_populateBalanceObjWithCreditAcc(t *testing.T) {
	type args struct {
		account        *proto.AccountsResponseMetadata
		balanceObjects []*proto.AccountBalance
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want []*proto.AccountBalance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.populateBalanceObjWithCreditAcc(tt.args.account, tt.args.balanceObjects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.populateBalanceObjWithCreditAcc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_populateBalanceObjWithDepositAcct(t *testing.T) {
	type args struct {
		account        *proto.AccountsResponseMetadata
		balanceObjects []*proto.AccountBalance
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want []*proto.AccountBalance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.populateBalanceObjWithDepositAcct(tt.args.account, tt.args.balanceObjects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.populateBalanceObjWithDepositAcct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_getAccountBalances(t *testing.T) {
	type args struct {
		balanceObjects []*proto.AccountBalance
		account        *proto.AccountsResponseMetadata
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want []*proto.AccountBalance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.getAccountBalances(tt.args.balanceObjects, tt.args.account); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.getAccountBalances() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestServer_getVirtualAcct(t *testing.T) {
	type args struct {
		ctx     context.Context
		vAcctID uint64
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.AccountsResponseMetadata
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.getVirtualAcct(tt.args.ctx, tt.args.vAcctID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.getVirtualAcct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.getVirtualAcct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_getVirtualAcctFromDatastore(t *testing.T) {
	type args struct {
		ctx     context.Context
		vAcctID uint64
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.VirtualAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.getVirtualAcctFromDatastore(tt.args.ctx, tt.args.vAcctID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.getVirtualAcctFromDatastore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.getVirtualAcctFromDatastore() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
