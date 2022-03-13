// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package pindb

import (
	"context"
	"github.com/elek/easypin/pindb/dbx"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/private/dbutil"
	"storj.io/private/dbutil/pgutil"
	"storj.io/private/migrate"
	"storj.io/private/tagsql"
)

var (
	mon = monkit.Package()

	// Error is the default easypindb errs class.
	Error = errs.Class("easypindb")
)

// DB is easypindb database.
type DB struct {
	*dbx.DB
	log            *zap.Logger
	driver         string
	source         string
	implementation dbutil.Implementation
	migrationDB    tagsql.DB
}

// Open creates instance of storjscan DB.
func Open(ctx context.Context, log *zap.Logger, databaseURL string) (*DB, error) {
	driver, source, impl, err := dbutil.SplitConnStr(databaseURL)
	if err != nil {
		return nil, err
	}
	if impl != dbutil.Postgres {
		return nil, Error.New("unsupported driver %q", driver)
	}

	source, err = pgutil.CheckApplicationName(source, "easypin")
	if err != nil {
		return nil, err
	}

	dbxDB, err := dbx.Open(driver, source)
	if err != nil {
		return nil, Error.New("failed opening database via DBX at %q: %v", source, err)
	}
	log.Debug("Connected to:", zap.String("pindb source", source))

	dbutil.Configure(ctx, dbxDB.DB, "easypin", mon)

	db := &DB{
		DB:             dbxDB,
		log:            log,
		driver:         driver,
		source:         source,
		implementation: impl,
	}
	db.migrationDB = db

	return db, nil
}

// MigrateToLatest migrates pindb to the latest version.
func (db *DB) MigrateToLatest(ctx context.Context) error {
	var migration *migrate.Migration

	switch db.implementation {
	case dbutil.Postgres:
		migration = db.PostgresMigration()
	default:
		return migrate.Create(ctx, "database", db.DB)
	}
	return migration.Run(ctx, db.log)
}

// Pins creates new PinDB with current DB connection.
func (db *DB) Pins() *PinDB {
	return &PinDB{db: db.DB}
}

// PostgresMigration returns steps needed for migrating postgres database.
func (db *DB) PostgresMigration() *migrate.Migration {
	return &migrate.Migration{
		Table: "versions",
		Steps: []*migrate.Step{
			{
				DB:          &db.migrationDB,
				Description: "Initial setup",
				Version:     0,
				Action: migrate.SQL{
					`CREATE TABLE nodes (
	cid text NOT NULL,
	expired_at timestamp with time zone NOT NULL,
	amount bigint NOT NULL,
	created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY ( cid )
);
CREATE TABLE pins (
	tx text NOT NULL,
	ix integer NOT NULL,
	cid text NOT NULL,
	amount bigint NOT NULL,
	created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY ( tx, ix )
);
`,
				},
			},
			{
				DB:          &db.migrationDB,
				Description: "Initial setup",
				Version:     1,
				Action: migrate.SQL{
					`ALTER TABLE pins ADD COLUMN retry integer NOT NULL DEFAULT 0;
					ALTER TABLE pins ADD COLUMN error text NOT NULL DEFAULT '';
					ALTER TABLE pins ADD COLUMN parse boolean NOT NULL DEFAULT false;
`,
				},
			},
		},
	}
}
