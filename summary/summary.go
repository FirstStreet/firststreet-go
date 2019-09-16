package summary

import (
	"net/http"
	"strconv"

	"github.com/firststreet/floodiq-go/backend"
	"github.com/firststreet/floodiq-go/jsonize"
	geojson "github.com/paulmach/go.geojson"
)

// version is the current product version
const version = "1.0"

// A LocationGeometry contains a polygon (boundingbox)
// and the Bounds (viewport) of a location
type LocationGeometry struct {
	Polygon geojson.Geometry `json:"polygon"`
	Center  geojson.Geometry `json:"center"`
	Bounds  geojson.Geometry `json:"bounds"`
}

type Location struct {
	FSID int32
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

// Client is used for /data/{svc}
type Client struct {
	B   *backend.Backend
	Key string
}

// GetPropertyByID retreives a Property Parcel by its unique identifier
func (c Client) GetPropertyByID(id string) (*ParcelProperty, error) {
	path := backend.FormatURLPath("/data/"+version+"/parcel/%s?type=property&key=%s", id, c.Key)
	property := &ParcelProperty{}
	err := c.B.Call(http.MethodGet, path, c.Key, property)
	return property, err
}

// GetPropertyByLatLng pulls a parcel by lat lng
func (c Client) GetPropertyByLatLng(lat, lng float64) (*ParcelProperty, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)

	path := backend.FormatURLPath("/data/"+version+"/parcel?type=property&key=%s&lat=%s&lng=%s", c.Key, latStr, lngStr)
	property := &ParcelProperty{}
	err := c.B.Call(http.MethodGet, path, c.Key, property)
	return property, err
}

// GetCityByID retreives a City Parcel by its unique identifier
func (c Client) GetCityByID(id string) (*ParcelCity, error) {
	path := backend.FormatURLPath("/data/"+version+"/parcel/%s?type=city&key=%s", id, c.Key)
	city := &ParcelCity{}
	err := c.B.Call(http.MethodGet, path, c.Key, city)
	return city, err
}
