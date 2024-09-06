package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Khởi tạo kết nối cơ sở dữ liệu
func InitDB() {
	var err error
	dsn := "root:ngoctuan1072003@tcp(localhost:3306)/library"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Kết nối thành công đến MySQL!")
}
