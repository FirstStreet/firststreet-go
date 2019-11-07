package firststreet

import (
	"errors"
	"strconv"
)

// Lookup determines which type of lookup to request to the API service
type Lookup struct {
	FSID    int64
	Lat     float64
	Lng     float64
	Address string
}

// FSIDIsValid returns true if FSID is not 0
func (l *Lookup) FSIDIsValid() bool {
	return l.FSID != 0
}

// FSIDString returns FSID as a string
func (l *Lookup) FSIDString() string {
	return strconv.FormatInt(l.FSID, 10)
}

// Address is valid if address is not empty
func (l *Lookup) AddressIsValid() bool {
	return l.Address != ""
}

// LatLngIsValid returns true if lat and lng are not 0
func (l *Lookup) LatLngIsValid() bool {
	return l.Lat != 0 || l.Lng != 0
}

// LatLngString provides Lat or Lng as a string value
func (l *Lookup) LatLngString(latlng string) string {
	if !l.LatLngIsValid() {
		return ""
	}

	if latlng == "lat" {
		return strconv.FormatFloat(l.Lat, 'f', -1, 64)
	}

	if latlng == "lng" {
		return strconv.FormatFloat(l.Lng, 'f', -1, 64)
	}

	return ""
}

// LookupType returns the requested lookup type
func (l *Lookup) LookupType() (string, error) {
	if l.FSIDIsValid() {
		return "fsid", nil
	}

	if l.LatLngIsValid() {
		return "coordinate", nil
	}

	if l.AddressIsValid() {
		return "address", nil
	}

	return "", errors.New("Lookup type could not be determined: Expecting an FSID, LatLng, or Address.")
}
