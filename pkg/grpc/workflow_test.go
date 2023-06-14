package grpc

import (
	"context"
	"reflect"
	"testing"

	"go.temporal.io/sdk/client"
)

func TestServer_ExecuteWorkflow(t *testing.T) {
	type args struct {
		ctx      context.Context
		workflow interface{}
		args     []interface{}
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    client.WorkflowRun
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ExecuteWorkflow(tt.args.ctx, tt.args.workflow, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ExecuteWorkflow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ExecuteWorkflow() = %v, want %v", got, tt.want)
			}
		})
	}
}
