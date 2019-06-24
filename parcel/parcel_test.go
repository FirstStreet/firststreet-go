package parcel

import (
	"testing"
	"net/http"
	assert "github.com/stretchr/testify/require"
	"github.com/firststreet/floodiq-go/backend"
	"github.com/firststreet/floodiq-go/testutil"
)

var testBackend = &backend.Backend{
	HTTPClient: &http.Client{},
}

func TestGetPropertyByID(t *testing.T) {
	testutil.Once.Do(testutil.StartServer)
	testBackend.URL = testutil.ServerAddr

	c := &Client{
		B: testBackend,
		Key: "test",
	}
	property, err := c.GetPropertyByID("test")
	assert.Nil(t, err)
	assert.NotNil(t, property)
}
