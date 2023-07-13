package plaidhandler

import (
	"context"
	"reflect"
	"testing"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/plaid/plaid-go/v14/plaid"
	"go.uber.org/zap"
)

func TestPlaidWrapper_TriggerWebhookForTest(t *testing.T) {

	type args struct {
		ctx         context.Context
		clientId    string
		secret      string
		accesstoken string
		webhookCode string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		wantErr bool
	}{
		{
			name: "trigger webhook for test",
			p:    plaidTestClient,
			args: args{
				ctx:         context.Background(),
				clientId:    testClientId,
				secret:      "",
				accesstoken: testAccessToken,
				webhookCode: "SYNC_UPDATES_AVAILABLE",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.TriggerWebhookForTest(tt.args.ctx, tt.args.clientId, tt.args.secret, tt.args.accesstoken, tt.args.webhookCode); (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.TriggerWebhookForTest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlaidWrapper_CreateLinkToken(t *testing.T) {
	type fields struct {
		client             *plaid.APIClient
		InstrumentationSdk *instrumentation.Client
		Logger             *zap.Logger
		Environment        plaid.Environment
		ClientID           string
		SecretKey          string
		OAuthDomain        string
		WebhooksDomain     string
		WebhooksEnabled    bool
		EnabledProducts    []plaid.Products
	}
	type args struct {
		ctx     context.Context
		options *LinkTokenOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    LinkToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaidWrapper{
				client:             tt.fields.client,
				InstrumentationSdk: tt.fields.InstrumentationSdk,
				Logger:             tt.fields.Logger,
				Environment:        tt.fields.Environment,
				ClientID:           tt.fields.ClientID,
				SecretKey:          tt.fields.SecretKey,
				OAuthDomain:        tt.fields.OAuthDomain,
				WebhooksDomain:     tt.fields.WebhooksDomain,
				WebhooksEnabled:    tt.fields.WebhooksEnabled,
				EnabledProducts:    tt.fields.EnabledProducts,
			}
			got, err := p.CreateLinkToken(tt.args.ctx, tt.args.options)
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
	type fields struct {
		client             *plaid.APIClient
		InstrumentationSdk *instrumentation.Client
		Logger             *zap.Logger
		Environment        plaid.Environment
		ClientID           string
		SecretKey          string
		OAuthDomain        string
		WebhooksDomain     string
		WebhooksEnabled    bool
		EnabledProducts    []plaid.Products
	}
	type args struct {
		ctx         context.Context
		publicToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ItemToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaidWrapper{
				client:             tt.fields.client,
				InstrumentationSdk: tt.fields.InstrumentationSdk,
				Logger:             tt.fields.Logger,
				Environment:        tt.fields.Environment,
				ClientID:           tt.fields.ClientID,
				SecretKey:          tt.fields.SecretKey,
				OAuthDomain:        tt.fields.OAuthDomain,
				WebhooksDomain:     tt.fields.WebhooksDomain,
				WebhooksEnabled:    tt.fields.WebhooksEnabled,
				EnabledProducts:    tt.fields.EnabledProducts,
			}
			got, err := p.ExchangePublicToken(tt.args.ctx, tt.args.publicToken)
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

func TestPlaidWrapper_configureClient(t *testing.T) {
	type fields struct {
		client             *plaid.APIClient
		InstrumentationSdk *instrumentation.Client
		Logger             *zap.Logger
		Environment        plaid.Environment
		ClientID           string
		SecretKey          string
		OAuthDomain        string
		WebhooksDomain     string
		WebhooksEnabled    bool
		EnabledProducts    []plaid.Products
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlaidWrapper{
				client:             tt.fields.client,
				InstrumentationSdk: tt.fields.InstrumentationSdk,
				Logger:             tt.fields.Logger,
				Environment:        tt.fields.Environment,
				ClientID:           tt.fields.ClientID,
				SecretKey:          tt.fields.SecretKey,
				OAuthDomain:        tt.fields.OAuthDomain,
				WebhooksDomain:     tt.fields.WebhooksDomain,
				WebhooksEnabled:    tt.fields.WebhooksEnabled,
				EnabledProducts:    tt.fields.EnabledProducts,
			}
			p.configureClient()
		})
	}
}
