package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type AuthorBookHandler struct {
	IAuthorBookService services.IAuthorBookService
}

// API to create a new Author-Book relationship
func (_self AuthorBookHandler) CreateAuthorBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Check if the method is POST
	if r.Method != "POST" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return // Add return after method not allowed
	}

	var authorBook models.Author_Book
	// Decode the incoming JSON request body into the authorBook model
	err := json.NewDecoder(r.Body).Decode(&authorBook)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Call the service to create the relationship
	err = _self.IAuthorBookService.CreateAuthorBook(&authorBook)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Author-Book relationship created successfully"})
}

// API to get Books by authorName
func (_self AuthorBookHandler) GetBooksByAuthorName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Check if the method is GET
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get the author name from the URL query params
	authorName := r.URL.Query().Get("author_name")
	if authorName == "" {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Missing author name")
		return
	}

	// Call the service to get books by author name
	books, err := _self.IAuthorBookService.GetBooksByAuthorName(authorName)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// If no books were found
	if len(books) == 0 {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "No books found for the given author name")
		return
	}

	// Success response with the list of books
	json.NewEncoder(w).Encode(books)
}

// API to get all author-book relationship
func (_self AuthorBookHandler) GetAllAuthorBookRelationships(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Call the service layer method to get all relationships
	relationships, err := _self.IAuthorBookService.GetAllAuthorBookRelationships()
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Send back the relationships as JSON
	json.NewEncoder(w).Encode(relationships)
}

// API to get Author-Book relationship by Book ID
func (_self AuthorBookHandler) GetAuthorBookByBookID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if the method is GET
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get the book ID from the URL query params (change "book_id" to "id")
	bookIDStr := r.URL.Query().Get("id")
	if bookIDStr == "" {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Missing book ID")
		return
	}

	// Convert book ID from string to integer
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	// Call the service to get the author-book relationship by book ID
	authorBook, err := _self.IAuthorBookService.GetAuthorBookByBookID(bookID)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, "Error fetching author-book relationship")
		return
	}

	// If the relationship was not found
	if authorBook == nil {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "Author-Book relationship not found")
		return
	}

	// Success response with the author-book relationship
	json.NewEncoder(w).Encode(authorBook)
}
