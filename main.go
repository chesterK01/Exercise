package main

import (
	"Exercise1/db"
	"Exercise1/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Khởi tạo kết nối cơ sở dữ liệu
	db.InitDB()

	// Định tuyến các API
	http.HandleFunc("/author", handlers.CreateAuthor)
	http.HandleFunc("/book", handlers.CreateBook)
	http.HandleFunc("/author_book", handlers.CreateAuthorBook)
	http.HandleFunc("/authors", handlers.GetAuthors)
	http.HandleFunc("/books", handlers.GetBooks)
	http.HandleFunc("/author/id", handlers.GetAuthorByID)
	http.HandleFunc("/book/id", handlers.GetBookByID)
	http.HandleFunc("/books/author", handlers.GetBooksByAuthorName)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
