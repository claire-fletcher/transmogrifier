package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/claire-fletcher/transmogrifier/internal/alexa"
	"github.com/claire-fletcher/transmogrifier/internal/carbon"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

type Transmogrifier struct {
	cif carbon.CarbonItensityFinder
}

func NewTransmogrifier(c carbon.CarbonItensityFinder) Transmogrifier {
	return Transmogrifier{
		cif: c,
	}
}

func (t Transmogrifier) HandleGeneric() alexa.Response {
	return alexa.NewSimpleResponse("testing", "Hello from Lambda!")
}

func (t Transmogrifier) HandleCarbonIntensity() alexa.Response {
	ci, err := t.cif.GetCurrentCarbonIntensity()
	if err != nil {
		return alexa.NewSimpleResponse("Error", "There was an error getting the carbon intensity.")
	}

	return alexa.NewSimpleResponse("Carbon Intensity", "The current carbon intensity is "+fmt.Sprint(ci))
}

/** Below here is the lambda specific work. we could split out transmogrifier into its own thing too but it is the MAIN part
    Decide based on readability of what the code is doing.
**/

func IntentDispatcher(t Transmogrifier, request alexa.Request) alexa.Response {

	var response alexa.Response
	switch request.Body.Intent.Name {
	case "GetCurrentCarbonIntensity":
		response = t.HandleCarbonIntensity()
	default:
		response = t.HandleGeneric()
	}

	return response
}

// This is the specific lambda handler for a request coming in
func Handler(request alexa.Request) (alexa.Response, error) {
	ukcif := carbon.CreateCarbonIntensityFinder("https://api.carbonintensity.org.uk/intensity")
	t := NewTransmogrifier(ukcif)

	return IntentDispatcher(t, request), nil
}

func main() {
	// Trigger
	lambda.Start(otellambda.InstrumentHandler(Handler))
}
