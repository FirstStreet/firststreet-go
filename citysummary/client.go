package citysummary

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

func (c Client) Lookup(lookup *firststreet.Lookup) (*firststreet.CitySummary, error) {
	var path string
	var err error

	summaryResponse := &firststreet.CitySummary{}

	lookupType, err := lookup.LookupType()

	if err != nil {
		return summaryResponse, err
	}

	if lookupType == "" {
		return summaryResponse, errors.New("Lookup type could not be determined: Expecting an FSID, LatLng, or Address.")
	}

	if lookupType == firststreet.FSIDLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/summary/city/%s", lookup.FSIDString())
	}

	if lookupType == firststreet.CoordinateLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/summary/city?lat=%s&lng=%s", lookup.LatLngString("lat"), lookup.LatLngString("lng"))
	}

	if lookupType == firststreet.AddressLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/summary/city?address=%s", lookup.Address)
	}

	err = c.B.Call(http.MethodGet, path, summaryResponse)
	return summaryResponse, err
}
