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
func (c Client) GetPropertyByFSID(fsid string) (*firststreet.Summary, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/property/%s?key=%s", fsid)
	summaryResponse := &firststreet.Summary{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// GetPropertyByLatLng pulls a parcel by lat lng
func (c Client) GetPropertyByLatLng(lat, lng float64) (*firststreet.Summary, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/summary/property/?key=%s&lat=%s&lng=%s", latStr, lngStr)
	summaryResponse := &firststreet.Summary{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// GetPropertyByAddress pulls a parcel by lat lng
func (c Client) GetPropertyByAddress(address string) (*firststreet.Summary, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/property/?key=%s&address=%s", address)
	summaryResponse := &firststreet.Summary{}
	err := c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// // GetCityByID retreives a City Parcel by its unique identifier
// func (c Client) GetCityByID(id string) (*City, error) {
// 	path := backend.FormatURLPath("/data/"+version+"/summary/%s?type=city&key=%s", id)
// 	city := &City{}
// 	err := c.B.Call(http.MethodGet, path, city)
// 	return city, err
// }
