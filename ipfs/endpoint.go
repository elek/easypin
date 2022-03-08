package ipfs

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"net/http"
)

// ErrEndpoint - ipfs endpoint error class.
var ErrEndpoint = errs.Class("ipfs endpoint")

// architecture: Endpoint
type Endpoint struct {
	log     *zap.Logger
	service *Service
}

// NewEndpoint creates new ipfs endpoint instance.
func NewEndpoint(log *zap.Logger, service *Service) *Endpoint {
	return &Endpoint{
		log:     log,
		service: service,
	}
}

// Register registers endpoint methods on API server subroute.
func (endpoint *Endpoint) Register(router *mux.Router) {
	//TODO: make these available only for the admins
	router.HandleFunc("/peers", endpoint.Peers).Methods(http.MethodGet)
	router.HandleFunc("/pin/{cid}", endpoint.Pin).Methods(http.MethodPost)
}

// Peers endpoint retrieves all the active Peers
func (endpoint *Endpoint) Peers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	peers, err := endpoint.service.GetPeers(ctx)
	if err != nil {
		endpoint.serveJSONError(w, http.StatusInternalServerError, ErrEndpoint.Wrap(err))
		return
	}

	err = json.NewEncoder(w).Encode(peers)
	if err != nil {
		endpoint.log.Error("failed to write json pins response", zap.Error(ErrEndpoint.Wrap(err)))
		return
	}

}

// Pin endpoint pins one specific cid.
func (endpoint *Endpoint) Pin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	cid := mux.Vars(r)["cid"]

	if cid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pinned, err := endpoint.service.Pin(ctx, cid)
	if err != nil {
		endpoint.serveJSONError(w, http.StatusInternalServerError, ErrEndpoint.Wrap(err))
		return
	}

	err = json.NewEncoder(w).Encode(pinned)
	if err != nil {
		endpoint.log.Error("failed to write json error response", zap.Error(ErrEndpoint.Wrap(err)))
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
