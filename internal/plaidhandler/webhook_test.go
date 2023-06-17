package plaidhandler

import (
	"testing"
	"time"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/helper"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/pointer"
	"github.com/plaid/plaid-go/v12/plaid"
	"github.com/stretchr/testify/assert"
)

func TestNewWebhookVerificationKeyFromPlaid(t *testing.T) {
	type args struct {
		input plaid.JWKPublicKey
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{
				input: plaid.JWKPublicKey{
					Alg:       "alg",
					Crv:       "crv",
					Kid:       "kid",
					Kty:       "kty",
					Use:       "use",
					X:         "x",
					Y:         "y",
					CreatedAt: int32(helper.GenerateRandomId(100, 10000)),
					ExpiredAt: *plaid.NewNullableInt32(pointer.Int32P(int32(helper.GenerateRandomId(100, 10000)))),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWebhookVerificationKeyFromPlaid(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWebhookVerificationKeyFromPlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestNewInMemoryWebhookVerification(t *testing.T) {
	type args struct {
		cleanupInterval time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{
				cleanupInterval: time.Second * 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInMemoryWebhookVerification(tt.args.cleanupInterval)
			if (got == nil) != tt.wantErr {
				t.Errorf("NewInMemoryWebhookVerification() error = %v, wantErr %v", got, tt.wantErr)
				return
			}
		})
	}
}
