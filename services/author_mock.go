package services

import (
	"Exercise1/models"
	"github.com/stretchr/testify/mock"
)

type mockAuthorRepository struct {
	mock.Mock
}

func (_self mockAuthorRepository) CreateAuthor(author *models.Author) (int64, error) {
	args := _self.Called(author)
	return args.Get(0).(int64), args.Error(1)
}

func (_self mockAuthorRepository) GetAuthors(limit int) ([]models.Author, error) {
	args := _self.Called(limit)
	var authors []models.Author
	if args.Get(0) != nil {
		authors = args.Get(0).([]models.Author)
	}
	return authors, args.Error(1)
}

func (_self mockAuthorRepository) GetAuthorByID(authorID int) (*models.Author, error) {
	args := _self.Called(authorID)
	var author *models.Author
	if args.Get(0) != nil {
		author = args.Get(0).(*models.Author)
	}
	return author, args.Error(1)
}
