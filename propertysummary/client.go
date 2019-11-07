package propertysummary

import (
	"errors"
	"net/http"

	firststreet "github.com/firststreet/firststreet-go"
	"github.com/firststreet/firststreet-go/backend"
)

// Client is used for /data/{svc}
type Client struct {
	B *backend.Backend
}

func (c Client) Lookup(lookup *firststreet.Lookup) (*firststreet.PropertySummary, error) {
	var path string
	var err error

	summaryResponse := &firststreet.PropertySummary{}

	lookupType, err := lookup.LookupType()

	if err != nil {
		return summaryResponse, err
	}

	if lookupType == "" {
		return summaryResponse, errors.New("Lookup type could not be determined: Expecting an FSID, LatLng, or Address.")
	}

	if lookupType == "fsid" {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/summary/property/%s", lookup.FSIDString())
	}

	if lookupType == "coordinate" {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/summary/property?lat=%s&lng=%s", lookup.LatLngString("lat"), lookup.LatLngString("lng"))
	}

	if lookupType == "address" {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/summary/property?address=%s", lookup.Address)
	}

	err = c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}

// // GetCityByFSID retreives a City Parcel by its unique identifier
// func (c Client) GetCityByFSID(fsid string) (*firststreet.SummaryCity, error) {
// 	path := backend.FormatURLPath("/data/"+version+"/summary/city/%s", fsid)
// 	summaryResponse := &firststreet.SummaryCity{}
// 	err := c.B.Call(http.MethodGet, path, summaryResponse)
// 	return summaryResponse, err
// }

// // GetCityByLatLng retrieves a City Parcel by lat lng
// func (c Client) GetCityByLatLng(lat, lng float64) (*firststreet.SummaryCity, error) {
// 	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
// 	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

// 	path := backend.FormatURLPath("/data/"+version+"/summary/city?lat=%s&lng=%s", latStr, lngStr)
// 	summaryResponse := &firststreet.SummaryCity{}
// 	err := c.B.Call(http.MethodGet, path, summaryResponse)
// 	return summaryResponse, err
// }

// // GetCityByAddress pulls a parcel by lat lng
// func (c Client) GetCityByAddress(address string) (*firststreet.SummaryCity, error) {
// 	path := backend.FormatURLPath("/data/"+version+"/summary/property?address=%s", address)

// 	summaryResponse := &firststreet.SummaryCity{}
// 	err := c.B.Call(http.MethodGet, path, summaryResponse)
// 	return summaryResponse, err
// }
