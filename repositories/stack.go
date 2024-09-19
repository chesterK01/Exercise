package repositories

import (
	"database/sql"
	"errors"
)

type IStackRepository interface {
	CreateBookStockQuality(bookID int, stock int, quality string) error
}

type StackRepository struct {
	DB *sql.DB
}

// Hàm chèn stock và quality vào bảng stack
func (_self StackRepository) CreateBookStockQuality(bookID int, stock int, quality string) error {
	// Câu truy vấn để chèn dữ liệu stock và quality
	query := `INSERT INTO stack (book_id, stock, quality) VALUES (?, ?, ?)`

	// Thực hiện truy vấn
	result, err := _self.DB.Exec(query, bookID, stock, quality)
	if err != nil {
		return errors.New("failed to insert stock and quality") // Trả về lỗi nếu có vấn đề khi thực hiện truy vấn
	}

	// Kiểm tra số hàng bị ảnh hưởng
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to retrieve affected rows")
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, insertion failed")
	}

	return nil
}
