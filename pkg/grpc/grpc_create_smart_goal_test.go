package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestServer_CreateSmartGoal(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *proto.CreateSmartGoalRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.CreateSmartGoalResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateSmartGoal(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateSmartGoal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateSmartGoal() = %v, want %v", got, tt.want)
			}
		})
	}
}
