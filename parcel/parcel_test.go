package parcel

import (
	"net/http"
	"testing"

	"github.com/firststreet/floodiq-go/backend"
	"github.com/firststreet/floodiq-go/testutil"
	assert "github.com/stretchr/testify/require"
)

var testBackend = &backend.Backend{
	HTTPClient: &http.Client{},
}

func TestGetPropertyByID(t *testing.T) {
	testutil.Once.Do(testutil.StartServer)
	testBackend.URL = testutil.ServerAddr

	c := &Client{
		B:   testBackend,
		Key: "test",
	}
	property, err := c.GetPropertyByID("test")
	assert.Nil(t, err)
	assert.NotNil(t, property)

	property, err = c.GetPropertyByLatLng(39.4419892114799, -75.6453718684964)
	assert.Nil(t, err)
	assert.NotNil(t, property)
}

func TestAddressLookup(t *testing.T) {
	testutil.Once.Do(testutil.StartServer)
	testBackend.URL = testutil.ServerAddr

	c := &Client{
		B:   testBackend,
		Key: "test",
	}
	property, err := c.GetPropertyByAddress("212 appoquin s, middletown, delware")
	assert.Nil(t, err)
	assert.NotNil(t, property)

	city, err := c.GetCityByAddress("Miami, FL")
	assert.Nil(t, err)
	assert.NotNil(t, city)
}
