package routes

import (
	"Exercise1/handlers"
	middleware "Exercise1/middleware"
	"Exercise1/repositories"
	"Exercise1/services"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) (*gin.Engine, error) {
	router := gin.Default()

	// Create repository and service
	authorRepo := repositories.AuthorRepository{DB: db}
	authorService := services.AuthorService{IAuthorRepo: authorRepo}
	authorHandler := handlers.AuthorHandler{IAuthorService: authorService}

	bookRepo := repositories.BookRepository{DB: db}
	bookService := services.BookService{IBookRepo: bookRepo}
	bookHandler := handlers.BookHandler{IBookService: bookService}

	authorbookRepo := repositories.AuthorBookRepository{DB: db}
	authorbookService := services.AuthorBookService{IAuthorBookRepo: authorbookRepo}
	authorBookHandler := handlers.AuthorBookHandler{IAuthorBookService: authorbookService}

	stackRepo := repositories.StackRepository{DB: db}
	stackService := services.StackService{IStackRepo: stackRepo}
	stackHandler := handlers.StackHandler{
		IAuthorBookService: authorbookService,
		IStackService:      stackService,
	}
	// Initialize JWT Middleware
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		return nil, err
	}

	// Define route /login to login and get token
	router.POST("/login", authMiddleware.LoginHandler)

	// Group routes that require authentication
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// Author management routes
		auth.POST("/create-author", authorHandler.CreateAuthor)
		auth.GET("/authors", authorHandler.GetAuthors)
		auth.GET("/author", authorHandler.GetAuthorByID)

		// Book management routes
		auth.POST("/create-book", bookHandler.CreateBook)
		auth.GET("/books", bookHandler.GetBooks)
		auth.GET("/book", bookHandler.GetBookByID)

		// Author-Book relationship management routes
		auth.POST("/create-author-book", authorBookHandler.CreateAuthorBook)
		auth.GET("/author-book/name", authorBookHandler.GetBooksByAuthorName)
		auth.GET("/author-books", authorBookHandler.GetAllAuthorBookRelationships)
		auth.GET("/author-book/id", authorBookHandler.GetAuthorBookByBookID)

		// Stack management routes
		auth.POST("/create-stack", stackHandler.CreateBookStockQuality)
	}

	return router, nil
}
