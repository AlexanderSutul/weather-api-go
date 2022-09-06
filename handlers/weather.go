package handlers

import (
	"net/http"
	"weather-api-go/models"
	"weather-api-go/services"
)

func Weather(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()
	lat := queries.Get("lat")
	resp := &models.Response{}
	if lat == "" {
		resp.Error = "lat is not provided as query parameter"
		resp.SendResponse(w, http.StatusBadRequest)
		return
	}

	lon := queries.Get("lon")
	if lon == "" {
		resp.Error = "lon is not provided as query parameter"
		resp.SendResponse(w, http.StatusBadRequest)
		return
	}

	weatherApiResponse, err := services.WeatherService.GetWeaterExternalApi(lat, lon)
	if err != nil {
		resp.Error = err.Error()
		resp.SendResponse(w, http.StatusInternalServerError)
		return
	}

	resp.Data = models.InitWeatherResponse(weatherApiResponse)
	resp.SendResponse(w, http.StatusOK)
}
