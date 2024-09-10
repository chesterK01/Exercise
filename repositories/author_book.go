package repositories

import (
	"Exercise1/models"
	"database/sql"
)

type IAuthorBookRepository interface {
	CreateAuthorBook(*models.Author_Book) error
	GetBooksByAuthorName(authorName string) ([]models.Book, error)
}

type AuthorBookRepository struct {
	DB *sql.DB
}

// Create a new author-book relationship
func (_self AuthorBookRepository) CreateAuthorBook(authorBook *models.Author_Book) error {
	query := `INSERT INTO author_book (author_id, book_id) VALUES (?, ?)`
	_, err := _self.DB.Exec(query, authorBook.AuthorID, authorBook.BookID)
	return err
}

// Get books by author's name
func (_self AuthorBookRepository) GetBooksByAuthorName(authorName string) ([]models.Book, error) {
	query := `
		SELECT b.id, b.name
		FROM book b
		JOIN author_book ab ON b.id = ab.book_id
		JOIN author a ON ab.author_id = a.id
		WHERE a.name LIKE ?
	`
	rows, err := _self.DB.Query(query, "%"+authorName+"%")
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
