package utils

import (
	"encoding/json"
	"net/http"
)

// Structure for error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Return errors as JSON
func ReturnErrorJSON(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
