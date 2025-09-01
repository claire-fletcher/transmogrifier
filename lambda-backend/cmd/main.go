package main

import (
	"log"

	CarbonIntensityFinder "github.com/claire-fletcher/transmogrifier/pkg"
	"github.com/aws/aws-lambda-go/lambda"
)

//TODO: stripped back to basics to just get that to work first.

// AlexaResponse contains the message
type AlexaResponse struct {
	Version string  `json:"version"`
	Body    ResBody `json:"response"`
}

// ResBody is the actual body of the response
type ResBody struct {
	OutputSpeech     Payload ` json:"outputSpeech,omitempty"`
	ShouldEndSession bool    `json:"shouldEndSession"`
}

// Payload ...
type Payload struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

// NewResponse builds a simple Alexa session response
func NewResponse(speech string) Response {
	return AlexarResponse {
		Version: "1.0",
		Body: ResBody {
			OutputSpeech: Payload {
				Type: "PlainText",
				Text: speech,
			},
			ShouldEndSession: true,
		},
	}
}

// Handler is the lambda hander
func Handler() (Response, error) {
	return NewResponse("testing this basic version"), nil
}

func main() {

	// setup the carbon intensity finder for the UK CI API
	// ukci, err := CarbonIntensityFinder.CreateCarbonIntensityFinder("https://api.carbonintensity.org.uk/intensity")
	// if err != nil {
	// 	log.Panicf("Unable to create carbon intensity finder due to error: %v", err)
	// }

	// Trigger
	lambda.Start(Handler)
}
