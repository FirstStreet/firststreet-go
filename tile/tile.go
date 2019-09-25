package tile

import (
	"net/http"
	"strconv"

	"github.com/firststreet/firststreet-go/backend"
)

// TilePointCoordinate indicates a X, Y, or Z value
type TilePointCoordinate int32

// FloodType is the requested flood for the payload
type FloodType string

const (
	// Current version of the tile service for this release
	version = "1.0"
	// Category1 is a Category 1 Hurricane
	Category1 FloodType = "c1"
	// Category2 is a Category 2 Hurricane (Note, this may not be provided by the API)
	Category2 FloodType = "c2"
	// Category3 is a Category 3 Hurricane
	Category3 FloodType = "c3"
	// Category4 is a Category 4 Hurricane (Note, this may not be provided by the API)
	Category4 FloodType = "c4"
	// Category5 is a Category 5 Hurricane
	Category5 FloodType = "c5"
	// KingTide is King tide flooding
	KingTide FloodType = "kt"
	// Highest Annaul Tidal Flooding is the highest annual tidal flooding
	HighestAnnualTidalFlooding FloodType = "em"
)

// The TileMetadata describes the X, Y position based on a Zoom
type TileMetadata struct {
	X         TilePointCoordinate
	Y         TilePointCoordinate
	Z         TilePointCoordinate
	FloodType FloodType
	FloodYear int32
}

// NewTileMetadata creates a TileMetadata for use
func NewTileMetadata(z, x, y int32, floodType FloodType, floodYear int32) *TileMetadata {
	return &TileMetadata{
		X:         TilePointCoordinate(x),
		Y:         TilePointCoordinate(y),
		Z:         TilePointCoordinate(z),
		FloodType: floodType,
		FloodYear: floodYear,
	}
}

// Tile is the tile from the requested Metadata
type Tile struct {
	// Metadata describes information about the requested tile
	Metadata TileMetadata
	// URL is the absolute path to the tile, with the API key
	URL string
	// Bytes is the resolved tile data
	Bytes []byte
}

// Client is used for /data/{svc}
type Client struct {
	B   *backend.Backend
	Key string
}

func (pc TilePointCoordinate) String() string {
	return strconv.Itoa(int(pc))
}

//http://apidev.firststreet.org/tile/1.0/hurricane/12/1132/1744?key=VpqgUoepulokFjdxZvE4iMjP8bTtN2PG&type=kt&year=2018

// Hurricane retreives a hurricane Tile by TileMetadata
func (c Client) Hurricane(z, x, y int32, floodType FloodType, floodYear int32) (*Tile, error) {
	tileMetadata := NewTileMetadata(z, x, y, floodType, floodYear)
	// property := &ParcelProperty{}
	// err := c.B.Call(http.MethodGet, path, c.Key, property)
	// return property, err
	tile := &Tile{
		Metadata: *tileMetadata,
	}

	path := backend.FormatURLPath("/data/"+version+"/hurricane/%s/%s/%s.png?&type=%s&year=%s", tile.Metadata.Z.String(), tile.Metadata.X.String(), tile.Metadata.Y.String(), string(tile.Metadata.FloodType), "2018")
	bytes, err := c.B.CallBytes(http.MethodGet, path, c.Key)
	if err != nil {
		return nil, err
	}
	tile.Bytes = bytes
	tile.URL = path
	return tile, nil
}
