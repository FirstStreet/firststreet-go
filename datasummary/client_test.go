package datasummary

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

func summaryPropertyHandler() http.HandlerFunc {
	summarySample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/summary-property.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// nolint
		w.Write(summarySample)
	})
}

func summaryPropertyNullHandler() http.HandlerFunc {
	summarySample, err := ioutil.ReadFile(testutil.GetDirectoryPath() + "/fixtures/summary-property-null.json")
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(summarySample)
	})
}
func TestGetPropertyByID(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryPropertyHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}
	property, err := c.GetPropertyByFSID("100032470544")
	assert.Nil(t, err)
	assert.NotNil(t, property)
}

func TestGetPropertyByIDNull(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryPropertyNullHandler())
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
		testutil.StartServer(summaryPropertyHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	summaryResponse, err := c.GetPropertyByAddress("just a test")
	assert.Nil(t, err)
	assert.NotNil(t, summaryResponse)

	want := &firststreet.SummaryProperty{
		FSID:    summaryResponse.FSID,
		Results: summaryResponse.Results,
	}

	assert.Equal(t, want, summaryResponse)
}
func TestGetPropertyByAddressNull(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryPropertyNullHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	summaryResponse, err := c.GetPropertyByAddress("just a test")
	assert.Nil(t, err)
	assert.NotNil(t, summaryResponse)

	want := &firststreet.SummaryProperty{
		FSID:    summaryResponse.FSID,
		Results: summaryResponse.Results,
	}

	assert.Equal(t, want, summaryResponse)
}

func TestGetPropertyByLatLng(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryPropertyHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	summaryResponse, err := c.GetPropertyByLatLng(123.45, 67.8810)
	assert.Nil(t, err)
	assert.NotNil(t, summaryResponse)

	want := &firststreet.SummaryProperty{
		FSID:    summaryResponse.FSID,
		Results: summaryResponse.Results,
	}

	assert.Equal(t, want, summaryResponse)
}

func TestGetPropertyByLatLngNull(t *testing.T) {
	testutil.Once.Do(func() {
		testutil.StartServer(summaryPropertyNullHandler())
	})
	testBackend.URL = testutil.ServerAddr
	c := &Client{
		B: testBackend,
	}

	summaryResponse, err := c.GetPropertyByLatLng(123.45, 67.8810)
	assert.Nil(t, err)
	assert.NotNil(t, summaryResponse)

	want := &firststreet.SummaryProperty{
		FSID:    summaryResponse.FSID,
		Results: summaryResponse.Results,
	}

	assert.Equal(t, want, summaryResponse)
}
