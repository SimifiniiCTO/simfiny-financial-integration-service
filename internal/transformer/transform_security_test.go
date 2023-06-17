package transformer

import (
	"errors"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

func TestTransformPlaidInvestmentSecurity_EmptyInput(t *testing.T) {
	securities := []plaid.Security{}

	expectedErr := errors.New("invalid input argument. securities cannot be empty")

	result, err := TransformPlaidInvestmentSecurity(securities)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if result != nil {
		t.Errorf("Expected nil result, but got: %v", result)
	}

	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
	}
}

func TestTransformPlaidInvestmentSecurity(t *testing.T) {
	type args struct {
		securities []plaid.Security
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "EmptyInput",
			args:    args{securities: []plaid.Security{}},
			wantErr: true,
		},
		{
			name:    "SingleSecurity",
			args:    args{securities: generateMultipleSecurity(1)},
			wantErr: false,
		},
		{
			name:    "MultipleSecurities",
			args:    args{securities: generateMultipleSecurity(10)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformPlaidInvestmentSecurity(tt.args.securities)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformPlaidInvestmentSecurity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got) != len(tt.args.securities) {
					t.Errorf("Expected result length: %v, but got: %v", len(tt.args.securities), len(got))
				}
			}
		})
	}
}

func TestTransformSinglePlaidInvestmentSecurity(t *testing.T) {
	type args struct {
		security plaid.Security
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ValidSecurity",
			args:    args{security: *generateSinglSecurity()},
			wantErr: false,
		},
		{
			name:    "InvalidSecurity",
			args:    args{security: plaid.Security{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TransformSinglePlaidInvestmentSecurity(tt.args.security)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}
