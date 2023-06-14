package plaidhandler

import (
	"reflect"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v12/plaid"
)

func TestNewItemTokenFromPlaid(t *testing.T) {
	type args struct {
		input plaid.ItemPublicTokenExchangeResponse
	}
	tests := []struct {
		name    string
		args    args
		want    ItemToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewItemTokenFromPlaid(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewItemTokenFromPlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemTokenFromPlaid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidLinkToken_Token(t *testing.T) {
	tests := []struct {
		name string
		p    *PlaidLinkToken
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Token(); got != tt.want {
				t.Errorf("PlaidLinkToken.Token() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidLinkToken_Expiration(t *testing.T) {
	tests := []struct {
		name string
		p    *PlaidLinkToken
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Expiration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidLinkToken.Expiration() = %v, want %v", got, tt.want)
			}
		})
	}
}
