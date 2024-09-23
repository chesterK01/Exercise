package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	IBookService services.IBookService
}

// API to create a new Book
func (_self BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := _self.IBookService.CreateBook(&book)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "id": id})
}

// API to get all Books
func (_self BookHandler) GetBooks(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	books, err := _self.IBookService.GetBooks(limit)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

// API to get Book by bookID
func (_self BookHandler) GetBookByID(c *gin.Context) {
	bookIDStr := c.Query("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := _self.IBookService.GetBookByID(bookID)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if book == nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusNotFound, "Book not found")
		return
	}

	c.JSON(http.StatusOK, book)
}
