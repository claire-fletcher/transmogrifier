package CarbonIntensityFinder

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

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

func (cif CarbonIntensityFinder) GetCurrentCarbonIntensity() (int, error) {

	// custom http client with timeout is needed as the default one has no timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Current half an hour intensity
	// TODO: can create the request outside if we want to
	// Could also create a generic make request thing for http helpers, only if start to do it a lot
	resp, err := client.Do(&http.Request{
		Method: "GET",
		URL:    cif.CurrentIntensitySource,
	})
	if err != nil {
		return 0, err
	}
	// TODO: handle responses

	defer resp.Body.Close()
	// Read and print response
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
