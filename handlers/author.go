package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	IAuthorService services.IAuthorService
}

// API to create a new Author
func (_self AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := _self.IAuthorService.CreateAuthor(&author)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Author created successfully", "id": id})
}

// API to get all Authors
func (_self AuthorHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	authors, err := _self.IAuthorService.GetAuthors(limit)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(authors)
}

// API to get Author by authorID
func (_self AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get author_id from query parameters
	authorIDStr := r.URL.Query().Get("id")
	authorID, err := strconv.Atoi(authorIDStr)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid Author ID")
		return
	}

	// Call service to get Author's information
	author, err := _self.IAuthorService.GetAuthorByID(authorID)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if author == nil {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "Author not found")
		return
	}

	// Return Author data as JSON
	json.NewEncoder(w).Encode(author)
}
