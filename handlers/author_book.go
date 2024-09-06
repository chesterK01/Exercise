package handlers

import (
	"Exercise1/db"
	"Exercise1/models"
	"Exercise1/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// API tạo mới mối quan hệ Author-Book
func CreateAuthorBook(w http.ResponseWriter, r *http.Request) {
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

	// Kiểm tra author_id và book_id
	var authorExists bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM author WHERE id=?)", authorBook.AuthorID).Scan(&authorExists)
	if err != nil || !authorExists {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Author ID does not exist")
		return
	}

	var bookExists bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM book WHERE id=?)", authorBook.BookID).Scan(&bookExists)
	if err != nil || !bookExists {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Book ID does not exist")
		return
	}

	query := "INSERT INTO author_book (author_id, book_id) VALUES (?, ?)"
	_, err = db.DB.Exec(query, authorBook.AuthorID, authorBook.BookID)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Author-Book relationship created successfully"})
}

// API lấy danh sách Books theo tên tác giả (tìm kiếm gần đúng)
func GetBooksByAuthorName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Lấy author_name từ query parameters
	authorName := r.URL.Query().Get("author_name")
	if authorName == "" {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Missing author name")
		return
	}

	// Tìm kiếm sách dựa trên tên tác giả (gần đúng)
	query := `
		SELECT b.id, b.name
		FROM book b
		JOIN author_book ab ON b.id = ab.book_id
		JOIN author a ON ab.author_id = a.id
		WHERE a.name LIKE ?
	`
	rows, err := db.DB.Query(query, fmt.Sprintf("%%%s%%", authorName))
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, "Error querying books")
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name)
		if err != nil {
			utils.ReturnErrorJSON(w, http.StatusInternalServerError, "Error scanning book data")
			return
		}
		books = append(books, book)
	}

	if len(books) == 0 {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "No books found for the given author name")
		return
	}

	// Trả về danh sách sách dưới dạng JSON
	json.NewEncoder(w).Encode(books)
}
