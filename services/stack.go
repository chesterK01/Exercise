package services

import (
	"Exercise1/models"
	"Exercise1/repositories"
)

type IStackService interface {
	UpdateBookStock(bookID int, stock int) error
	UpdateBookQuality(bookID int, quality string) error
	GetAllBooks() ([]models.Stack, error)
}

type StackService struct {
	StackRepo repositories.IStackRepository
}

func (_self *StackService) UpdateBookStock(bookID int, stock int) error {
	return _self.StackRepo.UpdateBookStock(bookID, stock)
}

func (_self *StackService) UpdateBookQuality(bookID int, quality string) error {
	return _self.StackRepo.UpdateBookQuality(bookID, quality)
}
func (_self *StackService) GetAllBooks() ([]models.Stack, error) {
	return _self.StackRepo.GetAllBooks()
}
