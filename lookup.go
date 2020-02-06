package firststreet

import (
	"errors"
	"strconv"
)

// Lookup determines which type of lookup to request to the API service
type Lookup struct {
	FSID    string
	Lat     float64
	Lng     float64
	Address string
}

// LookupType describes which type of lookup to use
type LookupType string

const FSIDLookup = LookupType("fsid")
const AddressLookup = LookupType("address")
const CoordinateLookup = LookupType("coordinate")

type LocationType string

const PropertyLocationType = LocationType("property")
const CityLocationType = LocationType("city")

// FSIDIsValid returns true if FSID is not ""
func (l *Lookup) FSIDIsValid() bool {
	return l.FSID != ""
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
func (l *Lookup) LookupType() (LookupType, error) {
	if l.FSIDIsValid() {
		return FSIDLookup, nil
	}

	if l.LatLngIsValid() {
		return CoordinateLookup, nil
	}

	if l.AddressIsValid() {
		return AddressLookup, nil
	}

	return "", errors.New("Lookup type could not be determined: Expecting an FSID, Lat and Lng, or Address.")
}
