package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(resp, req)

	if http.StatusOK != resp.Code {
		t.Errorf("expected code %v, actual code %v", http.StatusOK, resp.Code)
	}

	body := resp.Body.String()

	if body == "" {
		t.Error("cannot parse response body")
	}
}
