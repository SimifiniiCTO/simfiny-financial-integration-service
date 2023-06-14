package secrets

import (
	"context"
	"reflect"
	"testing"
)

func TestNewMockAwsKms(t *testing.T) {
	tests := []struct {
		name    string
		want    KeyManagement
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMockAwsKms()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMockAwsKms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockAwsKms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockKeyManagement_Decrypt(t *testing.T) {
	type args struct {
		ctx     context.Context
		keyID   string
		version string
		input   []byte
	}
	tests := []struct {
		name       string
		m          *MockKeyManagement
		args       args
		wantResult []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.m.Decrypt(tt.args.ctx, tt.args.keyID, tt.args.version, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockKeyManagement.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MockKeyManagement.Decrypt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMockKeyManagement_Encrypt(t *testing.T) {
	type args struct {
		ctx   context.Context
		input []byte
	}
	tests := []struct {
		name        string
		m           *MockKeyManagement
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
			gotKeyID, gotVersion, gotResult, err := tt.m.Encrypt(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockKeyManagement.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKeyID != tt.wantKeyID {
				t.Errorf("MockKeyManagement.Encrypt() gotKeyID = %v, want %v", gotKeyID, tt.wantKeyID)
			}
			if gotVersion != tt.wantVersion {
				t.Errorf("MockKeyManagement.Encrypt() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MockKeyManagement.Encrypt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_generateRandomBytes(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomBytes(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRandomBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomString(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
