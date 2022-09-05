package models

import (
	"time"
)

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
	Base        string                        `json:"base"`
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
	Base        string                      `json:"base"`
	Coordinates *WeatherResponseCoordinates `json:"coordinates"`
	FeelsLike   float32                     `json:"feels_like"`
	Humidity    int                         `json:"humidity"`
	Temperature float32                     `json:"temperature"`
	City        string                      `json:"city"`
	Country     string                      `json:"country"`
	SunsetTime  string                      `json:"sunset_time"`
	SunriseTime string                      `json:"sunrise_time"`
}

func InitWeatherResponse(war *WeatherApiResponse) *WeatherResponse {
	return &WeatherResponse{
		Base: war.Base,
		Coordinates: &WeatherResponseCoordinates{
			Latitude:  war.Coordinates.Latitude,
			Longitude: war.Coordinates.Longitude,
		},
		FeelsLike:   war.Main.FeelsLike,
		Humidity:    war.Main.Humidity,
		Temperature: war.Main.Temperature,
		City:        war.Name,
		Country:     war.Sys.Country,
		SunsetTime:  a.getTime(war.Sys.Sunset),
		SunriseTime: a.getTime(war.Sys.Sunrise),
	}
}

type Adaper struct{}

var a Adaper

func (a Adaper) getTime(original int64) string {
	tm := time.Unix(original, 0)
	return tm.UTC().String()
}
