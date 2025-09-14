package main_test

import (
	"testing"

	"go.uber.org/mock/gomock"

	. "github.com/claire-fletcher/transmogrifier/cmd"
	mock_carbon "github.com/claire-fletcher/transmogrifier/internal/carbon/mock"
)

func TestHandleCarbonIntensityReturnsTheIntensity(t *testing.T) {

	/** Arrange **/
	testIntensity := 123
	mockCtrl := gomock.NewController(t)
	mockCFI := mock_carbon.NewMockCarbonItensityFinder(mockCtrl)
	mockCFI.EXPECT().GetCurrentCarbonIntensity().Return(testIntensity, nil)

	/** Act **/
	response := HandleCarbonIntensity(mockCFI)

	/** Assert **/
	expectedSpeech := "The current carbon intensity is 123"
	if response.Body.OutputSpeech.Text != expectedSpeech {
		t.Errorf("Expected response to be '%s' but got '%s'", expectedSpeech, response.Body.OutputSpeech.Text)
	}
}
