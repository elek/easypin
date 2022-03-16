// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package pin

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
)

// ErrEndpoint - pin endpoint error class.
var ErrEndpoint = errs.Class("pin endpoint")

// Endpoint for querying ERC20 token information from ethereum chain.
//
// architecture: Endpoint
type Endpoint struct {
	log     *zap.Logger
	service *Service
}

// NewEndpoint creates new payments endpoint instance.
func NewEndpoint(log *zap.Logger, service *Service) *Endpoint {
	return &Endpoint{
		log:     log,
		service: service,
	}
}

// Register registers endpoint methods on API server subroute.
func (endpoint *Endpoint) Register(router *mux.Router) {
	router.HandleFunc("/pin/all", endpoint.Pins).Methods(http.MethodGet)
	router.HandleFunc("/pin/config", endpoint.Config).Methods(http.MethodGet)
	router.HandleFunc("/pin/{cid}", endpoint.PinOfHash).Methods(http.MethodGet)
	router.HandleFunc("/block/all", endpoint.Blocks).Methods(http.MethodGet)
	router.HandleFunc("/block/{cid}", endpoint.Block).Methods(http.MethodGet)
}

// Pins endpoint retrieves all Pin request events.
func (endpoint *Endpoint) Pins(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	payments, err := endpoint.service.AllPins(ctx)
	if err != nil {
		endpoint.serveJSONError(w, http.StatusInternalServerError, ErrEndpoint.Wrap(err))
		return
	}

	err = json.NewEncoder(w).Encode(payments)
	if err != nil {
		endpoint.log.Error("failed to write json pins response", zap.Error(ErrEndpoint.Wrap(err)))
		return
	}
}

// Blocks endpoint retrieves all Blocks which are pinned.
func (endpoint *Endpoint) Blocks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	payments, err := endpoint.service.Cids(ctx)
	if err != nil {
		endpoint.serveJSONError(w, http.StatusInternalServerError, ErrEndpoint.Wrap(err))
		return
	}

	err = json.NewEncoder(w).Encode(payments)
	if err != nil {
		endpoint.log.Error("failed to write json pins response", zap.Error(ErrEndpoint.Wrap(err)))
		return
	}
}

// Block endpoint retrieves one Block which is pinned.
func (endpoint *Endpoint) Block(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	cid := mux.Vars(r)["cid"]
	payments, err := endpoint.service.Cid(ctx, cid)
	if errors.Is(err, sql.ErrNoRows) {
		endpoint.serveJSONError(w, http.StatusNotFound, ErrEndpoint.Wrap(err))
		return
	}
	if err != nil {
		endpoint.serveJSONError(w, http.StatusInternalServerError, ErrEndpoint.Wrap(err))
		return
	}

	err = json.NewEncoder(w).Encode(payments)
	if err != nil {
		endpoint.log.Error("failed to write json pins response", zap.Error(ErrEndpoint.Wrap(err)))
		return
	}
}

// Config returns with the actual config required by the web UI.
func (endpoint *Endpoint) Config(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	config, err := endpoint.service.Config(ctx)
	if err != nil {
		endpoint.serveJSONError(w, http.StatusInternalServerError, ErrEndpoint.Wrap(err))
		return
	}

	err = json.NewEncoder(w).Encode(config)
	if err != nil {
		endpoint.log.Error("failed to write json config response", zap.Error(ErrEndpoint.Wrap(err)))
		return
	}
}

// serveJSONError writes JSON error to response output stream.
func (endpoint *Endpoint) serveJSONError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		endpoint.log.Error("failed to write json error response", zap.Error(ErrEndpoint.Wrap(err)))
		return
	}
}

func (endpoint *Endpoint) PinOfHash(w http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	cid := mux.Vars(request)["cid"]
	payments, err := endpoint.service.PinsOfHash(ctx, cid)
	if errors.Is(err, sql.ErrNoRows) {
		endpoint.serveJSONError(w, http.StatusNotFound, ErrEndpoint.Wrap(err))
		return
	}
	if err != nil {
		endpoint.serveJSONError(w, http.StatusInternalServerError, ErrEndpoint.Wrap(err))
		return
	}

	err = json.NewEncoder(w).Encode(payments)
	if err != nil {
		endpoint.log.Error("failed to write json pins response", zap.Error(ErrEndpoint.Wrap(err)))
		return
	}
}
