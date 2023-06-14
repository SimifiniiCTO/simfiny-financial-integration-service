package plaidhandler

import (
	"context"
	"reflect"
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
		want    *SyncResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Sync(tt.args.ctx, tt.args.cursor, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaidWrapper.Sync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaidWrapper.Sync() = %v, want %v", got, tt.want)
			}
		})
	}
}
