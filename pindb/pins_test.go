package pindb

import (
	"github.com/stretchr/testify/require"
	"storj.io/common/testcontext"
	"testing"
)

func TestPinMissing(t *testing.T) {
	TestRun(t, func(ctx *testcontext.Context, t *testing.T, db *DB) {

		err := db.Pins().Create(ctx, "tx1", 1, "cid1", 10)
		require.NoError(t, err)

		err = db.Pins().Create(ctx, "tx2", 2, "cid1", 10)
		require.NoError(t, err)

		err = db.Pins().Create(ctx, "tx3", 3, "cid2", 10)
		require.NoError(t, err)

		missing, err := db.Pins().FindNew(ctx)
		require.NoError(t, err)
		require.Equal(t, 2, len(missing))
	})

}
