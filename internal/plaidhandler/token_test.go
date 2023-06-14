package plaidhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
)

func TestPlaidWrapper_CreateLinkToken(t *testing.T) {
	type args struct {
		ctx     context.Context
		options *LinkTokenOptions
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    LinkToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.CreateLinkToken(tt.args.ctx, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.CreateLinkToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.CreateLinkToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_ExchangePublicToken(t *testing.T) {
	type args struct {
		ctx         context.Context
		publicToken string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    *ItemToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.ExchangePublicToken(tt.args.ctx, tt.args.publicToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.ExchangePublicToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.ExchangePublicToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_getAccessTokenForSandboxAcct(t *testing.T) {
	tests := []struct {
		name    string
		p       *PlaidWrapper
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.getAccessTokenForSandboxAcct()
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.getAccessTokenForSandboxAcct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PlaidWrapper.getAccessTokenForSandboxAcct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_getPlublicTokenForSandboxAcct(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    plaid.SandboxPublicTokenCreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.getPlublicTokenForSandboxAcct(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.getPlublicTokenForSandboxAcct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.getPlublicTokenForSandboxAcct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_configureClient(t *testing.T) {
	tests := []struct {
		name string
		p    *PlaidWrapper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.configureClient()
		})
	}
}
