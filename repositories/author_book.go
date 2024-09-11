package repositories

import (
	"Exercise1/models"
	"database/sql"
)

type IAuthorBookRepository interface {
	CreateAuthorBook(*models.Author_Book) error
	GetBooksByAuthorName(authorName string) ([]models.Book, error)
	GetAllAuthorBookRelationships() ([]models.Author_Book, error)
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

// Get all relationships with author_name and book_name
func (_self *AuthorBookRepository) GetAllAuthorBookRelationships() ([]models.Author_Book, error) {
	// Query to join author, book, and author_book tables
	query := `
		SELECT a.id, a.name, b.id, b.name
		FROM author_book ab
		INNER JOIN author a ON ab.author_id = a.id
		INNER JOIN book b ON ab.book_id = b.id`

	rows, err := _self.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relationships []models.Author_Book
	for rows.Next() {
		var authorBook models.Author_Book
		if err := rows.Scan(&authorBook.AuthorID, &authorBook.AuthorName, &authorBook.BookID, &authorBook.BookName); err != nil {
			return nil, err
		}
		relationships = append(relationships, authorBook)
	}

	return relationships, nil
}
