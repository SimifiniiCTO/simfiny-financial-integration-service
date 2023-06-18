package database

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/api/v1"
)

func TestDb_AddInvestmentHoldings(t *testing.T) {
	type args struct {
		ctx                 context.Context
		userId              *uint64
		investmentAccountId *uint64
		investmentHoldings  []*schema.InvesmentHolding
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    *schema.InvestmentAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.AddInvestmentHoldings(tt.args.ctx, tt.args.userId, tt.args.investmentAccountId, tt.args.investmentHoldings)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.AddInvestmentHoldings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.AddInvestmentHoldings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_UpdateInvestmentHoldings(t *testing.T) {
	type args struct {
		ctx                 context.Context
		userId              *uint64
		investmentAccountId *uint64
		investmentHoldings  []*schema.InvesmentHolding
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    *schema.InvestmentAccount
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.UpdateInvestmentHoldings(tt.args.ctx, tt.args.userId, tt.args.investmentAccountId, tt.args.investmentHoldings)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateInvestmentHoldings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.UpdateInvestmentHoldings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetInvestmentHoldings(t *testing.T) {
	type args struct {
		ctx                  context.Context
		investmentHoldingIds []uint64
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    []*schema.InvesmentHolding
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetInvestmentHoldings(tt.args.ctx, tt.args.investmentHoldingIds...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetInvestmentHoldings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetInvestmentHoldings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_DeleteInvestmentHoldings(t *testing.T) {
	type args struct {
		ctx                  context.Context
		investmentHoldingIds []uint64
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
			if err := tt.db.DeleteInvestmentHoldings(tt.args.ctx, tt.args.investmentHoldingIds...); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteInvestmentHoldings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
