// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package easypin

import (
	"context"
	"github.com/elek/easypin/pin"
	pindb "github.com/elek/easypin/pindb"
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
}

// DB is a collection of storjscan databases.
type DB interface {
	Pins() *pindb.PinDB
}

// App is the storjscan process that runs API endpoint.
//
// architecture: Peer
type App struct {
	Log     *zap.Logger
	DB      DB
	Servers *lifecycle.Group

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
}

// NewApp creates new storjscan application instance.
func NewApp(log *zap.Logger, config Config, db DB) (*App, error) {
	app := &App{
		Log: log,
		DB:  db,

		Servers: lifecycle.NewGroup(log.Named("servers")),
	}

	{ // pin
		token, err := tokens.AddressFromHex(config.Pin.TokenAddress)
		if err != nil {
			return nil, err
		}

		app.Pin.Service = pin.NewService(log.Named("pin:service"),
			config.Pin.Endpoint,
			token)

		app.Pin.Endpoint = pin.NewEndpoint(log.Named("pin:endpoint"), app.Pin.Service)
	}

	{ // API
		var err error

		app.API.Listener, err = net.Listen("tcp", config.API.Address)
		if err != nil {
			return nil, err
		}

		app.API.Server = api.NewServer(log.Named("api:server"), app.API.Listener)
		app.API.Server.NewAPI("/pin", app.Pin.Endpoint.Register)

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

	return group.Wait()
}

// Close closes all the resources.
func (app *App) Close() error {
	return app.Servers.Close()
}
