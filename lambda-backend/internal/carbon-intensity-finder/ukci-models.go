package CarbonIntensityFinder

type UKCIResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Intensity Intensity `json:"intensity"`
}

type Intensity struct {
	Forecast int    `json:"forecast"`
	Actual   int    `json:"actual"`
	Index    string `json:"index"`
}
