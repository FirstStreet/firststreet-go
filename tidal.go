package firststreet

type DepthValuesTidal struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type FloodDataTidal struct {
	Year     int64              `json:"year"`
	RiskData []DepthValuesTidal `json:"data"`
}

type ResultsTidal struct {
	FloodID   string           `json:"floodID"`
	FloodData []FloodDataTidal `json:"floodData"`
}

type Tidal struct {
	FSID    int64          `json:"FSID"`
	Results []ResultsTidal `json:"results"`
}
