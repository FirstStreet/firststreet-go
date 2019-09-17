package summary

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/firststreet/floodiq-go/backend"
	"github.com/firststreet/floodiq-go/testutil"
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

	// 	property, err = c.GetPropertyByLatLng(39.4419892114799, -75.6453718684964)
	// 	assert.Nil(t, err)
	// 	assert.NotNil(t, property)
}
