package util

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func WriteErrorMessage(w http.ResponseWriter, statusCode int, err string) {
	WriteJsonResponse(w, statusCode, ErrorResponse{ErrorMessage: err})
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, response any) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
