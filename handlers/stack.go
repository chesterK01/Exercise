package handlers

import (
	"Exercise1/models"
	"Exercise1/services"
	"Exercise1/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type StackHandler struct {
	IAuthorBookService services.IAuthorBookService
	IStackService      services.IStackService
}

// API to create a new stock and quality for a book
func (_self StackHandler) CreateBookStockQuality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Lấy book ID từ query parameter id
	bookIDStr := r.URL.Query().Get("id")
	if bookIDStr == "" {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Missing book ID")
		return
	}

	// Chuyển đổi bookID từ chuỗi sang số nguyên
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	// Decode body để lấy thông tin stock và quality
	var data models.UpdateStockQualityRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Lấy thông tin Author_Book từ AuthorBookService
	authorBook, err := _self.IAuthorBookService.GetAuthorBookByBookID(bookID)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, "Error fetching author-book relationship")
		return
	}
	if authorBook == nil {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "Book not found")
		return
	}

	// Gọi StackService để tạo stock và quality cho book
	err = _self.IStackService.CreateBookStockQuality(bookID, data.Stock, data.Quality)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Trả về JSON với thông tin cập nhật
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Stock and quality created successfully",
		"bookTitle":  authorBook.BookName,
		"authorID":   authorBook.AuthorID,
		"authorName": authorBook.AuthorName,
		"newStock":   data.Stock,
		"newQuality": data.Quality,
	})
}
