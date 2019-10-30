package firststreet

type DepthValuesTidal struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type DataTidal struct {
	Year     int64              `json:"year"`
	RiskData []DepthValuesTidal `json:"data"`
}

type ResultsTidal struct {
	FloodType string      `json:"floodType"`
	FloodID   string      `json:"floodID"`
	Data      []DataTidal `json:"data"`
}

type Tidal struct {
	FSID    int64          `json:"FSID"`
	Results []ResultsTidal `json:"results"`
}
