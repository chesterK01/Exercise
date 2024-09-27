package main

import (
	"Exercise1/db"
	"Exercise1/routes"
	"log"
	"os"
)

func main() {
	// Initialize database connection
	database := db.InitDB()
	defer database.Close()

	// Initialize routes
	router, err := routes.SetupRoutes(database)
	if err != nil {
		log.Fatal("Error setting up routes: ", err)
	}

	// Configure server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Không thể chạy server: %v", err)
	}
}
