package services

import "Exercise1/repositories"

type IStackService interface {
	CreateBookStockQuality(bookID int, stock int, quality string) error
}

type StackService struct {
	IStackRepo repositories.IStackRepository
}

// tạo stock và quality cho book
func (_self StackService) CreateBookStockQuality(bookID int, stock int, quality string) error {
	// Gọi hàm CreateBookStockQuality của repository
	err := _self.IStackRepo.CreateBookStockQuality(bookID, stock, quality)
	if err != nil {
		return err
	}

	return nil
}
