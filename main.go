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

	token := os.Getenv("TOKEN")
	if token == "" {
		panic("WEATHER API token is not provided")
	}
	port := os.Getenv("PORT")
	if port == "" {
		panic("port is not provided")
	}

	go utils.ShowUptimeMessage()

	services.WeatherService.Token = token

	http.HandleFunc("/weather", environment.MiddlewareJSON(handlers.Weather))

	fmt.Println("Serve the app")
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
