package pin

import (
	"context"
	"github.com/elek/easypin/ipfs"
	"github.com/elek/easypin/pin/contract"
	"github.com/elek/easypin/pindb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	errs "github.com/zeebo/errs/v2"
	"go.uber.org/zap"
	"math/big"
	"storj.io/common/sync2"
	"time"
)

// Chore persists new on-chain events to the db
type Chore struct {
	db           *pindb.PinDB
	log          *zap.Logger
	endpoint     string
	contract     string
	Loop         *sync2.Cycle
	IPFS         *ipfs.Service
	ByteDayPrice *big.Int //price in raw token(1e-18) / unit(=1byte) / month
}

func NewChore(log *zap.Logger, db *pindb.PinDB, service *ipfs.Service, endpoint string, contract string) *Chore {
	return &Chore{
		log:          log,
		db:           db,
		endpoint:     endpoint,
		contract:     contract,
		Loop:         sync2.NewCycle(1 * time.Minute),
		IPFS:         service,
		ByteDayPrice: big.NewInt(133333), // 4 USD / 1 TB (1E12) * token unit (1E18) / 1 USD/STORJ / 30 days
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
			iter.Event.Amount)
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
		c.log.Info("Pinning new IPFS entry", zap.String("cid", p.Cid))

		err := c.Pin(ctx, p.Transaction, p.LogIndex, p.Cid, p.Amount)
		if err != nil {
			c.log.Error("Pinning is failed", zap.String("cid", p.Cid), zap.Error(err))
		}
	}
	return nil
}

func (c *Chore) Pin(ctx context.Context, txHash string, ix uint, cid string, amount *big.Int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	pinned, err := c.IPFS.Pin(ctx, cid)
	if err != nil {
		return err
	}

	until := calculateUntil(time.Now(), c.ByteDayPrice, amount, pinned.Size)

	//TODO: fix amount
	err = c.db.CreateNode(ctx, txHash, int(ix), cid, until, amount)
	if err != nil {
		return err
	}
	c.log.Error("IPFS Block is pinned", zap.String("cid", cid), zap.Time("until", until))
	return nil
}

func calculateUntil(from time.Time, basePrice *big.Int, paidToken *big.Int, size uint64) time.Time {
	pricePerDay := new(big.Int).Mul(basePrice, big.NewInt(int64(size)))
	paidSeconds := new(big.Int).Div(new(big.Int).Mul(paidToken, big.NewInt(24*60*60)), pricePerDay)
	return from.Add(time.Duration(paidSeconds.Int64()) * time.Second)
}
