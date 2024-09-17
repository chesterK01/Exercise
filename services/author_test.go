package services

import (
	"errors"
	"testing"

	"Exercise1/models"
	"github.com/stretchr/testify/require"
)

// Test CreateAuthor method
func TestAuthorService_CreateAuthor(t *testing.T) {
	testCases := []struct {
		name           string
		input          *models.Author
		expectedID     int64
		expectedErr    error
		mockRepoResult int64
		mockRepoError  error
	}{
		{
			name: "Create author successfully",
			input: &models.Author{
				Name: "J.K. Rowling",
			},
			expectedID:     1,
			expectedErr:    nil,
			mockRepoResult: 1,
			mockRepoError:  nil,
		},
		{
			name: "Create author failed",
			input: &models.Author{
				Name: "George Orwell",
			},
			expectedID:     0,
			expectedErr:    errors.New("create author failed"),
			mockRepoResult: 0,
			mockRepoError:  errors.New("create author failed"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockAuthorRepo := new(mockAuthorRepository)
			mockAuthorRepo.On("CreateAuthor", testCase.input).Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := AuthorService{
				IAuthorRepo: mockAuthorRepo,
			}

			// When
			id, err := service.CreateAuthor(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedID, id)
			}
		})
	}
}

// Test GetAuthors method
func TestAuthorService_GetAuthors(t *testing.T) {
	testCases := []struct {
		name           string
		limit          int
		expectedData   []models.Author
		expectedErr    error
		mockRepoResult []models.Author
		mockRepoError  error
	}{
		{
			name:  "Get authors successfully",
			limit: 2,
			expectedData: []models.Author{
				{ID: 1, Name: "J.K. Rowling"},
				{ID: 2, Name: "George Orwell"},
			},
			expectedErr: nil,
			mockRepoResult: []models.Author{
				{ID: 1, Name: "J.K. Rowling"},
				{ID: 2, Name: "George Orwell"},
			},
			mockRepoError: nil,
		},
		{
			name:           "Get authors failed",
			limit:          2,
			expectedErr:    errors.New("get authors failed"),
			mockRepoResult: nil,
			mockRepoError:  errors.New("get authors failed"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockAuthorRepo := new(mockAuthorRepository)
			mockAuthorRepo.On("GetAuthors", testCase.limit).Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := AuthorService{
				IAuthorRepo: mockAuthorRepo,
			}

			// When
			authors, err := service.GetAuthors(testCase.limit)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedData, authors)
			}
		})
	}
}

// Test GetAuthorByID method
func TestAuthorService_GetAuthorByID(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedData   *models.Author
		expectedErr    error
		mockRepoResult *models.Author
		mockRepoError  error
	}{
		{
			name:  "Get author by ID successfully",
			input: 1,
			expectedData: &models.Author{
				ID:   1,
				Name: "J.K. Rowling",
			},
			expectedErr: nil,
			mockRepoResult: &models.Author{
				ID:   1,
				Name: "J.K. Rowling",
			},
			mockRepoError: nil,
		},
		{
			name:           "Get author by ID failed",
			input:          1,
			expectedErr:    errors.New("get author by ID failed"),
			mockRepoResult: nil,
			mockRepoError:  errors.New("get author by ID failed"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockAuthorRepo := new(mockAuthorRepository)
			mockAuthorRepo.On("GetAuthorByID", testCase.input).Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := AuthorService{
				IAuthorRepo: mockAuthorRepo,
			}

			// When
			author, err := service.GetAuthorByID(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedData, author)
			}
		})
	}
}
