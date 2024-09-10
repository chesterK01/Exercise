package repositories

import (
	"Exercise1/models"
	"database/sql"
)

type IAuthorRepository interface {
	CreateAuthor(*models.Author) (int64, error)
	GetAuthors(limit int) ([]models.Author, error)
	GetAuthorByID(authorID int) (*models.Author, error)
}

type AuthorRepository struct {
	DB *sql.DB
}

func (_self AuthorRepository) CreateAuthor(author *models.Author) (int64, error) {
	query := "INSERT INTO author (name) VALUES (?)"
	result, err := _self.DB.Exec(query, author.Name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (_self AuthorRepository) GetAuthors(limit int) ([]models.Author, error) {
	query := "SELECT id, name FROM author LIMIT ?"
	rows, err := _self.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (_self AuthorRepository) GetAuthorByID(authorID int) (*models.Author, error) {
	var author models.Author
	query := "SELECT id, name FROM author WHERE id = ?"
	err := _self.DB.QueryRow(query, authorID).Scan(&author.ID, &author.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &author, nil
}
