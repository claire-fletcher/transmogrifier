package alexa

// NewSimpleResponse creates a simple Alexa response with the given title and text. Specifically a speeech only response.
func NewSimpleResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: &Payload{
				Type: "PlainText",
				Text: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

// Response represents the structure of an Alexa response.
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              ResBody                `json:"response"`
}

// ResBody represents the body of an Alexa response. This currently only includes speech.
type ResBody struct {
	OutputSpeech     *Payload `json:"outputSpeech,omitempty"`
	ShouldEndSession bool     `json:"shouldEndSession"`
}

// Payload represents the payload of an Alexa response, including type, title, and text.
type Payload struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}
