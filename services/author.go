package services

import (
	"Exercise1/models"
	"Exercise1/repositories"
)

type IAuthorService interface {
	CreateAuthor(*models.Author) (int64, error)
	GetAuthors(limit int) ([]models.Author, error)
	GetAuthorByID(authorID int) (*models.Author, error)
}

type AuthorService struct {
	AuthorRepo repositories.IAuthorRepository
}

func (_self AuthorService) CreateAuthor(author *models.Author) (int64, error) {
	return _self.AuthorRepo.CreateAuthor(author)
}

func (_self AuthorService) GetAuthors(limit int) ([]models.Author, error) {
	return _self.AuthorRepo.GetAuthors(limit)
}

func (_self AuthorService) GetAuthorByID(authorID int) (*models.Author, error) {
	author, err := _self.AuthorRepo.GetAuthorByID(authorID)
	if err != nil {
		return nil, err
	}
	return author, nil
}
