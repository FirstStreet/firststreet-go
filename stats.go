package firststreet

type FloodStatistic struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type FloodData struct {
	Year     int64            `json:"year"`
	RiskData []FloodStatistic `json:"data"`
}

type ResultsData struct {
	FloodID   string      `json:"floodID"`
	FloodData []FloodData `json:"floodData"`
}

type FloodRiskData struct {
	FSID    int64         `json:"FSID"`
	Results []ResultsData `json:"results"`
}
