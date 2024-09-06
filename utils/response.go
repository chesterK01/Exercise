package utils

import (
	"encoding/json"
	"net/http"
)

// Cấu trúc để phản hồi lỗi
type ErrorResponse struct {
	Error string `json:"error"`
}

// Trả về lỗi dưới dạng JSON
func ReturnErrorJSON(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
