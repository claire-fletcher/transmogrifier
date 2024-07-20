package main

import (
	"log"

	CarbonIntensityFinder "github.com/claire-fletcher/transmogrifier/pkg"
	"github.com/gin-gonic/gin"
)

func main() {

	// setup the carbon intensity finder for the UK CI API
	ukci, err := CarbonIntensityFinder.CreateCarbonIntensityFinder("https://api.carbonintensity.org.uk/intensity")
	if err != nil {
		log.Panicf("Unable to create carbon intensity finder due to error: %v", err)
	}

	// setup server with routes for getting the carbon intensity
	router := gin.Default()
	router.GET("/getCurrentCarbonIntensity", func(c *gin.Context) {
		ukci.GetCurrentCarbonIntensity(c)
	})

	router.Run(":8080")
}
