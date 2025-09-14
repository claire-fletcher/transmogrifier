//go:generate mockgen -source=intensity-finder.go -destination=mock/intensity-finder.go
package carbon

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// TODO: consider correct interface and mock implementation in consumer
type CarbonIntensityFinder interface {
	GetCurrentCarbonIntensity() (int, error)
}

type finder struct {
	currentIntensitySource string
}

func CreateCarbonIntensityFinder(url string) finder {
	return finder{currentIntensitySource: url}
}

// TODO: separate out to generic http helpers
func (f finder) GetCurrentCarbonIntensity() (int, error) {
	// Create a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", f.currentIntensitySource, nil)
	if err != nil {
		return 0, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	// TODO: handle responses
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	data := UKCIResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	return data.Data[0].Intensity.Actual, nil //TODO: fix the assumption we always get one data response
}
