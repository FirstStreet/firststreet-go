package datasummary

import (
	"net/http"
	"strconv"

	firststreet "github.com/firststreet/firststreet-go"
	"github.com/firststreet/firststreet-go/backend"
)

// version is the current product version
const version = "v1.0"

// Client is used for /data/{svc}
type Client struct {
	B *backend.Backend
}

// GetPropertyByFSID retreives a Location Summary by its unique identifier
func (c Client) GetPropertyByFSID(fsid string) (*firststreet.SummaryProperty, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/property/%s", fsid)
	summaryResponse := &firststreet.SummaryProperty{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// GetPropertyByLatLng pulls a parcel by lat lng
func (c Client) GetPropertyByLatLng(lat, lng float64) (*firststreet.SummaryProperty, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/summary/property?lat=%s&lng=%s", latStr, lngStr)
	summaryResponse := &firststreet.SummaryProperty{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// GetPropertyByAddress pulls a parcel by lat lng
func (c Client) GetPropertyByAddress(address string) (*firststreet.SummaryProperty, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/property?address=%s", address)
	summaryResponse := &firststreet.SummaryProperty{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// GetCityByFSID retreives a City Parcel by its unique identifier
func (c Client) GetCityByFSID(fsid string) (*firststreet.SummaryCity, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/city/%s", fsid)
	summaryResponse := &firststreet.SummaryCity{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// GetCityByLatLng retrieves a City Parcel by lat lng
func (c Client) GetCityByLatLng(lat, lng float64) (*firststreet.SummaryCity, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/summary/city?lat=%s&lng=%s", latStr, lngStr)
	summaryResponse := &firststreet.SummaryCity{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// GetCityByAddress pulls a parcel by lat lng
func (c Client) GetCityByAddress(address string) (*firststreet.SummaryCity, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/property?address=%s", address)
	summaryResponse := &firststreet.SummaryCity{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

