package plaidhandler

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
)

// setup sets up a database connection to the test db node
func setup() {
	var err error
	plaidTestClient, err = GetPlaidWrapperForTest()
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestPlaidWrapper_GetWebhookVerificationKey(t *testing.T) {
	type args struct {
		ctx   context.Context
		keyId string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    *WebhookVerificationKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetWebhookVerificationKey(tt.args.ctx, tt.args.keyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetWebhookVerificationKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetWebhookVerificationKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_GetInstitution(t *testing.T) {
	type args struct {
		ctx           context.Context
		institutionId string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		want    *plaid.Institution
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetInstitution(tt.args.ctx, tt.args.institutionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.GetInstitution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.GetInstitution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_GetWebhooksURL(t *testing.T) {
	tests := []struct {
		name string
		p    *PlaidWrapper
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetWebhooksURL(); got != tt.want {
				t.Errorf("PlaidWrapper.GetWebhooksURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_GetRedirectURI(t *testing.T) {
	tests := []struct {
		name string
		p    *PlaidWrapper
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetRedirectURI(); got != tt.want {
				t.Errorf("PlaidWrapper.GetRedirectURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_ToJSON(t *testing.T) {
	tests := []struct {
		name string
		p    *PlaidWrapper
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ToJSON(); got != tt.want {
				t.Errorf("PlaidWrapper.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaidWrapper_EnabledProductsToString(t *testing.T) {
	tests := []struct {
		name string
		p    *PlaidWrapper
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.EnabledProductsToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.EnabledProductsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
