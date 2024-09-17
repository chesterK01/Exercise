package repositories

import (
	"Exercise1/models"
	"database/sql"
	"errors"
)

type IStackRepository interface {
	UpdateBookStock(bookID int, stock int) error
	UpdateBookQuality(bookID int, quality string) error
	GetAllBooks() ([]models.Stack, error)
}

type StackRepository struct {
	DB *sql.DB
}

// Add stock for a specific book
func (_self StackRepository) UpdateBookStock(bookID int, stock int) error {
	result, err := _self.DB.Exec("UPDATE stack SET stock = ? WHERE id = ?", stock, bookID)
	if err != nil {
		return errors.New("failed to update stock")
	}

	// Kiểm tra số hàng bị ảnh hưởng
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows")
	}

	if rowsAffected == 0 {
		return errors.New("book not found") // Trả về lỗi nếu không có bản ghi nào bị cập nhật
	}

	return nil
}

// Save book quality
func (_self StackRepository) UpdateBookQuality(bookID int, quality string) error {
	result, err := _self.DB.Exec("UPDATE stack SET quality = ? WHERE id = ?", quality, bookID)
	if err != nil {
		return errors.New("failed to update quality")
	}

	// Kiểm tra số hàng bị ảnh hưởng
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows")
	}

	if rowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

// Get list of all books
func (_self StackRepository) GetAllBooks() ([]models.Stack, error) {
	rows, err := _self.DB.Query("SELECT id, title, stock, quality FROM stack")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stacks := []models.Stack{}
	for rows.Next() {
		var stack models.Stack
		if err := rows.Scan(&stack.ID, &stack.Title, &stack.Stock, &stack.Quality); err != nil {
			return nil, err
		}
		stacks = append(stacks, stack)
	}

	return stacks, nil
}
