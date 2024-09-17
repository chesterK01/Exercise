package handlers

import (
	"Exercise1/models"
	"github.com/stretchr/testify/mock"
)

type mockBookService struct {
	mock.Mock
}

func (_self mockBookService) CreateBook(book *models.Book) (int64, error) {
	args := _self.Called(book)
	return args.Get(0).(int64), args.Error(1)
}

func (_self mockBookService) GetBooks(limit int) ([]models.Book, error) {
	args := _self.Called(limit)
	var books []models.Book
	if args.Get(0) == nil {
		books = args.Get(1).([]models.Book)
	}
	return books, args.Error(1)
}

func (_self mockBookService) GetBookByID(bookID int) (*models.Book, error) {
	args := _self.Called(bookID)
	var book *models.Book
	if args.Get(0) != nil {
		book = args.Get(0).(*models.Book)
	}
	return book, args.Error(1)
}
