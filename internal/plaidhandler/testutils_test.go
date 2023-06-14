package plaidhandler

import (
	"reflect"
	"testing"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/plaid/plaid-go/v12/plaid"
	"go.uber.org/zap"
)

func TestNewMockPlaidClient(t *testing.T) {
	tests := []struct {
		name string
		want *plaid.APIClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMockPlaidClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockPlaidClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newNewRelicClient(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *newrelic.Application
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newNewRelicClient(tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("newNewRelicClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNewRelicClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newLogger(t *testing.T) {
	tests := []struct {
		name string
		want *zap.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPlaidWrapperForTest(t *testing.T) {
	tests := []struct {
		name    string
		want    *PlaidWrapper
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPlaidWrapperForTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlaidWrapperForTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlaidWrapperForTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
