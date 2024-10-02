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

// Function to insert stock and quality into stack table
func (_self StackRepository) CreateBookStockQuality(bookID int, stock int, quality string) error {
	// Query to insert stock and quality data
	query := `INSERT INTO stack (book_id, stock, quality) VALUES (?, ?, ?)`

	// Execute query
	result, err := _self.DB.Exec(query, bookID, stock, quality)
	if err != nil {
		return errors.New("failed to insert stock and quality") // Returns an error if there is a problem executing the query
	}

	// Check the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to retrieve affected rows")
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected, insertion failed")
	}

	return nil
}
