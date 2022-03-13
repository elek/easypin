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
func (p *PinDB) Create(ctx context.Context, tx string, ix uint, cid string, amount int64) error {
	_, err := p.db.ExecContext(ctx, p.db.Rebind("INSERT INTO pins ( tx, ix, cid, amount ) VALUES ( ?, ?, ?, ?) ON CONFLICT DO NOTHING;"), tx, ix, cid, amount)
	if err != nil {
		return ErrPinDB.Wrap(err)
	}
	return ErrPinDB.Wrap(err)
}

func (p *PinDB) FindNew(ctx context.Context) ([]Pin, error) {
	var res []Pin

	rows, err := p.db.Query(ctx, "select paid.cid,paid.amount from (select pins.cid,sum(pins.amount) as amount from pins group by cid) paid left join nodes on nodes.cid=paid.cid WHERE nodes.cid is null;")
	if err != nil {
		return res, ErrPinDB.Wrap(err)
	}

	for rows.Next() {
		p := Pin{}
		err := rows.Scan(&p.Cid, &p.Amount)
		if err != nil {
			return res, err
		}
		res = append(res, p)
	}
	return res, ErrPinDB.Wrap(err)
}

// CreateNode inserts node record to the table (represents a pinned resource).
func (p *PinDB) CreateNode(ctx context.Context, cid string, expiry time.Time, amount int64) error {
	_, err := p.db.Create_Node(ctx, dbx.Node_Cid(cid), dbx.Node_ExpiredAt(expiry), dbx.Node_Amount(amount))
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
