package inmemoryverifier

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"
)

func TestNewWebhookVerificationKeyFromPlaid(t *testing.T) {
	type args struct {
		input plaid.JWKPublicKey
	}
	tests := []struct {
		name    string
		args    args
		want    plaidhandler.WebhookVerificationKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWebhookVerificationKeyFromPlaid(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWebhookVerificationKeyFromPlaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebhookVerificationKeyFromPlaid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInMemoryWebhookVerification(t *testing.T) {
	type args struct {
		log             *zap.Logger
		plaid           *plaidhandler.PlaidWrapper
		cleanupInterval time.Duration
	}
	tests := []struct {
		name string
		args args
		want WebhookVerification
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemoryWebhookVerification(tt.args.log, tt.args.plaid, tt.args.cleanupInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemoryWebhookVerification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryWebhookVerification_GetVerificationKey(t *testing.T) {
	type args struct {
		ctx   context.Context
		keyId string
	}
	tests := []struct {
		name    string
		m       *memoryWebhookVerification
		args    args
		want    *keyfunc.JWKS
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GetVerificationKey(tt.args.ctx, tt.args.keyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryWebhookVerification.GetVerificationKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryWebhookVerification.GetVerificationKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryWebhookVerification_cacheWorker(t *testing.T) {
	tests := []struct {
		name string
		m    *memoryWebhookVerification
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.cacheWorker()
		})
	}
}

func Test_memoryWebhookVerification_cleanup(t *testing.T) {
	tests := []struct {
		name string
		m    *memoryWebhookVerification
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.cleanup()
		})
	}
}

func Test_memoryWebhookVerification_Close(t *testing.T) {
	tests := []struct {
		name    string
		m       *memoryWebhookVerification
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Close(); (err != nil) != tt.wantErr {
				t.Errorf("memoryWebhookVerification.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
