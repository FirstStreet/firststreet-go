package summary

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/firststreet/firststreet-go/backend"
	"github.com/firststreet/firststreet-go/testutil"
	assert "github.com/stretchr/testify/require"
)

var testBackend = &backend.Backend{
	HTTPClient: &http.Client{},
}

func summaryHandler() http.HandlerFunc {
	summarySample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/summary.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(summarySample)
	})
}

func TestGetPropertyByID(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B:   testBackend,
		Key: "test",
	}
	property, err := c.GetPropertyByFSID("100032470544")
	assert.Nil(t, err)
	assert.NotNil(t, property)
}

func TestGetPropertyByAddress(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B:   testBackend,
		Key: "test",
	}

	summaryResponse, err := c.GetPropertyByAddress("just a test")
	assert.Nil(t, err)
	assert.NotNil(t, summaryResponse)

	want := &Summary{
		FSID:    summaryResponse.FSID,
		Results: summaryResponse.Results,
	}

	assert.Equal(t, want, summaryResponse)
}

func TestGetPropertyByLatLng(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B:   testBackend,
		Key: "test",
	}

	summaryResponse, err := c.GetPropertyByLatLng(123.45, 67.8810)
	assert.Nil(t, err)
	assert.NotNil(t, summaryResponse)

	want := &Summary{
		FSID:    summaryResponse.FSID,
		Results: summaryResponse.Results,
	}

	assert.Equal(t, want, summaryResponse)
}
