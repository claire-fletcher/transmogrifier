package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/claire-fletcher/transmogrifier/internal/alexa"
	CarbonIntensityFinder "github.com/claire-fletcher/transmogrifier/internal/carbon-intensity-finder"
)

func IntentDispatcher(request alexa.Request) alexa.Response {

	var response alexa.Response
	switch request.Body.Intent.Name {
	case "CarbonIntensityIntent":
		response = HandleCarbonIntensity()
	default:
		response = HandleGeneric()
	}

	return response
}

func HandleGeneric() alexa.Response {
	return alexa.NewSimpleResponse("testing", "Hello from Lambda!")
}

func HandleCarbonIntensity() alexa.Response {
	ukci, err := CarbonIntensityFinder.CreateCarbonIntensityFinder("https://api.carbonintensity.org.uk/intensity")
	if err != nil {
		return alexa.NewSimpleResponse("Error", "There was an error getting the carbon intensity.")
	}
	currentCI := ukci.GetCurrentCarbonIntensity()

	return alexa.NewSimpleResponse("Carbon Intensity", "The current carbon intensity is "+string(rune(currentCI)))
}

// This is the specific lambda handler for a request coming in
func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	// Trigger
	lambda.Start(Handler)
}
