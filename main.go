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
	// Initialize database connection
	dbConn := db.InitDB()

	// Create repositories
	authorRepo := repositories.AuthorRepository{DB: dbConn}
	bookRepo := repositories.BookRepository{DB: dbConn}
	authorBookRepo := repositories.AuthorBookRepository{DB: dbConn}
	stackRepo := repositories.StackRepository{DB: dbConn}

	// Create services
	authorService := services.AuthorService{IAuthorRepo: &authorRepo}
	bookService := services.BookService{IBookRepo: &bookRepo}
	authorBookService := services.AuthorBookService{IAuthorBookRepo: &authorBookRepo}
	stackService := services.StackService{IStackRepo: &stackRepo}

	// Create handlers
	authorHandler := handlers.AuthorHandler{IAuthorService: &authorService}
	bookHandler := handlers.BookHandler{IBookService: &bookService}
	authorBookHandler := handlers.AuthorBookHandler{IAuthorBookService: &authorBookService}
	stackHandler := handlers.StackHandler{IStackService: &stackService}

	// Routing API
	http.HandleFunc("/author", authorHandler.CreateAuthor)     // Create a new Author
	http.HandleFunc("/authors", authorHandler.GetAuthors)      // Get all Authors
	http.HandleFunc("/author/id", authorHandler.GetAuthorByID) // Get Author by authorID
	http.HandleFunc("/stock/update-stock", stackHandler.UpdateBookStock)

	http.HandleFunc("/book", bookHandler.CreateBook)     // Create a new Book
	http.HandleFunc("/books", bookHandler.GetBooks)      // Get all Books
	http.HandleFunc("/book/id", bookHandler.GetBookByID) // Get Book by bookID
	http.HandleFunc("/stock/update-quality", stackHandler.UpdateBookQuality)

	http.HandleFunc("/author_book", authorBookHandler.CreateAuthorBook)                            // Create a new Author-Book relationship
	http.HandleFunc("/books/author", authorBookHandler.GetBooksByAuthorName)                       // Get Book by Author_name
	http.HandleFunc("/author_book/relationships", authorBookHandler.GetAllAuthorBookRelationships) // Get all author-book relationships
	http.HandleFunc("/stocks", stackHandler.GetAllBooks)

	// Start server
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
