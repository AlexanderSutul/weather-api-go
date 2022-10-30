package handlers

import (
	"fmt"
	"net/http"
	"runtime"
	"weather-api-go/models"
)

func GetMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("%d MB", m.Alloc/uint64(1024))
}

func Health(w http.ResponseWriter, req *http.Request) {
	resp := models.ApiResponse{
		Data: models.Health{
			Cpus:   runtime.NumCPU(),
			Memory: GetMemUsage(),
		},
	}
	resp.SendResponse(w, http.StatusOK)
}
