package firststreet

import (
	geojson "github.com/paulmach/go.geojson"
)

// FirstFloodRisk defines a first risk for a location
type FirstFloodRisk struct {
	FloodID string `json:"floodID"`
	Year    int32  `json:"year"`
}

type FloodRisk struct {
	Year    int32    `json:"year"`
	FloodID []string `json:"floodID"`
}

type FloodRisks struct {
	FloodRisk []*FloodRisk `json:"floodRisks"`
}

// A LocationGeometry contains a polygon (boundingbox)
// and the Bounds (viewport) of a location
type LocationGeometry struct {
	Polygon geojson.Geometry `json:"polygon"`
	Center  geojson.Geometry `json:"center"`
	BBox    geojson.Geometry `json:"bbox"`
}

type Property struct {
	FSID          string            `json:"FSID"`
	PrimaryNumber string            `json:"primaryNumber"`
	StreetName    string            `json:"streetName"`
	LastUpdated   string            `json:"lastUpdated"`
	Predirection  string            `json:"predirection"`
	Postdirection string            `json:"postdirection"`
	ZipCode       int64             `json:"zipCode"`
	City          *PropertyCity     `json:"city"`
	State         string            `json:"state"`
	Geometry      *LocationGeometry `json:"geometry"`
	Elevation     int64             `json:"elevation"`
	CountyFIPS    int64             `json:"countyFips"`
}

type PropertyCity struct {
	FSID string `json:"FSID"`
	Name string `json:"Name"`
}

type City struct {
	FSID     string            `json:"FSID"`
	Name     string            `json:"name"`
	State    string            `json:"state"`
	Geometry *LocationGeometry `json:"geometry"`
}

type PropertySummaryResults struct {
	Location       Property
	FirstFloodRisk FirstFloodRisk
	FloodRisks     []FloodRisks
}

type CitySummaryResults struct {
	Location       City
	FirstFloodRisk FirstFloodRisk
	FloodRisks     []FloodRisks
}

type PropertySummary struct {
	FSID    string                 `json:"FSID"`
	Results PropertySummaryResults `json:"results"`
}

type CitySummary struct {
	FSID    string             `json:"FSID"`
	Results CitySummaryResults `json:"results"`
}
