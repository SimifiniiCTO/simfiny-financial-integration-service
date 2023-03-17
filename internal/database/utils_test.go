package database

import (
	"context"
	"reflect"
	"testing"

	core_database "github.com/SimifiniiCTO/core/core-database"
	newrelic "github.com/newrelic/go-agent"
	"github.com/stretchr/testify/assert"
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
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test connect to database",
			args: args{
				ctx: context.Background(),
				params: &ConnectionParameters{
					Host:         "localhost",
					User:         "service_db",
					Password:     "service_db",
					DatabaseName: "service_db",
					Port:         5433,
					SslMode:      "disable",
				},
				log: zap.NewNop(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := connectToDatabase(tt.args.ctx, tt.args.params, tt.args.log, tt.args.models...)
			if (err != nil) != tt.wantErr {
				t.Errorf("connectToDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == false {
				assert.NotNil(t, got)
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

func TestDb_startDatastoreSpan(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name string
		db   *Db
		args args
		want *newrelic.DatastoreSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.startDatastoreSpan(tt.args.ctx, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db.startDatastoreSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatDbConnectionString(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *ConnectionParameters
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatDbConnectionString(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatDbConnectionString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatDbConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
