package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (resp *ApiResponse) SendResponse(w http.ResponseWriter, responseCode int) {
	buffer, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Error = err.Error()
		fmt.Fprintln(w, resp)
		return
	}

	w.WriteHeader(responseCode)
	fmt.Fprintln(w, string(buffer))
}
