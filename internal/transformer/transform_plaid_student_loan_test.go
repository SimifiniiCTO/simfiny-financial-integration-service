package transformer

import (
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

func TestTransformStudentloanAccount(t *testing.T) {
	studentLoanAccts, acctMeta := generateManyPlaidStudentLoanAndAccountMetadata(10)

	type args struct {
		studentLoans    *[]plaid.StudentLoan
		acctIDToTypeMap map[string]*accountMetadata
	}
	tests := []struct {
		name    string
		args    args
		want    []*schema.StudentLoanAccount
		wantErr bool
	}{
		{
			name: "valid student loan liability",
			args: args{
				studentLoans:    &studentLoanAccts,
				acctIDToTypeMap: acctMeta,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformStudentloanAccount(tt.args.studentLoans, tt.args.acctIDToTypeMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformStudentloanAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
