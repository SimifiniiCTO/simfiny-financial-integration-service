package database

import (
	"context"
	"reflect"
	"testing"

	schema "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/generated/api/v1"
)

func TestDb_LinkExistsForItem(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint64
		itemID string
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.LinkExistsForItem(tt.args.ctx, tt.args.userID, tt.args.itemID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.LinkExistsForItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Db.LinkExistsForItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetLink(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint64
		linkID uint64
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    *schema.Link
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetLink(tt.args.ctx, tt.args.userID, tt.args.linkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_DeleteLink(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint64
		linkID uint64
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
			if err := tt.db.DeleteLink(tt.args.ctx, tt.args.userID, tt.args.linkID); (err != nil) != tt.wantErr {
				t.Errorf("Db.DeleteLink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDb_CreateLink(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint64
		link   *schema.Link
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    *schema.Link
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.CreateLink(tt.args.ctx, tt.args.userID, tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.CreateLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.CreateLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetLinkByItemId(t *testing.T) {
	type args struct {
		ctx    context.Context
		itemId string
	}
	tests := []struct {
		name    string
		db      *Db
		args    args
		want    *schema.Link
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetLinkByItemId(tt.args.ctx, tt.args.itemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetLinkByItemId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.GetLinkByItemId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_UpdateLink(t *testing.T) {
	type args struct {
		ctx  context.Context
		link *schema.Link
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
			if err := tt.db.UpdateLink(tt.args.ctx, tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("Db.UpdateLink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
