package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/claire-fletcher/transmogrifier/internal/alexa"
	"github.com/claire-fletcher/transmogrifier/internal/carbon"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

func HandleGeneric() alexa.Response {
	return alexa.NewSimpleResponse("testing", "Hello from Lambda!")
}

func HandleCarbonIntensity(cif carbon.CarbonIntensityFinder) alexa.Response {
	ci, err := cif.GetCurrentCarbonIntensity()
	if err != nil {
		return alexa.NewSimpleResponse("Error", "There was an error getting the carbon intensity.")
	}

	return alexa.NewSimpleResponse("Carbon Intensity", "The current carbon intensity is "+fmt.Sprint(ci))
}

/** Below here is the lambda specific work. we could split out transmogrifier into its own thing too but it is the MAIN part
    Decide based on readability of what the code is doing.
**/

func IntentDispatcher(cif carbon.CarbonIntensityFinder, request alexa.Request) alexa.Response {

	var response alexa.Response
	switch request.Body.Intent.Name {
	case "GetCurrentCarbonIntensity":
		response = HandleCarbonIntensity(cif)
	default:
		response = HandleGeneric()
	}

	return response
}

// This is the specific lambda handler for a request coming in
func Handler(request alexa.Request) (alexa.Response, error) {
	ukcif := carbon.CreateCarbonIntensityFinder("https://api.carbonintensity.org.uk/intensity")

	return IntentDispatcher(ukcif, request), nil
}

func main() {
	// Trigger
	lambda.Start(otellambda.InstrumentHandler(Handler))
}
