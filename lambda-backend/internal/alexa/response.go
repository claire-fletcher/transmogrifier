package alexa

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

type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              ResBody                `json:"response"`
}

type ResBody struct {
	OutputSpeech     *Payload `json:"outputSpeech,omitempty"`
	ShouldEndSession bool     `json:"shouldEndSession"`
}

type Payload struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}
