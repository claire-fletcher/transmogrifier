package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/claire-fletcher/transmogrifier/internal/alexa"
)

//TODO: stripped back to basics to just get that to work first.

func IntentDispatcher(request alexa.Request) alexa.Response {

	log.Printf("Received request: %+v", request) // TODO: Debugging

	var response alexa.Response
	switch request.Body.Intent.Name {
	default:
		response = HandleGeneric()
	}

	return response
}

func HandleGeneric() alexa.Response {
	return alexa.NewSimpleResponse("testing", "Hello from Lambda!")
}

// This is the specific lambda handler for a request coming in
func Handler(request alexa.Request) (alexa.Response, error) {
	log.Printf("Handler received request: %+v", request) // TODO: Debugging

	return IntentDispatcher(request), nil
}

func main() {
	// Trigger
	lambda.Start(Handler)
}
