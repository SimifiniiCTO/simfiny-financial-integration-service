package database

import (
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

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

func Test_genereateRandomUserProfileForTest(t *testing.T) {
	tests := []struct {
		name string
		want *schema.UserProfile
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genereateRandomUserProfileForTest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genereateRandomUserProfileForTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
