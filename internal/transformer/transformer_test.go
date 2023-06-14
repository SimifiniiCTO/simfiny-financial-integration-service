package transformer

import (
	"reflect"
	"testing"

	"github.com/plaid/plaid-go/v12/plaid"
)

func Test_getAccountMetadata(t *testing.T) {
	type args struct {
		accounts *[]plaid.AccountBase
	}
	tests := []struct {
		name string
		args args
		want map[string]*accountMetadata
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAccountMetadata(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAccountMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
