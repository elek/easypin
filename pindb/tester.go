// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package pindb

import (
	"context"
	"strings"
	"testing"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"storj.io/common/testcontext"
	"storj.io/private/dbutil"
	"storj.io/private/dbutil/pgtest"
	"storj.io/private/dbutil/pgutil"
	"storj.io/private/dbutil/tempdb"
)

// TestRun creates new storjscan test database, create tables and execute test function against that db.
func TestRun(t *testing.T, test func(ctx *testcontext.Context, t *testing.T, db *DB)) {
	t.Run("Postgres", func(t *testing.T) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		connStr := pgtest.PickPostgres(t)

		db, err := OpenDB(ctx, zaptest.NewLogger(t), connStr, t.Name(), "T")
		if err != nil {
			t.Fatal(err)
		}
		defer ctx.Check(db.Close)

		err = db.MigrateToLatest(ctx)
		if err != nil {
			t.Fatal(err)
		}

		test(ctx, t, db.DB)
	})
}

// TestDB is test storjscan database with unique schema which performs cleanup on close.
type TestDB struct {
	*DB
	tempDB *dbutil.TempDatabase
}

// OpenDB opens new unique temp storjscan test database.
func OpenDB(ctx context.Context, log *zap.Logger, connStr, testName, category string) (*TestDB, error) {
	schemaSuffix := pgutil.CreateRandomTestingSchemaName(6)
	schemaName := schemaName(testName, schemaSuffix, category)

	tempDB, err := tempdb.OpenUnique(ctx, connStr, schemaName)
	if err != nil {
		return nil, err
	}
	pinDB, err := Open(ctx, log, tempDB.ConnStr)
	if err != nil {
		return nil, errs.Combine(err, tempDB.Close())
	}

	return &TestDB{
		DB:     pinDB,
		tempDB: tempDB,
	}, nil
}

// Close closes test database and performs cleanup.
func (db *TestDB) Close() error {
	return errs.Combine(db.DB.Close(), db.tempDB.Close())
}

// schemaName create new postgres db schema name for testing.
func schemaName(testName, suffix, category string) string {
	maxTestNameLength := 64 - len(suffix) - len(category)
	if len(testName) > maxTestNameLength {
		testName = testName[:maxTestNameLength]
	}
	return strings.ToLower(testName + "/" + suffix + "/" + category)
}
