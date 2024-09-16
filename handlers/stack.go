package handlers

import (
	"Exercise1/services"
	"Exercise1/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type StackHandler struct {
	IStackService services.IStackService
}

// API 1: Nhập số lượng tồn kho từng sách
func (_self StackHandler) UpdateBookStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "PUT" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	bookIDStr := r.URL.Query().Get("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var data struct {
		Stock int `json:"stock"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	err = _self.IStackService.UpdateBookStock(bookID, data.Stock)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Stock updated successfully"})
}

// API 2: Lưu chất lượng sách
func (_self StackHandler) UpdateBookQuality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "PUT" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	bookIDStr := r.URL.Query().Get("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var data struct {
		Quality string `json:"quality"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	err = _self.IStackService.UpdateBookQuality(bookID, data.Quality)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Quality updated successfully"})
}

// API 3: Lấy danh sách sách
func (_self StackHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	stacks, err := _self.IStackService.GetAllBooks()
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(stacks)
}
