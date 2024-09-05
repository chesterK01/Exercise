package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// DB kết nối toàn cục
var db *sql.DB

// Cấu trúc Author
type Author struct {
	Name string `json:"name"`
}

// Cấu trúc Book
type Book struct {
	Name string `json:"name"`
}

// Cấu trúc AuthorBook cho mối quan hệ giữa tác giả và sách
type AuthorBook struct {
	AuthorID int `json:"author_id"`
	BookID   int `json:"book_id"`
}

// Kết nối tới cơ sở dữ liệu
func initDB() {
	var err error
	dsn := "root:ngoctuan1072003@tcp(localhost:3306)/library"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Kết nối thành công đến MySQL!")
}

// API tạo mới Author
func createAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var author Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO author (name) VALUES (?)"
	result, err := db.Exec(query, author.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Author created successfully with ID: %d", id)
}

// API tạo mới Book
func createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO book (name) VALUES (?)"
	result, err := db.Exec(query, book.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Book created successfully with ID: %d", id)
}

// API tạo mới mối quan hệ AuthorBook
func createAuthorBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var authorBook AuthorBook
	err := json.NewDecoder(r.Body).Decode(&authorBook)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Kiểm tra author_id có tồn tại không
	var authorExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM author WHERE id=?)", authorBook.AuthorID).Scan(&authorExists)
	if err != nil || !authorExists {
		http.Error(w, "Author ID does not exist", http.StatusBadRequest)
		return
	}

	// Kiểm tra book_id có tồn tại không
	var bookExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM book WHERE id=?)", authorBook.BookID).Scan(&bookExists)
	if err != nil || !bookExists {
		http.Error(w, "Book ID does not exist", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO author_book (author_id, book_id) VALUES (?, ?)"
	_, err = db.Exec(query, authorBook.AuthorID, authorBook.BookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Author-Book relationship created successfully!")
}

func main() {
	// Khởi tạo kết nối cơ sở dữ liệu
	initDB()

	// Định tuyến các API
	http.HandleFunc("/author", createAuthor)          // API tạo Author
	http.HandleFunc("/book", createBook)              // API tạo Book
	http.HandleFunc("/author_book", createAuthorBook) // API tạo mối quan hệ AuthorBook

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
