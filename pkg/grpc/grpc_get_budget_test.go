package grpc

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestServer_GetBudget(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *proto.GetBudgetRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *proto.GetBudgetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetBudget(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetBudget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetBudget() = %v, want %v", got, tt.want)
			}
		})
	}
}
