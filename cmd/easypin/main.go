// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"github.com/elek/easypin"
	pindb "github.com/elek/easypin/pindb"
	"github.com/spf13/pflag"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"log"

	"storj.io/private/cfgstruct"
)

// Flags contains storjscan app configuration.
var Flags struct {
	Database string
	easypin.Config
}

func init() {
	cfgstruct.Bind(pflag.CommandLine, &Flags)
}

func main() {
	pflag.Parse()

	if err := run(context.Background(), Flags.Config); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, config easypin.Config) error {
	logger := zap.NewExample()
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Println(err)
		}
	}()

	db, err := pindb.Open(ctx, logger.Named("easypin"), Flags.Database)
	if err != nil {
		return err
	}

	app, err := easypin.NewApp(logger.Named("easypin"), config, db)
	if err != nil {
		return err
	}

	runErr := app.Run(ctx)
	closeErr := app.Close()
	return errs.Combine(runErr, closeErr)
}
