package plaidhandler

import (
	"context"
	"testing"
)

func TestPlaidWrapper_DeleteItem(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken *string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		wantErr bool
	}{
		{
			name: "delete item",
			p:    plaidTestClient,
			args: args{
				ctx:         context.Background(),
				accessToken: &testAccessToken,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.DeleteItem(tt.args.ctx, tt.args.accessToken); (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
