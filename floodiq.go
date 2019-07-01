package floodiq

import (
	"net/http"

	"github.com/firststreet/floodiq-go/backend"
	"github.com/firststreet/floodiq-go/parcel"
)

const (
	// DefaultHost is where all API requests are handled
	DefaultHost = "http://apidev.firststreet.org"
	// Version is the current api version
	Version = "1.0"
)

// The API provides interfaces to Flood iQ services
type API struct {
	Parcel    *parcel.Client
	RateLimit *backend.RateLimit
}

// InitAPI initalizes services-level APIs
func (a *API) InitAPI(key string, backend *backend.Backend) {
	a.Parcel = &parcel.Client{B: backend, Key: key}
}

// newAPI creates an API for usage
func newAPI(key string, backend *backend.Backend) *API {
	a := &API{}
	a.InitAPI(key, backend)
	return a
}

func New(key string) *API {
	// The backend is used to interact with the Flood iQ API
	backend := &backend.Backend{
		HTTPClient: &http.Client{},
		URL:        DefaultHost,
		Version:    Version,
	}

	a := newAPI(key, backend)

	a.RateLimit = backend.RateLimit

	return a
}
