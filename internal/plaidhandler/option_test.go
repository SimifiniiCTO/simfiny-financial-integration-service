package plaidhandler

import (
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"
)

func TestWithInstrumentation(t *testing.T) {
	type args struct {
		instrumentation *instrumentation.Client
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithInstrumentation(tt.args.instrumentation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithInstrumentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLogger(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithEnvironment(t *testing.T) {
	type args struct {
		environment plaid.Environment
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithEnvironment(tt.args.environment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithClientID(t *testing.T) {
	type args struct {
		clientID string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithClientID(tt.args.clientID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithClientID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSecretKey(t *testing.T) {
	type args struct {
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithSecretKey(tt.args.secretKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSecretKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithOauthDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithOauthDomain(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithOauthDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithWebhooksDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithWebhooksDomain(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithWebhooksDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithWebhooksEnabled(t *testing.T) {
	type args struct {
		enabled bool
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithWebhooksEnabled(tt.args.enabled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithWebhooksEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithProducts(t *testing.T) {
	type args struct {
		products []string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithProducts(tt.args.products); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *PlaidWrapper
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
