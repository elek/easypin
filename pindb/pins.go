// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package storjscandb

import (
	"context"
	"github.com/elek/easypin/pindb/dbx"

	"github.com/zeebo/errs"
)

// ErrWalletsDB indicates about internal wallets DB error.
var ErrPinDB = errs.Class("PinDB")

// PinDB is database stores the pinning events.
//
// architecture: Database
type PinDB struct {
	db *dbx.DB
}

// Pin represents an entry in the pin table.
type Pin struct {
	Cid    []byte
	Amount int64
}

// Create inserts a new entry in the wallets table.
func (p *PinDB) Create(ctx context.Context, cid []byte, amount int64) (Pin, error) {
	w, err := p.db.Create_Pin(ctx, dbx.Pin_Cid(cid), dbx.Pin_Amount(amount))
	return Pin{Cid: w.Cid, Amount: w.Amount}, ErrPinDB.Wrap(err)
}
