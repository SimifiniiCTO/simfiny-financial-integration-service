package transformer

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestNewTransactionFromPlaid(t *testing.T) {
	type args struct {
		input plaid.Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTransactionFromPlaid(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTransactionFromPlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionFromPlaid() = %v, want %v", got, tt.want)
			}
		})
	}
}
