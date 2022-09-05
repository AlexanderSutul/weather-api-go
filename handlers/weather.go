package handlers

import (
	"net/http"
	"weather-api-go/models"
	"weather-api-go/services"
	"weather-api-go/utils"
)

func Weather(w http.ResponseWriter, req *http.Request) {
	lat := req.URL.Query().Get("lat")
	if lat == "" {
		utils.SendResponse(w, http.StatusBadRequest, &models.Response{
			Error: "lat is not provided as query parameter",
		})
		return
	}

	lon := req.URL.Query().Get("lon")
	if lon == "" {
		utils.SendResponse(w, http.StatusBadRequest, &models.Response{
			Error: "lon is not provided as query parameter",
		})
		return
	}

	weatherApiResponse, err := services.WeatherService.GetWeaterExternalApi(lat, lon)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, &models.Response{
			Error: err.Error(),
		})
		return
	}

	weatherResponse := models.InitWeatherResponse(weatherApiResponse)

	utils.SendResponse(w, http.StatusOK, &models.Response{
		Data: weatherResponse,
	})
}
