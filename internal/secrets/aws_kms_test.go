package secrets

import (
	"context"
	"reflect"
	"testing"
)

func TestAWSKMSConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       *AWSKMSConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AWSKMSConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAWSKMS(t *testing.T) {
	type args struct {
		config AWSKMSConfig
	}
	tests := []struct {
		name    string
		args    args
		want    KeyManagement
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAWSKMS(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAWSKMS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAWSKMS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAWSKMS_Encrypt(t *testing.T) {
	type args struct {
		ctx   context.Context
		input []byte
	}
	tests := []struct {
		name        string
		a           *AWSKMS
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
			gotKeyID, gotVersion, gotResult, err := tt.a.Encrypt(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AWSKMS.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKeyID != tt.wantKeyID {
				t.Errorf("AWSKMS.Encrypt() gotKeyID = %v, want %v", gotKeyID, tt.wantKeyID)
			}
			if gotVersion != tt.wantVersion {
				t.Errorf("AWSKMS.Encrypt() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AWSKMS.Encrypt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestAWSKMS_Decrypt(t *testing.T) {
	type args struct {
		ctx     context.Context
		keyID   string
		version string
		input   []byte
	}
	tests := []struct {
		name       string
		a          *AWSKMS
		args       args
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.a.Decrypt(tt.args.ctx, tt.args.keyID, tt.args.version, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AWSKMS.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AWSKMS.Decrypt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
