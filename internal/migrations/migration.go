package migrations

import (
	"context"
	"embed"
	"fmt"

	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chmigrate"
)

var Migrations = chmigrate.NewMigrations()

//go:embed *.sql
var sqlMigrations embed.FS

type MigrationEngine struct {
	migrator *chmigrate.Migrations
	db       *ch.DB
	sm       embed.FS
}

func New(db *ch.DB) (*MigrationEngine, error) {
	migrator := Migrations
	// discover any sql migration scritps
	if err := migrator.Discover(sqlMigrations); err != nil {
		return nil, err
	}

	engine := &MigrationEngine{
		migrator: migrator,
		db:       db,
		sm:       sqlMigrations,
	}

	return engine, nil
}

// Migrate migrates the various schema migrations
func (m *MigrationEngine) Migrate(ctx context.Context) error {
	var (
		migrator = chmigrate.NewMigrator(m.db, m.migrator)
		err      error
		// rollback func(ctx context.Context) = func(ctx context.Context) {
		// 	// if err := migrator.Lock(ctx); err != nil {
		// 	// 	return
		// 	// }
		// 	// defer migrator.Unlock(ctx) //nolint:errcheck

		// 	group, err := migrator.Rollback(ctx)
		// 	if err != nil {
		// 		return
		// 	}

		// 	if group.IsZero() {
		// 		fmt.Printf("there are no groups to roll back\n")
		// 		return
		// 	}
		// }
	)

	// if err := migrator.Lock(ctx); err != nil {
	// 	return err
	// }
	// defer migrator.Unlock(ctx) //nolint:errcheck
	// defer func(ctx context.Context) {
	// 	if err != nil {
	// 		rollback(ctx)
	// 	}
	// }(ctx)
	err = migrator.Init(ctx)
	if err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("applied %d migrations\n", len(group.Migrations))

	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return nil
	}

	return nil
}
