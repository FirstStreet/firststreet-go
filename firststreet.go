package firststreet

import (
	"net/http"

	"github.com/firststreet/firststreet-go/backend"
)

var Key string

const (
	// DefaultHost is where all API requests are handled
	DefaultHost = "https://api.firststreet.org"
	// Version is the current api version
	Version = "v0.1"
)

// NewBackend returns an API backend for a `client` to use
func NewBackend(key string) *backend.Backend {
	// The backend is used to interact with the First Street API
	backend := &backend.Backend{
		HTTPClient: &http.Client{},
		Key:        key,
		URL:        DefaultHost,
		Version:    Version,
	}

	return backend
}
