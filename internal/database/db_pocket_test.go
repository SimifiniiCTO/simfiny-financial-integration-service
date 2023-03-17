package database

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestDb_CreatePocket(t *testing.T) {
	type args struct {
		ctx           context.Context
		bankAccountID uint64
		pocket        *schema.Pocket
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		want    *schema.Pocket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.CreatePocket(tt.args.ctx, tt.args.bankAccountID, tt.args.pocket)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreatePocket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.CreatePocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_DeletePocket(t *testing.T) {
	type args struct {
		ctx      context.Context
		pocketID uint64
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.DeletePocket(tt.args.ctx, tt.args.pocketID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeletePocket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_GetPocket(t *testing.T) {
	type args struct {
		ctx      context.Context
		pocketID uint64
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		want    *schema.Pocket
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetPocket(tt.args.ctx, tt.args.pocketID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetPocket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetPocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_UpdatePocket(t *testing.T) {
	type args struct {
		ctx    context.Context
		pocket *schema.Pocket
	}
	tests := []struct {
		name    string
		d       *Db
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UpdatePocket(tt.args.ctx, tt.args.pocket); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdatePocket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
