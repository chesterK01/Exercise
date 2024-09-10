package main

import (
	"Exercise1/db"
	"Exercise1/handlers"
	"Exercise1/repositories"
	"Exercise1/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Khởi tạo kết nối cơ sở dữ liệu
	dbConn := db.InitDB()

	// Tạo các repository
	authorRepo := repositories.AuthorRepository{DB: dbConn}
	bookRepo := repositories.BookRepository{DB: dbConn}
	authorBookRepo := repositories.AuthorBookRepository{DB: dbConn}

	// Tạo các service
	authorService := services.AuthorService{AuthorRepo: &authorRepo}
	bookService := services.BookService{BookRepo: &bookRepo}
	authorBookService := services.AuthorBookService{AuthorBookRepo: &authorBookRepo}

	// Tạo các handler
	authorHandler := handlers.AuthorHandler{IAuthorService: &authorService}
	bookHandler := handlers.BookHandler{IBookService: &bookService}
	authorBookHandler := handlers.AuthorBookHandler{IAuthorBookService: &authorBookService}

	// Định tuyến các API
	http.HandleFunc("/author", authorHandler.CreateAuthor)     // Tạo Author
	http.HandleFunc("/authors", authorHandler.GetAuthors)      // Lấy danh sách Authors
	http.HandleFunc("/author/id", authorHandler.GetAuthorByID) // Lấy Author theo ID

	http.HandleFunc("/book", bookHandler.CreateBook)     // Tạo Book
	http.HandleFunc("/books", bookHandler.GetBooks)      // Lấy danh sách Books
	http.HandleFunc("/book/id", bookHandler.GetBookByID) // Lấy Book theo ID

	http.HandleFunc("/author_book", authorBookHandler.CreateAuthorBook)      // Tạo mối quan hệ Author-Book
	http.HandleFunc("/books/author", authorBookHandler.GetBooksByAuthorName) // Lấy Book theo Author_name

	// Khởi chạy server
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
