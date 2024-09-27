package main

import (
	"Exercise1/db"
	"Exercise1/routes"
	"log"
	"os"
)

func main() {
	// Khởi tạo kết nối cơ sở dữ liệu
	database := db.InitDB()
	defer database.Close()

	// Khởi tạo các route
	router, err := routes.SetupRoutes(database)
	if err != nil {
		log.Fatal("Error setting up routes: ", err)
	}

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
