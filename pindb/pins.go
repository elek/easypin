// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package pindb

import (
	"context"
	"github.com/elek/easypin/pindb/dbx"
	"github.com/jackc/pgx/v4"
	"math/big"
	"storj.io/private/dbutil/pgxutil"
	"time"

	"github.com/zeebo/errs"
)

// ErrPinDB indicates about internal wallets DB error.
var ErrPinDB = errs.Class("PinDB")

// PinDB is database stores the pinning events.
//
// architecture: Database
type PinDB struct {
	db *dbx.DB
}

// Pin represents an entry in the pin table.
type Pin struct {
	Cid         string
	Amount      *big.Int
	Transaction string
	LogIndex    uint
	Processed   bool
}

type Cid struct {
	Hash      string
	Pinned    time.Time
	ValidDays uint
}

// Create inserts a new entry in the wallets table.
func (p *PinDB) Create(ctx context.Context, tx string, ix uint, cid string, amount *big.Int) error {
	_, err := p.db.ExecContext(ctx, p.db.Rebind("INSERT INTO pins ( tx, ix, cid, amount ) VALUES ( ?, ?, ?, ?) ON CONFLICT DO NOTHING;"), tx, ix, cid, amount.String())
	if err != nil {
		return ErrPinDB.Wrap(err)
	}
	return ErrPinDB.Wrap(err)
}

func (p *PinDB) FindNew(ctx context.Context) ([]Pin, error) {
	var res []Pin

	rows, err := p.db.All_Pin_By_Processed_Equal_False_OrderBy_Asc_CreatedAt(ctx)
	if err != nil {
		return res, ErrPinDB.Wrap(err)
	}

	for _, r := range rows {
		res = append(res, dbToPin(r))
	}
	return res, ErrPinDB.Wrap(err)
}

// CreateNode inserts node record to the table (represents a pinned resource).
func (p *PinDB) CreateNode(ctx context.Context, txHash string, ix int, cid string, days uint) (err error) {

	return pgxutil.Conn(ctx, p.db, func(conn *pgx.Conn) error {
		var batch pgx.Batch

		batch.Queue("BEGIN TRANSACTION")
		statement := p.db.Rebind(
			`INSERT INTO nodes (cid,days,pinned_at)
					VALUES (?, ?,now())
					ON CONFLICT(cid)
					DO UPDATE SET days = GREATEST(nodes.days - (now()::date - nodes.pinned_at::date),0) + ?, pinned_at = now()`,
		)

		batch.Queue(statement, cid, days, days)

		statement = p.db.Rebind(
			`UPDATE pins SET processed = true WHERE tx = ? AND ix = ?`,
		)

		batch.Queue(statement, txHash, ix)

		batch.Queue("COMMIT TRANSACTION")
		results := conn.SendBatch(ctx, &batch)

		defer func() { err = errs.Combine(err, results.Close()) }()
		var errlist errs.Group
		for i := 0; i < batch.Len(); i++ {
			_, err := results.Exec()
			errlist.Add(err)
		}

		return errlist.Err()
	})

}

func (p *PinDB) AllNodes(ctx context.Context) ([]Cid, error) {
	var c []Cid
	res, err := p.db.All_Node_OrderBy_Desc_PinnedAt(ctx)
	if err != nil {
		return c, ErrPinDB.Wrap(err)
	}
	for _, r := range res {

		c = append(c, Cid{
			Hash:      r.Cid,
			Pinned:    r.PinnedAt,
			ValidDays: uint(r.Days),
		})
	}
	return c, nil
}

func (p *PinDB) Node(ctx context.Context, hash string) (Cid, error) {
	var c Cid
	res, err := p.db.Get_Node_By_Cid(ctx, dbx.Node_Cid(hash))
	if err != nil {
		return c, ErrPinDB.Wrap(err)
	}
	c.Hash = res.Cid
	c.ValidDays = uint(res.Days)
	c.Pinned = res.PinnedAt
	return c, nil
}

func (p *PinDB) AllPins(ctx context.Context) ([]Pin, error) {
	var c []Pin
	res, err := p.db.All_Pin_OrderBy_Desc_CreatedAt(ctx)
	if err != nil {
		return c, ErrPinDB.Wrap(err)
	}
	for _, r := range res {
		c = append(c, dbToPin(r))

	}
	return c, nil
}

func (p *PinDB) PinsOfHash(ctx context.Context, cid string) ([]Pin, error) {
	var c []Pin
	res, err := p.db.All_Pin_By_Cid_OrderBy_Desc_CreatedAt(ctx, dbx.Pin_Cid(cid))
	if err != nil {
		return c, ErrPinDB.Wrap(err)
	}
	for _, r := range res {
		c = append(c, dbToPin(r))
	}
	return c, nil
}

func dbToPin(r *dbx.Pin) Pin {
	amount := new(big.Int)
	if r.Amount != "" {
		amount, _ = new(big.Int).SetString(r.Amount, 10)

	}
	return Pin{
		Cid:         r.Cid,
		Amount:      amount,
		Transaction: r.Tx,
		LogIndex:    uint(r.Ix),
		Processed:   r.Processed,
	}
}
