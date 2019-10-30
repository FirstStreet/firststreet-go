package firststreet

type DepthValuesHurricane struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type DataHurricane struct {
	Year     int64                  `json:"year"`
	RiskData []DepthValuesHurricane `json:"data"`
}

type ResultsHurricane struct {
	FloodType string          `json:"floodType"`
	FloodID   string          `json:"floodID"`
	Data      []DataHurricane `json:"data"`
}

type Hurricane struct {
	FSID    int64              `json:"FSID"`
	Results []ResultsHurricane `json:"results"`
}
