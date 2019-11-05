package tidal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	firststreet "github.com/firststreet/firststreet-go"
	"github.com/firststreet/firststreet-go/backend"
	"github.com/firststreet/firststreet-go/testutil"
	assert "github.com/stretchr/testify/require"
)

var testBackend = &backend.Backend{
	HTTPClient: &http.Client{},
	Key:        "test",
}

func tidalPropertyHandler() http.HandlerFunc {
	tidalSample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/tidal-property.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// nolint
		w.Write(tidalSample)
	})
}

func tidalCityHandler() http.HandlerFunc {
	tidalSample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/tidal-city.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// nolint
		w.Write(tidalSample)
	})
}

func TestGetPropertyByID(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalPropertyHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}
	property, err := c.GetPropertyByFSID("100032470544")
	assert.Nil(t, err)
	assert.NotNil(t, property)
}

func TestGetPropertyByAddress(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalPropertyHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	response, err := c.GetPropertyByAddress("just a test")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	want := &firststreet.Tidal{
		FSID:    response.FSID,
		Results: response.Results,
	}

	assert.Equal(t, want, response)
}

func TestGetPropertyByLatLng(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalPropertyHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	response, err := c.GetPropertyByLatLng(123.45, 67.8810)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	want := &firststreet.Tidal{
		FSID:    response.FSID,
		Results: response.Results,
	}

	assert.Equal(t, want, response)
}

func TestGetCityByID(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalCityHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}
	property, err := c.GetCityByFSID("100032470544")
	assert.Nil(t, err)
	assert.NotNil(t, property)
}

func TestGetCityByAddress(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalCityHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	response, err := c.GetCityByAddress("just a test")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	want := &firststreet.Tidal{
		FSID:    response.FSID,
		Results: response.Results,
	}

	assert.Equal(t, want, response)
}

func TestGetCityByLatLng(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalCityHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	response, err := c.GetCityByLatLng(123.45, 67.8810)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	want := &firststreet.Tidal{
		FSID:    response.FSID,
		Results: response.Results,
	}

	assert.Equal(t, want, response)
}
