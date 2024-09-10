package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() *sql.DB {
	// Chuỗi kết nối đến MySQL
	dsn := "root:ngoctuan1072003@tcp(localhost:3306)/library"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database")

	// Lưu kết nối vào biến toàn cục
	DB = db
	return db
}
