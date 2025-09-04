package CarbonIntensityFinder

type UKCIResponse struct {
	Data []struct {
		From      string `json:"from"`
		To        string `json:"to"`
		Intensity struct {
			Forecast int    `json:"forecast"`
			Actual   int    `json:"actual"`
			Index    string `json:"index"`
		} `json:"intensity"`
	} `json:"data"`
}
