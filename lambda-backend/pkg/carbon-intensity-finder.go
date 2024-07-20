package CarbonIntensityFinder

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type AlexaSpeechResponse struct {
	Type         string       `json:"type"`
	Version      string       `json:"version"`
	MainTemplate MainTemplate `json:"mainTemplate"`
}

type MainTemplate struct {
	Parameters []string `json:"parameters"`
	Item       Item     `json:"item"`
}

type Item struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type CarbonIntensityFinder struct {
	CurrentIntensitySource url.URL
}

func CreateCarbonIntensityFinder(currentIntensitySource string) (*CarbonIntensityFinder, error) {
	u, err := url.Parse(currentIntensitySource)
	if err != nil {
		return nil, err
	}

	return &CarbonIntensityFinder{CurrentIntensitySource: *u}, nil
}

func (cif CarbonIntensityFinder) GetCurrentCarbonIntensity(c *gin.Context) {
	// TODO: get the actual intensity

	// Create response for Alexa
	audio := AlexaSpeechResponse{
		Type:    "APLA",
		Version: "0.91",
		MainTemplate: MainTemplate{
			Parameters: []string{"payload"},
			Item: Item{
				Type:    "Speech",
				Content: "The current carbon intensity is 200",
			},
		},
	}

	// Use the JSON response
	c.IndentedJSON(http.StatusOK, audio)
}

// TODO: check if this comes from amazon
// TODO: separate out the alexa responses into an alexa response writer
