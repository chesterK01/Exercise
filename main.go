package main

import (
	"Exercise1/db"
	"Exercise1/handlers"
	middlewares "Exercise1/middleware"
	"Exercise1/repositories"
	"Exercise1/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	// Khởi tạo kết nối cơ sở dữ liệu
	database := db.InitDB()
	defer database.Close()

	// Tạo các repository và service
	authorRepo := repositories.AuthorRepository{DB: database}
	authorService := services.AuthorService{IAuthorRepo: authorRepo}
	authorHandler := handlers.AuthorHandler{IAuthorService: authorService}

	bookRepo := repositories.BookRepository{DB: database}
	bookService := services.BookService{IBookRepo: bookRepo}
	bookHandler := handlers.BookHandler{IBookService: bookService}
	// Khởi tạo JWT Middleware
	authMiddleware, err := middlewares.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// Khởi tạo router
	router := gin.Default()

	// Định nghĩa route /login để đăng nhập và lấy token
	router.POST("/login", authMiddleware.LoginHandler)

	// Nhóm các route yêu cầu xác thực
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// Các route quản lý authors, yêu cầu quyền admin
		auth.POST("/create-author", authorHandler.CreateAuthor)
		auth.GET("/authors", authorHandler.GetAuthors)
		auth.GET("/author", authorHandler.GetAuthorByID)
	}

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/create-book", bookHandler.CreateBook)
		auth.GET("/books", bookHandler.GetBooks)
		auth.GET("/book", bookHandler.GetBookByID)
	}

	// Route mặc định
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Author API!",
		})
	})

	// Cấu hình cổng chạy server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Chạy server
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Không thể chạy server: %v", err)
	}
}
