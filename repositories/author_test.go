package repositories

import (
	"Exercise1/db"
	"Exercise1/models"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test hàm CreateAuthor
func TestAuthorRepository_CreateAuthor(t *testing.T) {
	testCases := []struct {
		name        string
		input       *models.Author
		expectedID  int64
		expectedErr error
		mockDB      *sql.DB
	}{
		{
			name: "Create author successfully",
			input: &models.Author{
				Name: "J.K. Rowling",
			},
			expectedID:  1, // Giả sử ID của author là 1 khi được thêm thành công
			expectedErr: nil,
			mockDB:      db.InitDB(), // Kết nối cơ sở dữ liệu thật
		},
		{
			name: "Create author failed due to DB error",
			input: &models.Author{
				Name: "George Orwell",
			},
			expectedID:  0,
			expectedErr: errors.New("Error connecting to the database: invalid connection"),
			mockDB:      nil, // Giả lập lỗi kết nối CSDL
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given: Tạo một repository với database tương ứng
			authorRepo := AuthorRepository{
				DB: testCase.mockDB,
			}

			// When: Gọi hàm CreateAuthor
			id, err := authorRepo.CreateAuthor(testCase.input)

			// Then: Kiểm tra kết quả trả về
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedID, id)
			}
		})
	}
}

// Test hàm GetAuthors
func TestAuthorRepository_GetAuthors(t *testing.T) {
	testCases := []struct {
		name         string
		limit        int
		expectedData []models.Author
		expectedErr  error
		mockDB       *sql.DB
	}{
		{
			name:  "Get authors successfully",
			limit: 2,
			expectedData: []models.Author{
				{ID: 1, Name: "J.K. Rowling"},
				{ID: 2, Name: "George Orwell"},
			},
			expectedErr: nil,
			mockDB:      db.InitDB(),
		},
		{
			name:        "Get authors failed due to DB error",
			limit:       1,
			expectedErr: errors.New("Error connecting to the database: invalid connection"),
			mockDB:      nil, // Giả lập lỗi kết nối
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given: Khởi tạo AuthorRepository với mockDB
			authorRepo := AuthorRepository{
				DB: testCase.mockDB,
			}

			// When: Gọi hàm GetAuthors
			authors, err := authorRepo.GetAuthors(testCase.limit)

			// Then: Kiểm tra kết quả trả về
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedData, authors)
			}
		})
	}
}

// Test hàm GetAuthorByID
func TestAuthorRepository_GetAuthorByID(t *testing.T) {
	testCases := []struct {
		name         string
		input        int
		expectedData *models.Author
		expectedErr  error
		mockDB       *sql.DB
	}{
		{
			name:  "Get author by ID successfully",
			input: 1, // Giả sử ID của author là 1
			expectedData: &models.Author{
				ID:   1,
				Name: "J.K. Rowling",
			},
			expectedErr: nil,
			mockDB:      db.InitDB(),
		},
		{
			name:         "Author not found",
			input:        99, // ID không tồn tại
			expectedData: nil,
			expectedErr:  nil,
			mockDB:       db.InitDB(),
		},
		{
			name:        "Get author by ID failed due to DB error",
			input:       1,
			expectedErr: errors.New("Error connecting to the database: invalid connection"),
			mockDB:      nil, // Giả lập lỗi kết nối
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given: Khởi tạo AuthorRepository với mockDB
			authorRepo := AuthorRepository{
				DB: testCase.mockDB,
			}

			// When: Gọi hàm GetAuthorByID
			author, err := authorRepo.GetAuthorByID(testCase.input)

			// Then: Kiểm tra kết quả trả về
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedData, author)
			}
		})
	}
}
