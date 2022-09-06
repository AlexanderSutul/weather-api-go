package utils

import (
	"fmt"
	"time"
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
