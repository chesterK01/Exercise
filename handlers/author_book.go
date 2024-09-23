package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthorBookHandler struct {
	IAuthorBookService services.IAuthorBookService
}

// API to create a new Author-Book relationship using Gin
func (_self AuthorBookHandler) CreateAuthorBook(c *gin.Context) {
	var authorBook models.Author_Book
	// Decode the incoming JSON request body into the authorBook model
	if err := c.ShouldBindJSON(&authorBook); err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid input")
		return
	}

	// Call the service to create the relationship
	err := _self.IAuthorBookService.CreateAuthorBook(&authorBook)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{"message": "Author-Book relationship created successfully"})
}

// API to get Books by authorName using Gin
func (_self AuthorBookHandler) GetBooksByAuthorName(c *gin.Context) {
	// Get the author name from the URL query params
	authorName := c.Query("author_name")
	if authorName == "" {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Missing author name")
		return
	}

	// Call the service to get books by author name
	books, err := _self.IAuthorBookService.GetBooksByAuthorName(authorName)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// If no books were found
	if len(books) == 0 {
		utils.ReturnErrorJSON(c.Writer, http.StatusNotFound, "No books found for the given author name")
		return
	}

	// Success response with the list of books
	c.JSON(http.StatusOK, books)
}

// API to get all author-book relationship using Gin
func (_self AuthorBookHandler) GetAllAuthorBookRelationships(c *gin.Context) {
	// Call the service layer method to get all relationships
	relationships, err := _self.IAuthorBookService.GetAllAuthorBookRelationships()
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	// Send back the relationships as JSON
	c.JSON(http.StatusOK, relationships)
}

// API to get Author-Book relationship by Book ID using Gin
func (_self AuthorBookHandler) GetAuthorBookByBookID(c *gin.Context) {
	// Get the book ID from the URL query params (change "book_id" to "id")
	bookIDStr := c.Query("id")
	if bookIDStr == "" {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Missing book ID")
		return
	}

	// Convert book ID from string to integer
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid book ID")
		return
	}

	// Call the service to get the author-book relationship by book ID
	authorBook, err := _self.IAuthorBookService.GetAuthorBookByBookID(bookID)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, "Error fetching author-book relationship")
		return
	}

	// If the relationship was not found
	if authorBook == nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusNotFound, "Author-Book relationship not found")
		return
	}

	// Success response with the author-book relationship
	c.JSON(http.StatusOK, authorBook)
}
