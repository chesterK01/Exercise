package handlers

import (
	"Exercise1/middleware"
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	IAuthorService services.IAuthorService
}

// API to create new Author (only admin can access)
func (_self AuthorHandler) CreateAuthor(c *gin.Context) {
	// Apply middleware to check admin rights
	middlewares.RoleMiddleware("admin")(c)
	if c.IsAborted() {
		return
	}

	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil || author.Name == "" {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := _self.IAuthorService.CreateAuthor(&author)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Author created successfully", "id": id})
}

// API to get all authors
func (_self AuthorHandler) GetAuthors(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	authors, err := _self.IAuthorService.GetAuthors(limit)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, authors)
}

// API to get Author by ID (Only admin can access)
func (_self AuthorHandler) GetAuthorByID(c *gin.Context) {
	// Apply middleware to check admin rights
	middlewares.RoleMiddleware("admin")(c)
	if c.IsAborted() {
		return
	}

	authorIDStr := c.Query("id")
	authorID, err := strconv.Atoi(authorIDStr)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusBadRequest, "Invalid Author ID")
		return
	}

	author, err := _self.IAuthorService.GetAuthorByID(authorID)
	if err != nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}
	if author == nil {
		utils.ReturnErrorJSON(c.Writer, http.StatusNotFound, "Author not found")
		return
	}

	c.JSON(http.StatusOK, author)
}
