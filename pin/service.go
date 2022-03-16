// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package pin

import (
	"context"
	"github.com/elek/easypin/pin/contract"
	pindb "github.com/elek/easypin/pindb"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
)

// ErrService - pin service error class.
var ErrService = errs.Class("pin service")

// Config holds pin service configuration.
type Config struct {
	EthereumEndpoint string
	TokenContract    string
	PinContract      string
}

// Service for querying ERC20 token information from ethereum chain.
//
// architecture: Service
type Service struct {
	db            *pindb.PinDB
	log           *zap.Logger
	endpoint      string
	pinContract   Address
	tokenContract Address
}

// NewService creates new token service instance.
func NewService(log *zap.Logger, db *pindb.PinDB, endpoint string, token Address, pin Address) *Service {
	return &Service{
		db:            db,
		log:           log,
		endpoint:      endpoint,
		pinContract:   pin,
		tokenContract: token,
	}
}

// PinsFromChain returns with all on-chain pin request
func (service *Service) PinsFromChain(ctx context.Context) (_ []pindb.Pin, err error) {
	defer mon.Task()(&ctx)(&err)

	client, err := ethclient.DialContext(ctx, service.endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	pin, err := contract.NewStorjPin(service.pinContract, client)
	if err != nil {
		return nil, ErrService.Wrap(err)
	}

	opts := &bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	}
	iter, err := pin.FilterPinned(opts, nil)
	if err != nil {
		return nil, ErrService.Wrap(err)
	}
	defer func() { err = errs.Combine(err, ErrService.Wrap(iter.Close())) }()

	var pins []pindb.Pin
	for iter.Next() {
		pins = append(pins, pindb.Pin{
			Cid:         iter.Event.Hash,
			Amount:      iter.Event.Amount,
			Transaction: iter.Event.Raw.TxHash.Hex(),
			LogIndex:    iter.Event.Raw.Index,
		})
	}

	return pins, nil
}

type WebConfig struct {
	PinContract   string
	TokenContract string
}

// Config returns the UI configuration
func (service *Service) Config(ctx context.Context) (cfg WebConfig, err error) {
	defer mon.Task()(&ctx)(&err)

	return WebConfig{
		TokenContract: service.tokenContract.Hex(),
		PinContract:   service.pinContract.Hex(),
	}, nil
}

func (service *Service) Cids(ctx context.Context) (result []pindb.Cid, err error) {
	defer mon.Task()(&ctx)(&err)
	return service.db.AllNodes(ctx)
}

func (service *Service) Cid(ctx context.Context, hash string) (result pindb.Cid, err error) {
	defer mon.Task()(&ctx)(&err)
	return service.db.Node(ctx, hash)
}

func (service *Service) AllPins(ctx context.Context) (result []pindb.Pin, err error) {
	defer mon.Task()(&ctx)(&err)
	return service.db.AllPins(ctx)

}

func (service *Service) PinsOfHash(ctx context.Context, cid string) (result []pindb.Pin, err error) {
	defer mon.Task()(&ctx)(&err)
	return service.db.PinsOfHash(ctx, cid)

}
