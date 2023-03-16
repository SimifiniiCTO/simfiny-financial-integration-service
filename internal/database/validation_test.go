package database

import (
	"context"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/proto"
)

func TestDb_validateVirtualAcct(t *testing.T) {
	type args struct {
		ctx      context.Context
		vAcctOrm *schema.VirtualAccountORM
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
			if err := tt.db.validateVirtualAcct(tt.args.ctx, tt.args.vAcctOrm); (err != nil) != tt.wantErr {
				t.Errorf("Db.validateVirtualAcct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
