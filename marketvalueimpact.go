package firststreet

type ImpactValues struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type FloodDataMVI struct {
	Year int64          `json:"year"`
	Data []ImpactValues `json:"data"`
}

type ResultsMVI struct {
	FloodID   string         `json:"floodID"`
	FloodData []FloodDataMVI `json:"floodData"`
}

type MVI struct {
	FSID    int64        `json:"FSID"`
	Results []ResultsMVI `json:"results"`
}
