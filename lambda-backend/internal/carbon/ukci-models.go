package carbon

// UKCIResponse represents the structure of the response from the UK Carbon Intensity API
type UKCIResponse struct {
	Data []Data `json:"data"`
}

// Data represents the structure of each data entry in the UK Carbon Intensity API response
type Data struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Intensity Intensity `json:"intensity"`
}

// Intensity represents the intensity details in the UK Carbon Intensity API response
type Intensity struct {
	Forecast int    `json:"forecast"`
	Actual   int    `json:"actual"`
	Index    string `json:"index"`
}
