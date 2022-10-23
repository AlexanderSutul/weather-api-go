package utils

import (
	"fmt"
	"log"
	"time"
	"weather-api-go/constants"
)

func ShowUptimeMessage() {
	start := time.Now()
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		fmt.Printf("Current seconds uptime: %.0f\n", time.Since(start).Seconds())
	}

	fmt.Println("no more messages will be displayed")
}

func GetTime(original int64) string {
	tm := time.Unix(original, 0)
	return tm.UTC().String()
}

func IsInvalidLastUpdate(t string) bool {
	if t == "" {
		return true
	}
	cacheTime, err := time.Parse(constants.TIME_FORMAT, t)
	if err != nil {
		log.Println(err)
		return true
	}
	now := time.Now()
	return now.Sub(cacheTime).Minutes() > 60
}
