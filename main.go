package main

import (
	"fmt"
	"net/http"
	"os"
	"weather-api-go/constants"
	"weather-api-go/environment"
	"weather-api-go/handlers"
	"weather-api-go/services"
	"weather-api-go/utils"
)

func main() {

	err := environment.LoadEnvs()
	if err != nil {
		panic(err)
	}

	apiToken := os.Getenv(constants.API_TOKEN)
	if apiToken == "" {
		panic("WEATHER API token is not provided")
	}
	port := os.Getenv(constants.PORT)
	if port == "" {
		panic("port is not provided")
	}

	go utils.ShowUptimeMessage()

	services.WeatherService.Token = apiToken

	http.HandleFunc("/health", environment.MiddlewareJSON(handlers.Health))
	http.HandleFunc("/weather", environment.MiddlewareJSON(handlers.Weather))

	fmt.Println("Serve the app")
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
