package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestTransformStudentloanAccount(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformStudentloanAccount(tt.args.studentLoans, tt.args.acctIDToTypeMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformStudentloanAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformStudentloanAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
