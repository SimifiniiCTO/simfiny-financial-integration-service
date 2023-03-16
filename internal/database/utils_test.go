package database

import (
	"context"
	"reflect"
	"testing"

	core_database "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-database"
	"go.uber.org/zap"
)

func Test_connectToDatabase(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *ConnectionParameters
		log    *zap.Logger
		models []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *core_database.DatabaseConn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := connectToDatabase(tt.args.ctx, tt.args.params, tt.args.log, tt.args.models...)
			if (err != nil) != tt.wantErr {
				t.Errorf("connectToDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectToDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pingDatabase(t *testing.T) {
	type args struct {
		ctx    context.Context
		dbConn *core_database.DatabaseConn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pingDatabase(tt.args.ctx, tt.args.dbConn); (err != nil) != tt.wantErr {
				t.Errorf("pingDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_configureConnectionString(t *testing.T) {
	type args struct {
		ctx      context.Context
		host     string
		user     string
		password string
		dbname   string
		port     int
		sslMode  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configureConnectionString(tt.args.ctx, tt.args.host, tt.args.user, tt.args.password, tt.args.dbname, tt.args.port, tt.args.sslMode); got != tt.want {
				t.Errorf("configureConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setConnectionConfigs(t *testing.T) {
	type args struct {
		ctx    context.Context
		dbConn *core_database.DatabaseConn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setConnectionConfigs(tt.args.ctx, tt.args.dbConn)
		})
	}
}

func Test_migrateSchemas(t *testing.T) {
	type args struct {
		ctx    context.Context
		db     *core_database.DatabaseConn
		log    *zap.Logger
		models []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := migrateSchemas(tt.args.ctx, tt.args.db, tt.args.log, tt.args.models...); (err != nil) != tt.wantErr {
				t.Errorf("migrateSchemas() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
