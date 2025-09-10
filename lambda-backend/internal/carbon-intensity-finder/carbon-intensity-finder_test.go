package CarbonIntensityFinder_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/claire-fletcher/transmogrifier/internal/carbon-intensity-finder"
)

func createMockUKCIResponse(actual int) string {
	return fmt.Sprintf(`{
		"data": {
			"from": "2024-01-01T00:00Z",
			"to": "2024-01-01T00:30Z",
			"intensity": {
				"forecast": 150,
				"actual": %d,
				"index": "moderate"
			}
		}
	}`, actual)
}

// TODO: consider the testify suite instead for better assertions
// TODO: should we be creating json or just creating a struct using a create method and then unmarshalling to json string?
func TestCarbonIntensityFinderReturnsValue(t *testing.T) {

	/** Arrange **/
	testForecast := 123
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, createMockUKCIResponse(testForecast))
	}))
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse the carbon intensity finder url: %v", err)
	}
	testCIF, err := CreateCarbonIntensityFinder(u.String())
	if err != nil {
		t.Fatalf("Failed to create the carbon intensity finder: %v", err)
	}

	/** Act **/
	result, err := testCIF.GetCurrentCarbonIntensity()
	if err != nil {
		t.Fatalf("Failed to get the current carbon intensity: %v", err)
	}

	/** Assert **/
	if result != testForecast {
		t.Errorf("Expected result to be %d, got %d", testForecast, result)
	}
}
