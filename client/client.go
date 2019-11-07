package client

import (
	"github.com/firststreet/firststreet-go/backend"
	"github.com/firststreet/firststreet-go/citysummary"
	"github.com/firststreet/firststreet-go/propertysummary"
)

// The API provides interfaces to Flood iQ services
type API struct {
	RateLimit       *backend.RateLimit
	PropertySummary *propertysummary.Client
	CitySummary     *citysummary.Client
}

// InitAPI initalizes services-level APIs
func (a *API) InitAPI(backend *backend.Backend) {
	a.PropertySummary = &propertysummary.Client{B: backend}
	a.CitySummary = &citysummary.Client{B: backend}
}

// New creates an API Client
func New(backend *backend.Backend) *API {
	a := &API{}
	a.InitAPI(backend)
	return a
}
