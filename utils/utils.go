package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-api-go/models"
)

func SendResponse(w http.ResponseWriter, responseCode int, response *models.Response) {
	buffer, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = "cannot parse response model"
		fmt.Fprintln(w, response)
		return
	}

	w.WriteHeader(responseCode)
	fmt.Fprintln(w, string(buffer))
}
