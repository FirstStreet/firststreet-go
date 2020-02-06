package hurricane

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

func (c Client) Lookup(locationType firststreet.LocationType, lookup *firststreet.Lookup) (*firststreet.FloodRiskData, error) {
	var path string
	var err error

	mviResponse := &firststreet.FloodRiskData{}

	lookupType, err := lookup.LookupType()

	if err != nil {
		return mviResponse, err
	}

	if lookupType == "" {
		return mviResponse, errors.New("Lookup type could not be determined: Expecting an FSID, LatLng, or Address.")
	}

	if lookupType == firststreet.FSIDLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/hurricane/"+string(locationType)+"/%s", lookup.FSID)
	}

	if lookupType == firststreet.CoordinateLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/hurricane/"+string(locationType)+"?lat=%s&lng=%s", lookup.LatLngString("lat"), lookup.LatLngString("lng"))
	}

	if lookupType == firststreet.AddressLookup {
		path = backend.FormatURLPath("/data/"+c.B.Version+"/hurricane/"+string(locationType)+"?address=%s", lookup.Address)
	}

	err = c.B.Call(http.MethodGet, path, mviResponse)
	return mviResponse, err
}
