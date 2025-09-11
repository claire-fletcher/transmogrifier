package main_test

import (
	"testing"

	mock_CarbonIntensityFinder "github.com/claire-fletcher/transmogrifier/internal/carbon-intensity-finder/mock"
	"go.uber.org/mock/gomock"

	. "github.com/claire-fletcher/transmogrifier/cmd"
)

func TestHandleCarbonIntensityReturnsTheIntensity(t *testing.T) {

	/** Arrange **/
	testIntensity := 123
	mockCtrl := gomock.NewController(t)
	mockCFI := mock_CarbonIntensityFinder.NewMockCarbonItensityFinder(mockCtrl)
	mockCFI.EXPECT().GetCurrentCarbonIntensity().Return(testIntensity, nil)

	testTransmogrifier := NewTransmogrifier(mockCFI)

	/** Act **/
	response := testTransmogrifier.HandleCarbonIntensity()

	/** Assert **/
	expectedSpeech := "The current carbon intensity is 123"
	if response.Body.OutputSpeech.Text != expectedSpeech {
		t.Errorf("Expected response to be '%s' but got '%s'", expectedSpeech, response.Body.OutputSpeech.Text)
	}
}
