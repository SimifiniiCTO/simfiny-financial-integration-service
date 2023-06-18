package clickhousedatabase

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"
)

func Test_generateRandomInvestmentTransaction(t *testing.T) {
	tests := []struct {
		name string
		want *schema.InvestmentTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomInvestmentTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomInvestmentTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomReOccurringTransaction(t *testing.T) {
	tests := []struct {
		name string
		want *schema.ReOccuringTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomReOccurringTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomReOccurringTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomTransaction(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomTransaction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateMultipleInvestmentTransaction(t *testing.T) {
	type args struct {
		numTransactions int
	}
	tests := []struct {
		name string
		args args
		want []*schema.InvestmentTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMultipleInvestmentTransaction(tt.args.numTransactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMultipleInvestmentTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateMultipleTransaction(t *testing.T) {
	type args struct {
		numTransactions int64
	}
	tests := []struct {
		name string
		args args
		want []*schema.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMultipleTransaction(tt.args.numTransactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMultipleTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateMultipleReOcurringTransactions(t *testing.T) {
	type args struct {
		numTransactions int
	}
	tests := []struct {
		name string
		args args
		want []*schema.ReOccuringTransaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMultipleReOcurringTransactions(tt.args.numTransactions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMultipleReOcurringTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomId(t *testing.T) {
	tests := []struct {
		name string
		want *uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRandomId(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateStringList(t *testing.T) {
	type args struct {
		numElements int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateStringList(tt.args.numElements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateStringList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointerUint64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want *uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointerUint64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointerUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
