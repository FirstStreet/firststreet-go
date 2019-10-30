package marketvalueimpact

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

// GetPropertyByFSID retreives a location's risk MVI by its unique identifier
func (c Client) GetPropertyByFSID(fsid string) (*firststreet.MVI, error) {
	path := backend.FormatURLPath("/data/"+version+"/market-value-impact/property/%s", fsid)
	response := &firststreet.MVI{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetPropertyByLatLng pulls a parcel's risk MVI by lat lng
func (c Client) GetPropertyByLatLng(lat, lng float64) (*firststreet.MVI, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/market-value-impact/property?lat=%s&lng=%s", latStr, lngStr)
	response := &firststreet.MVI{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetPropertyByAddress pulls a parcel's risk MVI by address
func (c Client) GetPropertyByAddress(address string) (*firststreet.MVI, error) {
	path := backend.FormatURLPath("/data/"+version+"/market-value-impact/property?address=%s", address)
	response := &firststreet.MVI{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByFSID retreives City Risk MVI by its unique identifier
func (c Client) GetCityByFSID(fsid string) (*firststreet.MVI, error) {
	path := backend.FormatURLPath("/data/"+version+"/market-value-impact/city/%s", fsid)
	response := &firststreet.MVI{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByLatLng retrieves City Risk MVI by lat lng
func (c Client) GetCityByLatLng(lat, lng float64) (*firststreet.MVI, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/market-value-impact/city?lat=%s&lng=%s", latStr, lngStr)
	response := &firststreet.MVI{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByAddress pulls City Risk MVI by lat lng
func (c Client) GetCityByAddress(address string) (*firststreet.MVI, error) {
	path := backend.FormatURLPath("/data/"+version+"/market-value-impact/property?address=%s", address)

	response := &firststreet.MVI{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

