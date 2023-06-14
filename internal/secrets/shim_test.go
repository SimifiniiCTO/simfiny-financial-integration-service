package secrets

import (
	"context"
	"reflect"
	"testing"
)

func TestShimKMS_Decrypt(t *testing.T) {
	type args struct {
		ctx     context.Context
		keyID   string
		version string
		input   []byte
	}
	tests := []struct {
		name       string
		s          *ShimKMS
		args       args
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.s.Decrypt(tt.args.ctx, tt.args.keyID, tt.args.version, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShimKMS.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ShimKMS.Decrypt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestShimKMS_Encrypt(t *testing.T) {
	type args struct {
		ctx   context.Context
		input []byte
	}
	tests := []struct {
		name        string
		s           *ShimKMS
		args        args
		wantKeyID   string
		wantVersion string
		wantResult  []byte
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeyID, gotVersion, gotResult, err := tt.s.Encrypt(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShimKMS.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKeyID != tt.wantKeyID {
				t.Errorf("ShimKMS.Encrypt() gotKeyID = %v, want %v", gotKeyID, tt.wantKeyID)
			}
			if gotVersion != tt.wantVersion {
				t.Errorf("ShimKMS.Encrypt() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ShimKMS.Encrypt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
