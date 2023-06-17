package plaidhandler

import (
	"context"
	"testing"
)

func TestPlaidWrapper_Sync(t *testing.T) {
	type args struct {
		ctx         context.Context
		cursor      *string
		accessToken *string
	}
	tests := []struct {
		name    string
		p       *PlaidWrapper
		args    args
		wantErr bool
	}{
		{
			name: "test",
			p:    plaidTestClient,
			args: args{
				ctx:         context.Background(),
				cursor:      nil,
				accessToken: &testAccessToken,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Sync(tt.args.ctx, tt.args.cursor, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.Sync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got.New) == 0 && len(got.Updated) == 0 && len(got.Deleted) == 0 {
					t.Errorf("PlaidWrapper.Sync() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
