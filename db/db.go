package db

import "weather-api-go/models"

var DatabaseInstance models.DB = *models.InitDatabase()
