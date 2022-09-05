package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weather-api-go/models"
	"weather-api-go/utils"
)

type weatherService struct {
	Token string
}

var WeatherService *weatherService = &weatherService{}

func (ws *weatherService) GetWeaterExternalApi(lat, lon string) (*models.WeatherApiResponse, error) {
	url := fmt.Sprintf(
		utils.API_URL_TEMPLATE,
		lat,
		lon,
		ws.Token,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	wap := &models.WeatherApiResponse{}
	err = json.Unmarshal(body, wap)
	if err != nil {
		return nil, err
	}

	return wap, nil
}
