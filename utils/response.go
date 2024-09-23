package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Structure for error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Return errors as JSON
func ReturnErrorJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(gin.H{"error": message})
}
