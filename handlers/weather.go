package handlers

import (
	"net/http"
	"weather-api-go/constants"
	"weather-api-go/db"
	"weather-api-go/models"
	"weather-api-go/services"
	"weather-api-go/utils"
)

func Weather(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()
	resp := &models.ApiResponse{}

	lat := queries.Get("lat")
	if lat == "" {
		resp.Error = "lat is not provided as a query parameter"
		resp.SendResponse(w, http.StatusBadRequest)
		return
	}

	if !utils.IsNumber(lat) {
		resp.Error = "lat has to be a number"
		resp.SendResponse(w, http.StatusBadRequest)
		return
	}

	lon := queries.Get("lon")
	if lon == "" {
		resp.Error = "lon is not provided as a query parameter"
		resp.SendResponse(w, http.StatusBadRequest)
		return
	}

	if !utils.IsNumber(lon) {
		resp.Error = "lon has to be a number"
		resp.SendResponse(w, http.StatusBadRequest)
		return
	}

	weather, err := getWeather(w, models.Coords{Lat: lat, Lon: lon})
	if err != nil {
		resp.Error = err.Error()
		resp.SendResponse(w, http.StatusInternalServerError)
		return
	}

	resp.Data = models.InitWeatherResponse(weather)
	resp.SendResponse(w, http.StatusOK)
}

func getWeather(w http.ResponseWriter, coords models.Coords) (*models.WeatherApiResponse, error) {
	war, err := db.DatabaseInstance.Fetch(coords)
	if err != nil {
		weatherApiResponse, err := services.WeatherService.GetWeather(coords)
		if err != nil {
			return nil, err
		}
		db.DatabaseInstance.Add(coords, weatherApiResponse)
		w.Header().Add(constants.WEATHER_HEADER, "false")
		return weatherApiResponse, nil
	}
	w.Header().Add(constants.WEATHER_HEADER, "true")
	return war, nil
}
