package firststreet

type DepthValues struct {
	Type  string  `json:"type"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type Data struct {
	Year     int64         `json:"year"`
	RiskData []DepthValues `json:"data"`
}

type ResultsHurricane struct {
	FloodType string `json:"floodType"`
	FloodID   string `json:"floodID"`
	Data      []Data `json:"data"`
}

type Hurricane struct {
	FSID    int64              `json:"FSID"`
	Results []ResultsHurricane `json:"results"`
}
