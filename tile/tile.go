package tile

import (
	"fmt"

	"github.com/firststreet/floodiq-go/backend"
)

// version is the current product version
const version = "1.0"

// The TileMetadata describes the X, Y position based on a Zoom
type TileMetadata struct {
	X int32
	Y int32
	Z int32
}

// The Tile provides an exact URL of where query the Tile
type Tile struct {
	URL       string
	Metadata  *TileMetadata
	FloodType string
	FloodYear int32
}

type URLPrefix string

// Client is used for /data/{svc}
type Client struct {
	B   *backend.Backend
	Key string
}

//http://apidev.firststreet.org/tile/1.0/hurricane/12/1132/1744?key=VpqgUoepulokFjdxZvE4iMjP8bTtN2PG&type=kt&year=2018

// Hurricane retreives a hurricane Tile by TileMetadata
func (c Client) Hurricane(id string) (*Tile, error) {
	// path := backend.FormatURLPath("/data/"+version+"/parcel/%s?type=property&key=%s", id, c.Key)
	// property := &ParcelProperty{}
	// err := c.B.Call(http.MethodGet, path, c.Key, property)
	// return property, err
	return &Tile{}, nil
}

// FromURLPrefix returns a valid URL from a URL Prefix
func (c Client) FromURLPrefix(prefix *URLPrefix) string {
	path := backend.FormatURLPath("/tile/" + version + "/" + prefix.String())

	return path
}

// url := t.FloodType + "/" + t.Metadata.Z + "/" + t.Metadata.X + "/" + t.Metadata.Y + "?type=" + t.FloodType + "&year=" + t.FloodYear
func (t *Tile) GetURLPrefix() *URLPrefix {
	url := new(URLPrefix)
	url = t.FloodType + "/" + t.Metadata.Z + "/" + t.Metadata.X + "/" + t.Metadata.Y + "?type=" + t.FloodType + "&year=" + t.FloodYear
	return url
}

// Returns URLPrefix as a string
func (p *URLPrefix) String() string {
	return fmt.Sprintf("%b", p)
}

// Takes an int32 and returns a string
func yearToStr(year int32) string {

}
