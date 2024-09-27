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

	// Tạo các repository và service
	authorRepo := repositories.AuthorRepository{DB: db}
	authorService := services.AuthorService{IAuthorRepo: authorRepo}
	authorHandler := handlers.AuthorHandler{IAuthorService: authorService}

	bookRepo := repositories.BookRepository{DB: db}
	bookService := services.BookService{IBookRepo: bookRepo}
	bookHandler := handlers.BookHandler{IBookService: bookService}

	// Khởi tạo JWT Middleware
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		return nil, err
	}

	// Định nghĩa route /login để đăng nhập và lấy token
	router.POST("/login", authMiddleware.LoginHandler)

	// Nhóm các route yêu cầu xác thực
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// Các route quản lý authors
		auth.POST("/create-author", authorHandler.CreateAuthor)
		auth.GET("/authors", authorHandler.GetAuthors)
		auth.GET("/author", authorHandler.GetAuthorByID)

		// Các route quản lý books
		auth.POST("/create-book", bookHandler.CreateBook)
		auth.GET("/books", bookHandler.GetBooks)
		auth.GET("/book", bookHandler.GetBookByID)
	}

	return router, nil
}
