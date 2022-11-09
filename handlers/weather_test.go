package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWether(t *testing.T) {
	cases := []struct {
		RequestUrl   string
		ResponseCode int
		ErrorDescr   string
	}{
		{RequestUrl: "/weather", ResponseCode: http.StatusBadRequest, ErrorDescr: ErrLatIsNotProvided},
		{RequestUrl: "/weather?lat=10qwe", ResponseCode: http.StatusBadRequest, ErrorDescr: ErrLatIsNotNumber},
		{RequestUrl: "/weather?lat=10", ResponseCode: http.StatusBadRequest, ErrorDescr: ErrLonIsNotProvided},
		{RequestUrl: "/weather?lat=10&lon=10qwe", ResponseCode: http.StatusBadRequest, ErrorDescr: ErrLonIsNotNumber},
		{RequestUrl: "/weather?lat=10&lon=10", ResponseCode: http.StatusOK, ErrorDescr: ""},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("current test case %d", i+1), func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, c.RequestUrl, nil)
			res := httptest.NewRecorder()

			Weather(res, req)

			if res.Code != c.ResponseCode {
				t.Errorf("expected response code %d, actual response code %d", c.ResponseCode, res.Code)
			}

			body := res.Body.String()
			fmt.Println(body)
			if !strings.Contains(body, c.ErrorDescr) {
				t.Errorf("expected error: %s \n actual error: %s", c.ErrorDescr, body)
			}
		})
	}
}
