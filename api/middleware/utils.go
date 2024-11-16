package middleware

import (
	"encoding/json"
	"net/http"
)

type httpEror struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BadRequest(w http.ResponseWriter, err error) {
	HttpJSONError(w, err.Error(), http.StatusBadGateway)
}

func InternalError(w http.ResponseWriter, err error) {
	HttpJSONError(w, err.Error(), http.StatusInternalServerError)
}

func UnauthenticatedError(w http.ResponseWriter, err error) {
	HttpJSONError(w, err.Error(), http.StatusUnauthorized)
}

func HttpJSONError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(httpEror{
		Code:    code,
		Message: error,
	})
}
