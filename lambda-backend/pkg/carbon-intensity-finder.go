package CarbonIntensityFinder

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

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
	response, err := http.Get(cif.CurrentIntensitySource.Host)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
	}

	// Read response information
	// TODO: need to create a response which alexa can handle
	// TODO: this can be in a different function for a generic "create alexa response"
}
