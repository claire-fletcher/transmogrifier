# transmogrifier
An Alexa skill to get and use information about the current carbon intensity using the carbon aware SDK. Enabling "transmogrification" of electricity usage into carbon aware usage - a.k.a. enabling demand shifting. 

# Initial Design - MVP v1.0.0
There are a few options including an SDK with more control over the data we collect or alternatively using already hosted APIs to get the data.
To start with it is one less complexity to use the APIs and this will specifically use the UK's national grid API. This has good documentation, no tiered usage, and the most aligned terms of use.

![Initial Design](initial-design.png)

## Front End
The front end will  be created with the developer skills framework in the console. This can then generate the initial json which will then be stored here and can be updated with an API CI job.

### Invocation:
- "electro transmorgify"

### Intents and utterance mapping:
- Utterance: "What is the current carbon intensity" -> GetCurrentCarbonIntensity

Note: There will be no slots for this. They may be used for the forecast extension to give a specific time. Future considerations: a dialog model.


## Back End
This is where the requests from alexa will come in and be handled. These will be API routes based on the different intents. They will then get the relevant information from the CI APIs, transform it to visual/words in a request response, and send this back to Alexa.


## Links:
- https://carbon-aware-sdk.greensoftware.foundation/docs/tutorial-basics/carbon-aware-webapi
- https://carbon-intensity.github.io/api-definitions/#carbon-intensity-api-v2-0-0
- Alternative APIs: https://developers.thegreenwebfoundation.org/grid-intensity-cli/explainer/providers/
- https://developer.amazon.com/en-US/docs/alexa/custom-skills/steps-to-build-a-custom-skill.html

## Future
1. Using the SDK instead also hosted as a function
2. Additional options for the skill such as more nuanced questions, data analysis etc.
3. "What is the forecasted carbon intensity for the next hour"