//go:generate mockgen -source=carbon-intensity-finder.go -destination=mock/carbon-intensity-finder.go
package CarbonIntensityFinder

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CarbonItensityFinder interface {
	GetCurrentCarbonIntensity() (int, error)
}

type CarbonIntensityFinder struct {
	CurrentIntensitySource *url.URL
}

func CreateCarbonIntensityFinder(currentIntensitySource string) (*CarbonIntensityFinder, error) {
	u, err := url.Parse(currentIntensitySource)
	if err != nil {
		return nil, err
	}

	return &CarbonIntensityFinder{CurrentIntensitySource: u}, nil
}

// TODO: separate out to generic http helpers
func (cif CarbonIntensityFinder) GetCurrentCarbonIntensity() (int, error) {
	// custom http client with timeout is needed as the default one has no timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(&http.Request{
		Method: "GET",
		URL:    cif.CurrentIntensitySource,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	})
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
