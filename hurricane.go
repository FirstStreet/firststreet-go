package firststreet

type DepthValuesHurricane struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type FloodDataHurricane struct {
	Year     int64                  `json:"year"`
	RiskData []DepthValuesHurricane `json:"data"`
}

type ResultsHurricane struct {
	FloodID   string               `json:"floodID"`
	FloodData []FloodDataHurricane `json:"floodData"`
}

type Hurricane struct {
	FSID    int64              `json:"FSID"`
	Results []ResultsHurricane `json:"results"`
}
