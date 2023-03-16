package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/generated/api/v1"
)

func TestServer_DeleteUserProfile(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *proto.DeleteUserProfileRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.DeleteUserProfileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteUserProfile(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteUserProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
