package grpc

import (
	"reflect"
	"testing"

	"google.golang.org/grpc"
)

func TestServer_RegisterGrpcServer(t *testing.T) {
	type args struct {
		srv *grpc.Server
	}
	tests := []struct {
		name   string
		server *Server
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.server.RegisterGrpcServer(tt.args.srv)
		})
	}
}

func TestParams_Validate(t *testing.T) {
	tests := []struct {
		name    string
		p       *Params
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Params.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewServer(t *testing.T) {
	type args struct {
		param *Params
	}
	tests := []struct {
		name    string
		args    args
		want    *Server
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServer(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
