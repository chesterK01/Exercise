package repositories

import (
	"Exercise1/models"
	"database/sql"
	"fmt"
)

type IAuthorBookRepository interface {
	CreateAuthorBook(*models.Author_Book) error
	GetBooksByAuthorName(authorName string) ([]models.Book, error)
	GetAllAuthorBookRelationships() ([]models.Author_Book, error)
	GetAuthorBookByBookID(bookID int) (*models.Author_Book, error)
}

type AuthorBookRepository struct {
	DB *sql.DB
}

// Create a new author-book relationship
func (_self AuthorBookRepository) CreateAuthorBook(authorBook *models.Author_Book) error {
	query := `INSERT INTO author_book (author_id, book_id) VALUES (?, ?)`
	_, err := _self.DB.Exec(query, authorBook.AuthorID, authorBook.BookID)
	if err != nil {
		return fmt.Errorf("error creating author-book relationship: %v", err)
	}
	return nil
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
		return nil, fmt.Errorf("error fetching books by author name: %v", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Name); err != nil {
			return nil, fmt.Errorf("error scanning book data: %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %v", err)
	}
	return books, nil
}

// Get all author-book relationships with author_name and book_name
func (_self AuthorBookRepository) GetAllAuthorBookRelationships() ([]models.Author_Book, error) {
	query := `
		SELECT a.id, a.name, b.id, b.name
		FROM author_book ab
		INNER JOIN author a ON ab.author_id = a.id
		INNER JOIN book b ON ab.book_id = b.id
	`

	rows, err := _self.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching author-book relationships: %v", err)
	}
	defer rows.Close()

	var relationships []models.Author_Book
	for rows.Next() {
		var authorBook models.Author_Book
		if err := rows.Scan(&authorBook.AuthorID, &authorBook.AuthorName, &authorBook.BookID, &authorBook.BookName); err != nil {
			return nil, fmt.Errorf("error scanning author-book data: %v", err)
		}
		relationships = append(relationships, authorBook)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %v", err)
	}

	return relationships, nil
}

// Get author-book relationship by bookID
func (_self AuthorBookRepository) GetAuthorBookByBookID(bookID int) (*models.Author_Book, error) {
	var authorBook models.Author_Book
	query := `
        SELECT a.id, a.name, b.id, b.name
        FROM author_book ab
        INNER JOIN author a ON ab.author_id = a.id
        INNER JOIN book b ON ab.book_id = b.id
        WHERE b.id = ?
    `
	err := _self.DB.QueryRow(query, bookID).Scan(
		&authorBook.AuthorID, &authorBook.AuthorName,
		&authorBook.BookID, &authorBook.BookName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &authorBook, nil
}
