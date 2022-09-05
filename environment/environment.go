package environment

import (
	"errors"
	"net/http"

	"github.com/joho/godotenv"
)

func LoadEnvs() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("cannot find file .env")
	}
	return nil
}

func MiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
