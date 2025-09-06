package CarbonIntensityFinder

import (
	"testing"

	mock "github.com/claire-fletcher/transmogrifier/internal/carbon-intensity-finder/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCarbonIntensityFinderReturnsValue(t *testing.T) {

	/** Arrange **/
	ctrl := gomock.NewController(t)
	mockCIF := mock.NewMockCarbonItensityFinder(ctrl)
	mockCIF.EXPECT().GetCurrentCarbonIntensity().Return(100, nil).Times(1)

	// TODO:  may need to split out the call to the api so that we can mock it here as otherwise we are just saying, return 100.

	/** Act **/
	intensity, err := mockCIF.GetCurrentCarbonIntensity()

	/** Assert **/
	assert.NoError(t, err, "There should be no error")
	assert.Equal(t, intensity, 100, "The carbon intensity should be 100")
}

// TODO: Test suites to set up the specific calls and returns so that we can test multiple scenarios including failures more easily
// TODO: also best practice to work with a context
// Can use interfacer to create interfaces from the struct to make mocking easier
