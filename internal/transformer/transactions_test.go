package transformer

import (
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransactionFromPlaid(t *testing.T) {
	type args struct {
		input *plaid.Transaction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid transaction",
			args: args{
				input: generatePlaidTransaction(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTransactionFromPlaid(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTransactionFromPlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
