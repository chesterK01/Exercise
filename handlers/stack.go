package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type StackHandler struct {
	IAuthorBookService services.IAuthorBookService
	IStackService      services.IStackService
}

// API to create a new stock and quality for a book
func (_self StackHandler) CreateBookStockQuality(c *gin.Context) {
	bookIDStr := c.Query("id")
	if bookIDStr == "" {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Missing book ID")
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var data models.UpdateStockQualityRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid input")
		return
	}

	authorBook, err := _self.IAuthorBookService.GetAuthorBookByBookID(bookID)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, "Error fetching author-book relationship")
		return
	}
	if authorBook == nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusNotFound, "Book not found")
		return
	}

	err = _self.IStackService.CreateBookStockQuality(bookID, data.Stock, data.Quality)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Stock and quality created successfully",
		"bookTitle":  authorBook.BookName,
		"authorID":   authorBook.AuthorID,
		"authorName": authorBook.AuthorName,
		"newStock":   data.Stock,
		"newQuality": data.Quality,
	})
}
