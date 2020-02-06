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
func TestTidalLookupProperty(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalPropertyHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	fsidLookup := &firststreet.Lookup{
		FSID: "450350219571",
	}

	frd, err := c.Lookup(firststreet.PropertyLocationType, fsidLookup)
	assert.Nil(t, err)
	assert.NotNil(t, frd)
	assert.Equal(t, frd.FSID, "450350219571")
	assert.NotNil(t, frd.Results)

	fsidLookupBad := &firststreet.Lookup{}
	_, err = c.Lookup(firststreet.PropertyLocationType, fsidLookupBad)
	assert.NotNil(t, err)

	latLngLookup := &firststreet.Lookup{
		Lat: -43.23,
		Lng: -23.12,
	}

	frd, err = c.Lookup(firststreet.PropertyLocationType, latLngLookup)
	assert.Nil(t, err)
	assert.NotNil(t, frd)
	assert.Equal(t, frd.FSID, "450350219571")
	assert.NotNil(t, frd.Results)

	addressLookup := &firststreet.Lookup{
		Address: "1327 Autumn Drive Fernandina Beach FL",
	}

	frd, err = c.Lookup(firststreet.PropertyLocationType, addressLookup)
	assert.Nil(t, err)
	assert.NotNil(t, frd)
	assert.Equal(t, frd.FSID, "450350219571")
	assert.NotNil(t, frd.Results)

}

func TestTidalLookupCity(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(tidalCityHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	fsidLookup := &firststreet.Lookup{
		FSID: "450350219571",
	}

	frd, err := c.Lookup(firststreet.CityLocationType, fsidLookup)
	assert.Nil(t, err)
	assert.NotNil(t, frd)
	assert.Equal(t, frd.FSID, "450350219571")
	assert.NotNil(t, frd.Results)

	fsidLookupBad := &firststreet.Lookup{}
	_, err = c.Lookup(firststreet.CityLocationType, fsidLookupBad)
	assert.NotNil(t, err)

	latLngLookup := &firststreet.Lookup{
		Lat: -43.23,
		Lng: -23.12,
	}

	frd, err = c.Lookup(firststreet.CityLocationType, latLngLookup)
	assert.Nil(t, err)
	assert.NotNil(t, frd)
	assert.Equal(t, frd.FSID, "450350219571")
	assert.NotNil(t, frd.Results)

	addressLookup := &firststreet.Lookup{
		Address: "1327 Autumn Drive Fernandina Beach FL",
	}

	frd, err = c.Lookup(firststreet.CityLocationType, addressLookup)
	assert.Nil(t, err)
	assert.NotNil(t, frd)
	assert.Equal(t, frd.FSID, "450350219571")
	assert.NotNil(t, frd.Results)
}
