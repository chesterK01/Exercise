package services

import "Exercise1/repositories"

type IStackService interface {
	CreateBookStockQuality(bookID int, stock int, quality string) error
}

type StackService struct {
	IStackRepo repositories.IStackRepository
}

// Create stock and quality for book
func (_self StackService) CreateBookStockQuality(bookID int, stock int, quality string) error {
	// Call the repository's CreateBookStockQuality function
	err := _self.IStackRepo.CreateBookStockQuality(bookID, stock, quality)
	if err != nil {
		return err
	}

	return nil
}
