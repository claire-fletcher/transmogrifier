package CarbonIntensityFinder

import (
	"net/url"
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

func (cif CarbonIntensityFinder) GetCurrentCarbonIntensity() int {
	// TODO: get the actual intensity
	// USE cif.CurrentIntensitySource.String() as an http url and set up http request to the api

	return 200
}
