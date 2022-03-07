package pin

import (
	"context"
	"github.com/elek/easypin/pin/contract"
	"github.com/elek/easypin/pindb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	errs "github.com/zeebo/errs/v2"
	"go.uber.org/zap"
	"storj.io/common/sync2"
	"time"
)

// Chore persists new on-chain events to the db
type Chore struct {
	db       *pindb.PinDB
	log      *zap.Logger
	endpoint string
	contract string
	Loop     *sync2.Cycle
}

func NewChore(log *zap.Logger, db *pindb.PinDB, endpoint string, contract string) *Chore {
	return &Chore{
		log:      log,
		db:       db,
		endpoint: endpoint,
		contract: contract,
		Loop:     sync2.NewCycle(1 * time.Minute),
	}
}

func (c *Chore) Run(ctx context.Context) error {
	return c.Loop.Run(ctx, func(ctx context.Context) error {
		c.log.Info("Refreshing database with onchain data")
		err := c.PersistRequests(ctx)
		if err != nil {
			return err
		}
		err = c.PinMissing(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (c *Chore) Close() error {
	c.Loop.Close()
	return nil
}

// PersistRequests reads all the pin requests and save them to the db.
func (c *Chore) PersistRequests(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	client, err := ethclient.DialContext(ctx, c.endpoint)
	if err != nil {
		return ErrService.Wrap(err)
	}
	defer client.Close()

	pin, err := contract.NewStorjPin(common.HexToAddress(c.contract), client)
	if err != nil {
		return ErrService.Wrap(err)
	}

	opts := &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	}
	iter, err := pin.FilterPinned(opts, nil)
	if err != nil {
		return ErrService.Wrap(err)
	}
	defer func() { err = errs.Combine(err, ErrService.Wrap(iter.Close())) }()

	for iter.Next() {
		err = c.db.Create(ctx, iter.Event.Raw.TxHash.Hex(),
			iter.Event.Raw.Index,
			iter.Event.Hash,
			iter.Event.Amount.Int64())
		if err != nil {
			return err
		}
	}

	return nil
}

// PinMissing pins the requests which are not yet pinned.
func (c *Chore) PinMissing(ctx context.Context) (err error) {
	pins, err := c.db.FindNew(ctx)
	if err != nil {
		return err
	}
	for _, p := range pins {
		c.log.Info("Oh, let me pin", zap.String("cid", p.Cid))
		//TODO calculate the exact period
		until := time.Now().Add(7 * time.Hour * 24)

		//TODO pin the cid

		err := c.db.CreateNode(ctx, p.Cid, until, p.Amount)
		if err != nil {
			c.log.Error("Couldn't pin the CID", zap.String("cid", p.Cid), zap.Error(err))
		}
		c.log.Error("IPFS Cid is pinned", zap.String("cid", p.Cid), zap.Time("until", until))
	}
	return nil
}
