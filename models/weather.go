package models

import "weather-api-go/utils"

type WeatherApiResponseCoordinates struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}
type WeatherApiResponseMain struct {
	FeelsLike   float32 `json:"feels_like"`
	Temperature float32 `json:"temp"`
	Humidity    int     `json:"humidity"`
}

type WeatherApiResponseSys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherApiResponse struct {
	Coordinates WeatherApiResponseCoordinates `json:"coord"`
	Main        WeatherApiResponseMain        `json:"main"`
	Name        string                        `json:"name"`
	Sys         WeatherApiResponseSys         `json:"sys"`
}

type WeatherResponseCoordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type WeatherResponse struct {
	Coordinates *WeatherResponseCoordinates `json:"coordinates"`
	FeelsLike   float32                     `json:"feelsLike"`
	Humidity    int                         `json:"humidity"`
	Temperature float32                     `json:"temperature"`
	City        string                      `json:"city"`
	Country     string                      `json:"country"`
	SunsetTime  string                      `json:"sunsetTime"`
	SunriseTime string                      `json:"sunriseTime"`
}

func InitWeatherResponse(war *WeatherApiResponse) *WeatherResponse {
	return &WeatherResponse{
		Coordinates: &WeatherResponseCoordinates{
			Latitude:  war.Coordinates.Latitude,
			Longitude: war.Coordinates.Longitude,
		},
		FeelsLike:   war.Main.FeelsLike,
		Humidity:    war.Main.Humidity,
		Temperature: war.Main.Temperature,
		City:        war.Name,
		Country:     war.Sys.Country,
		SunsetTime:  utils.GetTime(war.Sys.Sunset),
		SunriseTime: utils.GetTime(war.Sys.Sunrise),
	}
}
