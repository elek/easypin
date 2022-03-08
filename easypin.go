// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package easypin

import (
	"context"
	"github.com/elek/easypin/ipfs"
	"github.com/elek/easypin/pin"
	pindb "github.com/elek/easypin/pindb"
	"github.com/zeebo/errs"
	"net"
	"storj.io/storjscan/tokens"

	"github.com/elek/easypin/api"
	"github.com/spacemonkeygo/monkit/v3"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"storj.io/private/debug"
	"storj.io/storj/private/lifecycle"
)

var mon = monkit.Package()

// Config wraps storjscan configuration.
type Config struct {
	Debug debug.Config
	Pin   pin.Config
	API   api.Config
	IPFS  ipfs.Config
}

// DB is a collection of storjscan databases.
type DB interface {
	Pins() *pindb.PinDB
}

// App is the storjscan process that runs API endpoint.
//
// architecture: Peer
type App struct {
	Log      *zap.Logger
	DB       DB
	Servers  *lifecycle.Group
	Services *lifecycle.Group

	Debug struct {
		Listener net.Listener
		Server   *debug.Server
	}

	Pin struct {
		Service  *pin.Service
		Endpoint *pin.Endpoint
	}

	API struct {
		Listener net.Listener
		Server   *api.Server
	}

	IPFS struct {
		Service  *ipfs.Service
		Endpoint *ipfs.Endpoint
	}
}

// NewApp creates new storjscan application instance.
func NewApp(log *zap.Logger, config Config, db DB) (*App, error) {
	app := &App{
		Log:      log,
		DB:       db,
		Services: lifecycle.NewGroup(log.Named("services")),
		Servers:  lifecycle.NewGroup(log.Named("servers")),
	}

	{ // IPFS
		var err error

		if config.IPFS.Address == "" {
			return nil, errs.New("IPFS node address is empty")
		}
		app.IPFS.Service, err = ipfs.NewService(log.Named("ipfs:service"), config.IPFS.Address)
		if err != nil {
			return nil, err
		}
		app.IPFS.Endpoint = ipfs.NewEndpoint(log.Named("ipfs:endpoint"), app.IPFS.Service)
	}

	{ // pin
		token, err := tokens.AddressFromHex(config.Pin.TokenAddress)
		if err != nil {
			return nil, err
		}

		app.Pin.Service = pin.NewService(log.Named("pin:service"),
			db.Pins(),
			config.Pin.Endpoint,
			token)

		app.Pin.Endpoint = pin.NewEndpoint(log.Named("pin:endpoint"), app.Pin.Service)
	}

	{
		chore := pin.NewChore(log.Named("persister:core"), db.Pins(), config.Pin.Endpoint, config.Pin.TokenAddress)
		app.Services.Add(lifecycle.Item{
			Name:  "persister:chore",
			Run:   chore.Run,
			Close: chore.Close,
		})
	}
	{ // API
		var err error

		app.API.Listener, err = net.Listen("tcp", config.API.Address)
		if err != nil {
			return nil, err
		}

		app.API.Server = api.NewServer(log.Named("api:server"), app.API.Listener)
		app.API.Server.NewAPI("/pin", app.Pin.Endpoint.Register)
		app.API.Server.NewAPI("/ipfs", app.IPFS.Endpoint.Register)

		app.Servers.Add(lifecycle.Item{
			Name:  "api",
			Run:   app.API.Server.Run,
			Close: app.API.Server.Close,
		})
	}
	{ // wallets
		//TODO
	}

	return app, nil
}

// Run runs storjscan until it's either closed or it errors.
func (app *App) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	group, ctx := errgroup.WithContext(ctx)

	app.Servers.Run(ctx, group)
	app.Services.Run(ctx, group)

	return group.Wait()
}

// Close closes all the resources.
func (app *App) Close() error {
	return app.Services.Close()
	return app.Servers.Close()
}
