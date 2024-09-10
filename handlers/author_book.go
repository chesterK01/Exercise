package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"encoding/json"
	"net/http"
)

type AuthorBookHandler struct {
	IAuthorBookService services.IAuthorBookService
}

// API to create a new Author-Book relationship
func (_self AuthorBookHandler) CreateAuthorBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var authorBook models.Author_Book
	err := json.NewDecoder(r.Body).Decode(&authorBook)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	err = _self.IAuthorBookService.CreateAuthorBook(&authorBook)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Author-Book relationship created successfully"})
}

// API to get Books by authorName
func (_self AuthorBookHandler) GetBooksByAuthorName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	authorName := r.URL.Query().Get("author_name")
	if authorName == "" {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Missing author name")
		return
	}

	books, err := _self.IAuthorBookService.GetBooksByAuthorName(authorName)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(books) == 0 {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "No books found for the given author name")
		return
	}

	json.NewEncoder(w).Encode(books)
}
