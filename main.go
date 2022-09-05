package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"weather-api-go/environment"
	"weather-api-go/handlers"
	"weather-api-go/services"
)

func main() {

	err := environment.LoadEnvs()
	if err != nil {
		panic(err)
	}

	go showUptimeMessage()

	apiKey := os.Getenv("TOKEN")
	if apiKey == "" {
		panic("no WEATHER API token")
	}

	services.WeatherService.Token = apiKey

	http.HandleFunc("/weather", environment.MiddlewareJSON(handlers.Weather))

	fmt.Println("Serve the app")
	http.ListenAndServe(":4200", nil)
}

func showUptimeMessage() {
	start := time.Now()
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		fmt.Printf("Current seconds uptime: %.0f\n", time.Since(start).Seconds())
	}
}
