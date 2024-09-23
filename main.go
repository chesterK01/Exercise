package main

import (
	"Exercise1/db"
	"Exercise1/handlers"
	"Exercise1/repositories"
	"Exercise1/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	dbConn := db.InitDB()

	// Initialize Gin router
	r := gin.Default()

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
	stackHandler := handlers.StackHandler{
		IAuthorBookService: authorBookService,
		IStackService:      stackService,
	}

	// Use Gin for routing API
	r.POST("/author", authorHandler.CreateAuthor)    // Create a new Author
	r.GET("/authors", authorHandler.GetAuthors)      // Get all Authors
	r.GET("/author/id", authorHandler.GetAuthorByID) // Get Author by authorID

	r.POST("/book", bookHandler.CreateBook)    // Create a new Book
	r.GET("/books", bookHandler.GetBooks)      // Get all Books
	r.GET("/book/id", bookHandler.GetBookByID) // Get Book by bookID

	r.POST("/author-book", authorBookHandler.CreateAuthorBook)
	r.POST("/author-books", authorBookHandler.GetAllAuthorBookRelationships)
	r.POST("/author-book/id", authorBookHandler.GetAuthorBookByBookID)
	r.POST("/author-book/name", authorBookHandler.GetBooksByAuthorName)

	r.POST("/stack/create", stackHandler.CreateBookStockQuality) // Create a new stock and quality for a book

	// Start the server
	r.Run(":8080") // Default runs on localhost:8080
}
