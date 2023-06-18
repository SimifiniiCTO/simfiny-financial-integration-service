package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestServer_GetMilestonesBySmartGoalId(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *proto.GetMilestonesBySmartGoalIdRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.GetMilestonesBySmartGoalIdResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMilestonesBySmartGoalId(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetMilestonesBySmartGoalId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetMilestonesBySmartGoalId() = %v, want %v", got, tt.want)
			}
		})
	}
}
