package tidal

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

// Lookup provides tidal lookup
func (c Client) Lookup(locationType firststreet.LocationType, lookup *firststreet.Lookup) (*firststreet.FloodRiskData, error) {
	var path string
	var err error

	tidalResponse := &firststreet.FloodRiskData{}

	lookupType, err := lookup.LookupType()

	if err != nil {
		return tidalResponse, err
	}

	if lookupType == "" {
		return tidalResponse, errors.New("Lookup type could not be determined: Expecting an FSID, LatLng, or Address.")
	}

	if lookupType == firststreet.FSIDLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/tidal/"+string(locationType)+"/%s", lookup.FSIDString())
	}

	if lookupType == firststreet.CoordinateLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/tidal/"+string(locationType)+"?lat=%s&lng=%s", lookup.LatLngString("lat"), lookup.LatLngString("lng"))
	}

	if lookupType == firststreet.AddressLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/tidal/"+string(locationType)+"?address=%s", lookup.Address)
	}

	err = c.B.Call(http.MethodGet, path, tidalResponse)
	return tidalResponse, err
}
