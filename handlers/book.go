package handlers

import (
	"Exercise1/db"
	"Exercise1/models"
	"Exercise1/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

// API tạo mới Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	query := "INSERT INTO book (name) VALUES (?)"
	result, err := db.DB.Exec(query, book.Name)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Book created successfully", "id": id})
}

// API để lấy danh sách Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Lấy tham số "limit" từ query string, nếu không có sẽ mặc định là 10
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}

	query := "SELECT id, name FROM book LIMIT ?"
	rows, err := db.DB.Query(query, limit)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name)
		if err != nil {
			utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

// API lấy Book theo ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Lấy book_id từ URL
	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Book ID not found")
		return
	}

	var book models.Book
	query := "SELECT id, name FROM book WHERE id = ?"
	err := db.DB.QueryRow(query, bookID).Scan(&book.ID, &book.Name)
	if err == sql.ErrNoRows {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "Book not found")
		return
	} else if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(book)
}
