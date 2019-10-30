package hurricane

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

// GetPropertyByFSID retreives a location hurricane risk by its unique identifier
func (c Client) GetPropertyByFSID(fsid string) (*firststreet.Hurricane, error) {
	path := backend.FormatURLPath("/data/"+version+"/hurricane/property/%s", fsid)
	response := &firststreet.Hurricane{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetPropertyByLatLng pulls a parcel's hurricane risk by lat lng
func (c Client) GetPropertyByLatLng(lat, lng float64) (*firststreet.Hurricane, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/hurricane/property?lat=%s&lng=%s", latStr, lngStr)
	response := &firststreet.Hurricane{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetPropertyByAddress pulls a parcel's hurricane risk by address
func (c Client) GetPropertyByAddress(address string) (*firststreet.Hurricane, error) {
	path := backend.FormatURLPath("/data/"+version+"/hurricane/property?address=%s", address)
	response := &firststreet.Hurricane{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByFSID retreives City Hurricane Risk by its unique identifier
func (c Client) GetCityByFSID(fsid string) (*firststreet.Hurricane, error) {
	path := backend.FormatURLPath("/data/"+version+"/hurricane/city/%s", fsid)
	response := &firststreet.Hurricane{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByLatLng retrieves City Hurricane Risk by lat lng
func (c Client) GetCityByLatLng(lat, lng float64) (*firststreet.Hurricane, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/hurricane/city?lat=%s&lng=%s", latStr, lngStr)
	response := &firststreet.Hurricane{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

// GetCityByAddress pulls City Hurricane Risk by lat lng
func (c Client) GetCityByAddress(address string) (*firststreet.Hurricane, error) {
	path := backend.FormatURLPath("/data/"+version+"/hurricane/property?address=%s", address)

	response := &firststreet.Hurricane{}
	err := c.B.Call(http.MethodGet, path, response)
	return response, err
}

