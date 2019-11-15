package marketvalueimpact

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

func marketValueImpactPropertyHandler() http.HandlerFunc {
	mviSample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/mvi-property.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// nolint
		w.Write(mviSample)
	})
}

func marketValueImpactCityHandler() http.HandlerFunc {
	mviSample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/mvi-city.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// nolint
		w.Write(mviSample)
	})
}

func TestMVIPropertyLookup(t *testing.T) {
	testutil.StartServer(marketValueImpactPropertyHandler())

	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	fsidLookup := &firststreet.Lookup{
		FSID: 450350219571,
	}

	mvi, err := c.Lookup(firststreet.PropertyLocationType, fsidLookup)
	assert.Nil(t, err)
	assert.NotNil(t, mvi)
	assert.Equal(t, mvi.FSID, int64(450350219571))
	assert.NotNil(t, mvi.Results)

	fsidLookupBad := &firststreet.Lookup{}
	_, err = c.Lookup(firststreet.PropertyLocationType, fsidLookupBad)
	assert.NotNil(t, err)

	latLngLookup := &firststreet.Lookup{
		Lat: -43.23,
		Lng: -23.12,
	}

	mvi, err = c.Lookup(firststreet.PropertyLocationType, latLngLookup)
	assert.Nil(t, err)
	assert.NotNil(t, mvi)
	assert.Equal(t, mvi.FSID, int64(450350219571))
	assert.NotNil(t, mvi.Results)

	addressLookup := &firststreet.Lookup{
		Address: "1327 Autumn Drive Fernandina Beach FL",
	}

	mvi, err = c.Lookup(firststreet.PropertyLocationType, addressLookup)
	assert.Nil(t, err)
	assert.NotNil(t, mvi)
	assert.Equal(t, mvi.FSID, int64(450350219571))
	assert.NotNil(t, mvi.Results)

}

func TestMVICityLookup(t *testing.T) {
	testutil.StartServer(marketValueImpactCityHandler())

	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	fsidLookup := &firststreet.Lookup{
		FSID: 1222175,
	}

	mvi, err := c.Lookup(firststreet.CityLocationType, fsidLookup)
	assert.Nil(t, err)
	assert.NotNil(t, mvi)
	assert.Equal(t, mvi.FSID, int64(1222175))
	assert.NotNil(t, mvi.Results)

	fsidLookupBad := &firststreet.Lookup{}
	_, err = c.Lookup(firststreet.CityLocationType, fsidLookupBad)
	assert.NotNil(t, err)

	latLngLookup := &firststreet.Lookup{
		Lat: -43.23,
		Lng: -23.12,
	}

	mvi, err = c.Lookup(firststreet.CityLocationType, latLngLookup)
	assert.Nil(t, err)
	assert.NotNil(t, mvi)
	assert.Equal(t, mvi.FSID, int64(1222175))
	assert.NotNil(t, mvi.Results)

	addressLookup := &firststreet.Lookup{
		Address: "Fernandina Beach FL",
	}

	mvi, err = c.Lookup(firststreet.CityLocationType, addressLookup)
	assert.Nil(t, err)
	assert.NotNil(t, mvi)
	assert.Equal(t, mvi.FSID, int64(1222175))
	assert.NotNil(t, mvi.Results)

}
