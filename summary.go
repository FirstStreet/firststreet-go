package firststreet

import (
	"github.com/firststreet/firststreet-go/jsonize"
	geojson "github.com/paulmach/go.geojson"
)

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
	CountyFIPS    jsonize.JsonNullInt64  `json:"countyFips"`
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

type SummaryResults struct {
	Location Location
	// @TODO: Summary
}

type Summary struct {
	FSID    int64          `json:"FSID"`
	Results SummaryResults `json:"results"`
}
