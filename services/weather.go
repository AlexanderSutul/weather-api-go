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

func (ws *weatherService) GetWeaterExternalApi(coords models.Coords) (*models.WeatherApiResponse, error) {
	url := fmt.Sprintf(
		utils.API_URL_TEMPLATE,
		coords.Lat,
		coords.Lon,
		ws.Token,
	)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
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
