package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/claire-fletcher/transmogrifier/internal/alexa"
)

//TODO: stripped back to basics to just get that to work first.

// Handler is the lambda hander
func Handler() (alexa.Response, error) {
	return alexa.NewSimpleResponse("testing", "Hello from Lambda!"), nil
}

func main() {
	// Trigger
	lambda.Start(Handler)
}
