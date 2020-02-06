package firststreet

import (
	"errors"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestNewBackend(t *testing.T) {
	backend := NewBackend("abc123")
	assert.Equal(t, "abc123", backend.Key)
}

func TestFSIDLookup(t *testing.T) {
	fsid := "123456"
	fsidLookup := &Lookup{
		FSID: fsid,
	}

	assert.Equal(t, fsidLookup.FSID, fsid)
	assert.True(t, fsidLookup.FSIDIsValid())
	assert.False(t, fsidLookup.AddressIsValid())
	assert.False(t, fsidLookup.LatLngIsValid())

	lookupType, err := fsidLookup.LookupType()
	assert.Nil(t, err)
	assert.Equal(t, lookupType, FSIDLookup)

	latlngString := fsidLookup.LatLngString("lat")
	assert.Equal(t, latlngString, "")
}

func TestLatLngLookup(t *testing.T) {
	latLngLookup := &Lookup{
		Lat: 12.34,
		Lng: 56.78,
	}

	assert.True(t, latLngLookup.LatLngIsValid())

	assert.Equal(t, latLngLookup.LatLngString("lat"), "12.34")
	assert.Equal(t, latLngLookup.LatLngString("lng"), "56.78")
	assert.Equal(t, latLngLookup.LatLngString("gibberish"), "")

	latLngLookupType, err := latLngLookup.LookupType()

	assert.Nil(t, err)
	assert.Equal(t, latLngLookupType, CoordinateLookup)

}

func TestAddressLookup(t *testing.T) {
	addressLookup := &Lookup{
		Address: "1234 Main St",
	}

	assert.True(t, addressLookup.AddressIsValid())

	addressLookupType, err := addressLookup.LookupType()
	assert.Nil(t, err)
	assert.Equal(t, addressLookupType, AddressLookup)
}

func TestBadLookup(t *testing.T) {
	badlookup := &Lookup{}

	badType, err := badlookup.LookupType()

	assert.Equal(t, badType, LookupType(""))
	assert.Error(t, err, errors.New("Lookup type could not be determined: Expecting an FSID, Lat and Lng, or Address."))
}

// func TestIntegration(t *testing.T) {
// 	api := New("VpqgUoepulokFjdxZvE4iMjP8bTtN2PG")
// 	parcel, _ := api.Parcel.GetPropertyByID("100032470544")
// 	t.Log(parcel)
// }
