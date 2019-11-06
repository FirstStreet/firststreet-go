package tidal

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

// GetPropertyByFSID retreives a location Tidal risk by its unique identifier
func (c Client) GetPropertyByFSID(fsid string) (*firststreet.Tidal, error) {
	path := backend.FormatURLPath("/data/"+version+"/tidal/property/%s", fsid)
	response := &firststreet.Tidal{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetPropertyByLatLng pulls a parcel's Tidal risk by lat lng
func (c Client) GetPropertyByLatLng(lat, lng float64) (*firststreet.Tidal, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/tidal/property?lat=%s&lng=%s", latStr, lngStr)
	response := &firststreet.Tidal{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetPropertyByAddress pulls a parcel's Tidal risk by address
func (c Client) GetPropertyByAddress(address string) (*firststreet.Tidal, error) {
	path := backend.FormatURLPath("/data/"+version+"/tidal/property?address=%s", address)
	response := &firststreet.Tidal{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByFSID retreives City Tidal Risk by its unique identifier
func (c Client) GetCityByFSID(fsid string) (*firststreet.Tidal, error) {
	path := backend.FormatURLPath("/data/"+version+"/tidal/city/%s", fsid)
	response := &firststreet.Tidal{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByLatLng retrieves City Tidal Risk by lat lng
func (c Client) GetCityByLatLng(lat, lng float64) (*firststreet.Tidal, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/tidal/city?lat=%s&lng=%s", latStr, lngStr)
	response := &firststreet.Tidal{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByAddress pulls City Tidal Risk by lat lng
func (c Client) GetCityByAddress(address string) (*firststreet.Tidal, error) {
	path := backend.FormatURLPath("/data/"+version+"/tidal/property?address=%s", address)

	response := &firststreet.Tidal{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

