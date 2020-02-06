package citysummary

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

func summaryCityHandler() http.HandlerFunc {
	summarySample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/summary-city.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// nolint
		w.Write(summarySample)
	})
}

func TestCitySummary(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryCityHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	fsidLookup := &firststreet.Lookup{
		FSID: "1222175",
	}

	property, err := c.Lookup(fsidLookup)
	assert.Nil(t, err)
	assert.NotNil(t, property)
	assert.Equal(t, property.FSID, "1222175")
	assert.NotNil(t, property.Results.Location)

	fsidLookupBad := &firststreet.Lookup{}
	_, err = c.Lookup(fsidLookupBad)
	assert.NotNil(t, err)

	latLngLookup := &firststreet.Lookup{
		Lat: -43.23,
		Lng: -23.12,
	}

	property, err = c.Lookup(latLngLookup)
	assert.Nil(t, err)
	assert.NotNil(t, property)
	assert.Equal(t, property.FSID, "1222175")
	assert.NotNil(t, property.Results.Location)

	addressLookup := &firststreet.Lookup{
		Address: "1327 Autumn Drive Fernandina Beach FL",
	}

	property, err = c.Lookup(addressLookup)
	assert.Nil(t, err)
	assert.NotNil(t, property)
	assert.Equal(t, property.FSID, "1222175")
	assert.NotNil(t, property.Results.Location)

}
