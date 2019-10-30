package firststreet

type ImpactValues struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type DataMVI struct {
	Year int64          `json:"year"`
	Data []ImpactValues `json:"data"`
}

type ResultsMVI struct {
	FloodType string    `json:"floodType"`
	FloodID   string    `json:"floodID"`
	Data      []DataMVI `json:"data"`
}

type MVI struct {
	FSID    int64        `json:"FSID"`
	Results []ResultsMVI `json:"results"`
}
