package plaidhandler

import (
	"testing"
)

func TestNewMockPlaidClient(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMockPlaidClient()
			if (got == nil) != tt.wantErr {
				t.Errorf("NewMockPlaidClient() error = %v, wantErr %v", got, tt.wantErr)
				return
			}
		})
	}
}

func TestGetPlaidWrapperForTest(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPlaidWrapperForTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlaidWrapperForTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got == nil {
					t.Errorf("GetPlaidWrapperForTest() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
