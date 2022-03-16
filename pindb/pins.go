// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package pindb

import (
	"context"
	"github.com/elek/easypin/pindb/dbx"
	"math/big"
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
}

type Cid struct {
	Hash   string
	Pinned time.Time
	Expiry time.Time
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

	rows, err := p.db.All_Pin_Tx_Pin_Ix_Pin_Cid_Pin_Amount_By_Processed_Equal_False_OrderBy_Asc_CreatedAt(ctx)
	if err != nil {
		return res, ErrPinDB.Wrap(err)
	}

	var ok bool
	for _, row := range rows {
		p := Pin{}
		p.Cid = row.Cid
		p.LogIndex = uint(row.Ix)
		p.Transaction = row.Tx
		if row.Amount != "" {
			p.Amount, ok = new(big.Int).SetString(row.Amount, 10)
			if !ok {
				return res, errs.New("Couldn't parse amount %s", row.Amount)
			}
		}
		res = append(res, p)
	}
	return res, ErrPinDB.Wrap(err)
}

// CreateNode inserts node record to the table (represents a pinned resource).
func (p *PinDB) CreateNode(ctx context.Context, txHash string, ix int, cid string, expiry time.Time, amount *big.Int) (err error) {
	tx, err := p.db.Open(ctx)
	if err != nil {
		return ErrPinDB.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, tx.Commit())
	}()

	_, err = tx.Create_Node(ctx, dbx.Node_Cid(cid), dbx.Node_ExpiredAt(expiry), dbx.Node_Amount(amount.String()))
	if err != nil {
		return ErrPinDB.Wrap(err)
	}

	_, err = tx.Update_Pin_By_Tx_And_Ix(ctx, dbx.Pin_Tx(txHash), dbx.Pin_Ix(ix), dbx.Pin_Update_Fields{
		Processed: dbx.Pin_Processed(true),
	})
	if err != nil {
		return ErrPinDB.Wrap(err)
	}

	return ErrPinDB.Wrap(err)
}

func (p *PinDB) AllNodes(ctx context.Context) ([]Cid, error) {
	var c []Cid
	res, err := p.db.All_Node_OrderBy_Desc_CreatedAt(ctx)
	if err != nil {
		return c, ErrPinDB.Wrap(err)
	}
	for _, r := range res {
		c = append(c, Cid{
			Hash:   r.Cid,
			Pinned: r.CreatedAt,
			Expiry: r.ExpiredAt,
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
	c.Expiry = res.ExpiredAt
	c.Pinned = res.CreatedAt
	return c, nil
}
