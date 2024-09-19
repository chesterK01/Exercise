package services

import (
	"Exercise1/models"
	"Exercise1/repositories"
	"fmt"
	"log"
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

// Create a new author-book relationship
func (_self AuthorBookService) CreateAuthorBook(authorBook *models.Author_Book) error {
	log.Printf("Creating a new author-book relationship: %+v", authorBook)
	return _self.IAuthorBookRepo.CreateAuthorBook(authorBook)
}

// Get books by author name
func (_self AuthorBookService) GetBooksByAuthorName(authorName string) ([]models.Book, error) {
	log.Printf("Fetching books for author: %s", authorName)
	books, err := _self.IAuthorBookRepo.GetBooksByAuthorName(authorName)
	if err != nil {
		log.Printf("Error fetching books by author name: %v", err)
		return nil, fmt.Errorf("error fetching books by author name: %v", err)
	}
	return books, nil
}

// Get all author-book relationships
func (_self AuthorBookService) GetAllAuthorBookRelationships() ([]models.Author_Book, error) {
	log.Println("Fetching all author-book relationships")
	relationships, err := _self.IAuthorBookRepo.GetAllAuthorBookRelationships()
	if err != nil {
		log.Printf("Error fetching author-book relationships: %v", err)
		return nil, fmt.Errorf("error fetching author-book relationships: %v", err)
	}
	return relationships, nil
}

// Get author-book relationship by bookID
func (_self AuthorBookService) GetAuthorBookByBookID(bookID int) (*models.Author_Book, error) {
	log.Printf("Fetching author-book relationship for bookID: %d", bookID)
	authorBook, err := _self.IAuthorBookRepo.GetAuthorBookByBookID(bookID)
	if err != nil {
		log.Printf("Error fetching author-book relationship: %v", err)
		return nil, fmt.Errorf("error fetching author-book relationship: %v", err)
	}
	if authorBook == nil {
		log.Printf("No author-book relationship found for bookID: %d", bookID)
		return nil, fmt.Errorf("no author-book relationship found for bookID: %d", bookID)
	}
	return authorBook, nil
}
