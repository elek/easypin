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
	Endpoint     string
	TokenAddress string
}

// Service for querying ERC20 token information from ethereum chain.
//
// architecture: Service
type Service struct {
	db       *pindb.PinDB
	log      *zap.Logger
	endpoint string
	token    Address
}

// NewService creates new token service instance.
func NewService(log *zap.Logger, db *pindb.PinDB, endpoint string, token Address) *Service {
	return &Service{
		db:       db,
		log:      log,
		endpoint: endpoint,
		token:    token,
	}
}

// Pins returns with all on-chain pin request
func (service *Service) Pins(ctx context.Context) (_ []Pin, err error) {
	defer mon.Task()(&ctx)(&err)

	client, err := ethclient.DialContext(ctx, service.endpoint)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	pin, err := contract.NewStorjPin(service.token, client)
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

	var pins []Pin
	for iter.Next() {
		pins = append(pins, Pin{
			Cid:         iter.Event.Hash,
			TokenValue:  iter.Event.Amount,
			Transaction: iter.Event.Raw.TxHash,
		})
	}

	return pins, nil
}

type WebConfig struct {
	TokenAddress string
}

// Config returns the UI configuration
func (service *Service) Config(ctx context.Context) (cfg WebConfig, err error) {
	defer mon.Task()(&ctx)(&err)

	return WebConfig{
		TokenAddress: service.token.Hex(),
	}, nil
}
