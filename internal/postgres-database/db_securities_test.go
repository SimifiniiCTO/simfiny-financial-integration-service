package database

import (
	"context"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestDb_UpdateSecurities(t *testing.T) {
	type args struct {
		ctx                 context.Context
		investmentAccountId *uint64
		securities          []*schema.InvestmentSecurity
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.UpdateSecurities(tt.args.ctx, tt.args.investmentAccountId, tt.args.securities); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateSecurities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
