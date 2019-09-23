package summary

import (
	"net/http"
	"strconv"

	"github.com/firststreet/floodiq-go/backend"
	"github.com/firststreet/floodiq-go/jsonize"
	geojson "github.com/paulmach/go.geojson"
)

// version is the current product version
const version = "v1.0"

// A LocationGeometry contains a polygon (boundingbox)
// and the Bounds (viewport) of a location
type LocationGeometry struct {
	Polygon geojson.Geometry `json:"polygon"`
	Center  geojson.Geometry `json:"center"`
	Bounds  geojson.Geometry `json:"bounds"`
}

type Location struct {
	FSID int64
}

type Property struct {
	Location
	PrimaryNumber jsonize.JsonNullString `json:"primaryNumber"`
	StreetName    jsonize.JsonNullString `json:"streetName"`
	LastUpdated   string                 `json:"lastUpdated"`
	Predirection  jsonize.JsonNullString `json:"predirection"`
	Postdirection jsonize.JsonNullString `json:"postdirection"`
	ZipCode       int                    `json:"zipCode"`
	City          *PropertyCity          `json:"city"`
	State         jsonize.JsonNullString `json:"state"`
	Geometry      *LocationGeometry      `json:"geometry"`
	Elevation     jsonize.JsonNullInt64  `json:"elevation"`
	// @TODO: FemaZone should be broken down into its own struct
	// There could also be more than 1 femazones
	FemaZone   jsonize.JsonNullString `json:"femaZone"`
	LotSize    jsonize.JsonNullInt64  `json:"lotSize"`
	HomeSize   jsonize.JsonNullInt64  `json:"floorArea"`
	LandUse    jsonize.JsonNullString `json:"landUse"`
	CountyFIPS jsonize.JsonNullInt64  `json:"countyFips"`
	Distance   float64                `json:"distance"`
}

type PropertyCity struct {
	FSID jsonize.JsonNullInt64  `json:"FSID"`
	Name jsonize.JsonNullString `json:"Name"`
}

type City struct {
	Location
	Name     jsonize.JsonNullString `json:"name"`
	State    jsonize.JsonNullString `json:"state"`
	Geometry *LocationGeometry      `json:"geometry"`
}

type Summary struct {
	FSID    int64          `json:"FSID"`
	Type    string         `json:"type"`
	Results SummaryResults `json:"results"`
}

type SummaryResults struct {
	Location Location
	// @TODO: Summary
}

// Client is used for /data/{svc}
type Client struct {
	B   *backend.Backend
	Key string
}

// GetPropertyByFSID retreives a Location Summary by its unique identifier
func (c Client) GetPropertyByFSID(fsid string) (*Summary, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/%s?type=property&key=%s", fsid, c.Key)
	summaryResponse := &Summary{}
	err := c.B.Call(http.MethodGet, path, c.Key, summaryResponse)
	return summaryResponse, err
}

// GetPropertyByLatLng pulls a parcel by lat lng
func (c Client) GetPropertyByLatLng(lat, lng float64) (*Summary, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/summary?type=property&key=%s&lat=%s&lng=%s", c.Key, latStr, lngStr)
	summaryResponse := &Summary{}
	err := c.B.Call(http.MethodGet, path, c.Key, summaryResponse)
	return summaryResponse, err
}

// GetPropertyByAddress pulls a parcel by lat lng
func (c Client) GetPropertyByAddress(address string) (*Summary, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary?locationType=property&key=%s&address=%s", c.Key, address)
	summaryResponse := &Summary{}
	err := c.B.Call(http.MethodGet, path, c.Key, summaryResponse)
	return summaryResponse, err
}

// GetCityByID retreives a City Parcel by its unique identifier
func (c Client) GetCityByID(id string) (*City, error) {
	path := backend.FormatURLPath("/data/"+version+"/summary/%s?type=city&key=%s", id, c.Key)
	city := &City{}
	err := c.B.Call(http.MethodGet, path, c.Key, city)
	return city, err
}

func (s Summary) Location() Location {
	return s.Results.Location
}
