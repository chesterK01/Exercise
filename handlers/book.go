package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type BookHandler struct {
	IBookService services.IBookService
}

// API to create a new Book
func (_self BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := _self.IBookService.CreateBook(&book)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Book created successfully", "id": id})
}

// API to get all Books
func (_self BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // If there is no limit, the default is 10
	}

	books, err := _self.IBookService.GetBooks(limit)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(books)
}

// API to get Book by bookID
func (_self BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	bookIDStr := r.URL.Query().Get("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := _self.IBookService.GetBookByID(bookID)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if book == nil {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "Book not found")
		return
	}

	// Return Book data as JSON
	json.NewEncoder(w).Encode(book)
}
