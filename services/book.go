package services

import (
	"Exercise1/models"
	"Exercise1/repositories"
)

type IBookService interface {
	CreateBook(book *models.Book) (int64, error)
	GetBooks(limit int) ([]models.Book, error)
	GetBookByID(bookID int) (*models.Book, error)
}

type BookService struct {
	BookRepo repositories.IBookRepository
}

func (_self BookService) CreateBook(book *models.Book) (int64, error) {
	return _self.BookRepo.CreateBook(book)
}

func (_self BookService) GetBooks(limit int) ([]models.Book, error) {
	return _self.BookRepo.GetBooks(limit)
}

func (_self BookService) GetBookByID(bookID int) (*models.Book, error) {
	return _self.BookRepo.GetBookByID(bookID)
}
