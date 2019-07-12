package parcel

import (
	"net/http"
	"strconv"

	"github.com/firststreet/floodiq-go/backend"
	geojson "github.com/paulmach/go.geojson"
)

// version is the current product version
const version = "1.0"

// A ParcelGeometry contains a polygon (boundingbox)
// and the Bounds (viewport) of a location
type ParcelGeometry struct {
	Polygon geojson.Geometry `json:"polygon"`
	Center  geojson.Geometry `json:"center"`
	Bounds  geojson.Geometry `json:"bounds"`
}

// A property parcel is a basic data struct for Property
type ParcelProperty struct {
	ID            int64                  `json:"ID"`
	PrimaryNumber string                 `json:"primaryNumber"`
	StreetName    string                 `json:"streetName"`
	LastUpdated   string                 `json:"lastUpdated"`
	Predirection  string                 `json:"predirection"`
	Postdirection string                 `json:"postdirection"`
	ZipCode       int                    `json:"zipCode"`
	City          *ParcelCityForProperty `json:"city"`
	State         string                 `json:"state"`
	Geometry      *ParcelGeometry        `json:"geometry"`
	Elevation     int64                  `json:"elevation"`
	LotSize       int64                  `json:"lotSize"`
	HomeSize      int64                  `json:"floorArea"`
	LandUse       string                 `json:"landUse"`
	CountyFIPS    int64                  `json:"countyFips"`
	Distance      float64                `json:"distance"`
}

type ParcelCityForProperty struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
}

type ParcelCity struct {
	ID       int64           `json:"ID"`
	Name     string          `json:"name"`
	State    string          `json:"state"`
	Geometry *ParcelGeometry `json:"geometry"`
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
