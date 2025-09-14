// Package CarbonItensityFinder provides a way to get the current carbon intensity from a given source
//
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

// CarbonIntensityFinder is a struct to hold relevant dependencies for finding the current carbon intensity.
type CarbonIntensityFinder struct {
	CurrentIntensitySource *url.URL
}

// CreateCarbonIntensityFinder creates a new instance of CarbonIntensityFinder with the given source URL.
func CreateCarbonIntensityFinder(currentIntensitySource string) (*CarbonIntensityFinder, error) {
	u, err := url.Parse(currentIntensitySource)
	if err != nil {
		return nil, err
	}

	return &CarbonIntensityFinder{CurrentIntensitySource: u}, nil
}

// TODO: separate out to generic http helpers
// GetCurrentCarbonIntensity fetches the current carbon intensity from the configured source.
// It will return a value in gCO2/kWh or an error if something goes wrong.
func (cif CarbonIntensityFinder) GetCurrentCarbonIntensity() (int, error) {
	// custom http client with timeout is needed as the default one has no timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// TODO: consider a context instead and use a request with context
	// instead of a client with timeout
	// This is more flexible as the request can have a timeout and the client can have different config.
	// each request can differ

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

// do a http request, check if that failed
// read the bosy, check if reading failed
// read that body into a given struct, check if that failed
// we could have a http package with request helper to do this for us and then store and err and only return at the end
// turn each step into a no-op if the previous failed? errRequester, with err() method like in the scan thing
// at least give it an attempt and see what happens
