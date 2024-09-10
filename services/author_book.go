package services

import (
	"Exercise1/models"
	"Exercise1/repositories"
)

type IAuthorBookService interface {
	CreateAuthorBook(*models.Author_Book) error
	GetBooksByAuthorName(authorName string) ([]models.Book, error)
}

type AuthorBookService struct {
	AuthorBookRepo repositories.IAuthorBookRepository
}

func (_self AuthorBookService) CreateAuthorBook(authorBook *models.Author_Book) error {
	return _self.AuthorBookRepo.CreateAuthorBook(authorBook)
}

func (_self AuthorBookService) GetBooksByAuthorName(authorName string) ([]models.Book, error) {
	return _self.AuthorBookRepo.GetBooksByAuthorName(authorName)
}
