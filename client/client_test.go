package client

import (
	"net/http"
	"testing"

	"github.com/firststreet/firststreet-go/backend"
	"github.com/stretchr/testify/assert"
)

var testBackend = &backend.Backend{
	HTTPClient: &http.Client{},
	Key:        "test",
}

func TestNew(t *testing.T) {
	a := New(testBackend)
	assert.IsType(t, a, new(API))
}
