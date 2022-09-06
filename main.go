package main

import (
	"fmt"
	"net/http"
	"os"
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

	go utils.ShowUptimeMessage()

	services.WeatherService.Token = os.Getenv("TOKEN")

	if services.WeatherService.Token == "" {
		panic("no WEATHER API token in weather service")
	}

	http.HandleFunc("/weather", environment.MiddlewareJSON(handlers.Weather))

	fmt.Println("Serve the app")
	http.ListenAndServe(":4200", nil)
}
