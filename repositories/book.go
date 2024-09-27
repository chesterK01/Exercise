package repositories

import (
	"Exercise1/models"
	"database/sql"
)

type IBookRepository interface {
	CreateBook(*models.Book) (int64, error)
	GetBooks(limit int) ([]models.Book, error)
	GetBookByID(bookID int) (*models.Book, error)
}

type BookRepository struct {
	DB *sql.DB
}

// Create a new Book
func (_self BookRepository) CreateBook(book *models.Book) (int64, error) {
	query := `INSERT INTO book (name) VALUES ?`
	result, err := _self.DB.Exec(query, book.Name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Get all Books
func (_self BookRepository) GetBooks(limit int) ([]models.Book, error) {
	query := `SELECT id, name FROM book LIMIT ?`
	rows, err := _self.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Name); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// Get Book by bookID
func (_self BookRepository) GetBookByID(bookID int) (*models.Book, error) {
	var book models.Book
	query := `SELECT id, name FROM book WHERE id = ?`
	err := _self.DB.QueryRow(query, bookID).Scan(&book.ID, &book.Name)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
