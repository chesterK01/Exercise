package handlers

import (
	"Exercise1/db"
	"Exercise1/models"
	"Exercise1/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

// API tạo mới Author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
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

	query := "INSERT INTO author (name) VALUES (?)"
	result, err := db.DB.Exec(query, author.Name)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Author created successfully", "id": id})

}

// API để lấy danh sách Authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
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

	query := "SELECT id, name FROM author LIMIT ?"
	rows, err := db.DB.Query(query, limit)
	if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.Name)
		if err != nil {
			utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		authors = append(authors, author)
	}

	json.NewEncoder(w).Encode(authors)
}

// API lấy Author theo ID
func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		utils.ReturnErrorJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Lấy author_id từ URL
	authorID := r.URL.Query().Get("id")
	if authorID == "" {
		utils.ReturnErrorJSON(w, http.StatusBadRequest, "Author ID not found")
		return
	}

	var author models.Author
	query := "SELECT id, name FROM author WHERE id = ?"
	err := db.DB.QueryRow(query, authorID).Scan(&author.ID, &author.Name)
	if err == sql.ErrNoRows {
		utils.ReturnErrorJSON(w, http.StatusNotFound, "Author not found")
		return
	} else if err != nil {
		utils.ReturnErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(author)
}
