package client

import (
	"github.com/firststreet/firststreet-go/backend"
	"github.com/firststreet/firststreet-go/datasummary"
)

// The API provides interfaces to Flood iQ services
type API struct {
	DataSummary *datasummary.Client
	RateLimit   *backend.RateLimit
}

// InitAPI initalizes services-level APIs
func (a *API) InitAPI(backend *backend.Backend) {
	a.DataSummary = &datasummary.Client{B: backend}
}

// New creates an API Client
func New(backend *backend.Backend) *API {
	a := &API{}
	a.InitAPI(backend)
	return a
}
