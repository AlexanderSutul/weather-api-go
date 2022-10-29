package models

import (
	"errors"
	"fmt"
	"log"
	"time"
	"weather-api-go/constants"
)

type Record struct {
	War        *WeatherApiResponse
	LastUpdate string
}

type Database = map[Coords]Record

type DB struct {
	db Database
}

func InitDatabase() *DB {
	m := make(Database)
	return &DB{db: m}
}

func (db *DB) Fetch(c Coords) (*WeatherApiResponse, error) {
	db.printDBValues()

	if w, ok := db.db[c]; ok {
		fmt.Println("data from database")
		return w.War, nil
	}
	return nil, errors.New(fmt.Errorf("no value with coords %v", c).Error())
}

func (db *DB) Add(c Coords, war *WeatherApiResponse) {
	db.db[c] = Record{War: war, LastUpdate: time.Now().Format(constants.TIME_FORMAT)}
}

func (db *DB) printDBValues() {
	if len(db.db) > 0 {
		log.Println("current db state:")
		for k, v := range db.db {
			log.Printf("#%v: %v \n", k, v)
		}
	}
}

func isLastUpdateValid(t string) bool {
	if t == "" {
		return false
	}

	parsedTime, err := time.Parse(constants.TIME_FORMAT, t)
	if err != nil {
		return true
	}

	now := time.Now()

	return now.Sub(parsedTime).Milliseconds() > 0
}
