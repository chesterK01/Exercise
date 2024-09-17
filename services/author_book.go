package services

import (
	"Exercise1/models"
	"Exercise1/repositories"
)

type IAuthorBookService interface {
	CreateAuthorBook(*models.Author_Book) error
	GetBooksByAuthorName(authorName string) ([]models.Book, error)
	GetAllAuthorBookRelationships() ([]models.Author_Book, error)
	GetAuthorBookByBookID(bookID int) (*models.Author_Book, error)
}

type AuthorBookService struct {
	IAuthorBookRepo repositories.IAuthorBookRepository
}

func (_self AuthorBookService) CreateAuthorBook(authorBook *models.Author_Book) error {
	return _self.IAuthorBookRepo.CreateAuthorBook(authorBook)
}

func (_self AuthorBookService) GetBooksByAuthorName(authorName string) ([]models.Book, error) {
	return _self.IAuthorBookRepo.GetBooksByAuthorName(authorName)
}

func (_self AuthorBookService) GetAllAuthorBookRelationships() ([]models.Author_Book, error) {
	return _self.IAuthorBookRepo.GetAllAuthorBookRelationships()
}
func (_self AuthorBookService) GetAuthorBookByBookID(bookID int) (*models.Author_Book, error) {
	return _self.IAuthorBookRepo.GetAuthorBookByBookID(bookID)
}
