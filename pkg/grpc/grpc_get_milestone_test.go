package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_GetMilestone(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *proto.GetMilestoneRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.GetMilestoneResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMilestone(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetMilestone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetMilestone() = %v, want %v", got, tt.want)
			}
		})
	}
}
