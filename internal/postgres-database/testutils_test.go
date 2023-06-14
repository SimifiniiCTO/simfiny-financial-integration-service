package database

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func Test_createUserProfileAndBankAccountForTest(t *testing.T) {
	type args struct {
		t    *testing.T
		ctx  context.Context
		conn *Db
	}
	tests := []struct {
		name  string
		args  args
		want  *schema.UserProfile
		want1 *schema.BankAccount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := createUserProfileAndBankAccountForTest(tt.args.t, tt.args.ctx, tt.args.conn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createUserProfileAndBankAccountForTest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("createUserProfileAndBankAccountForTest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_generateBankAccount(t *testing.T) {
	tests := []struct {
		name string
		want *schema.BankAccount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateBankAccount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateBankAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
