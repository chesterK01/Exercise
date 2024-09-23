package handlers

import (
	"Exercise1/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock CreateBook method
func (m *mockBookService) CreateBook(book *models.Book) (int64, error) {
	args := m.Called(book)
	return args.Get(0).(int64), args.Error(1)
}

func TestBookHandler_CreateBook(t *testing.T) {
	testCases := []struct {
		name                 string
		requestBody          interface{}
		expectedResponseBody string
		expectedStatus       int
		mockCreateBookInput  *models.Book
		mockCreateBookID     int64
		mockCreateBookErr    error
	}{
		{
			name: "Invalid request body",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody: `{"error":"Invalid input"}` + "\n",
			expectedStatus:       http.StatusBadRequest,
			mockCreateBookInput:  nil,
		},
		{
			name: "Successful book creation",
			requestBody: map[string]interface{}{
				"name": "Arthur",
			},
			expectedResponseBody: `{"id":1,"message":"Book created successfully"}` + "\n",
			expectedStatus:       http.StatusCreated,
			mockCreateBookInput:  &models.Book{Name: "Arthur"},
			mockCreateBookID:     1,
			mockCreateBookErr:    nil,
		},
		{
			name: "Book creation failed",
			requestBody: map[string]interface{}{
				"name": "Arthur",
			},
			expectedResponseBody: `{"error":"failed to create book"}` + "\n",
			expectedStatus:       http.StatusInternalServerError,
			mockCreateBookInput:  &models.Book{Name: "Arthur"},
			mockCreateBookID:     0,
			mockCreateBookErr:    errors.New("failed to create book"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocking BookService
			mockBookService := new(mockBookService)

			if testCase.mockCreateBookInput != nil {
				mockBookService.On("CreateBook", testCase.mockCreateBookInput).
					Return(testCase.mockCreateBookID, testCase.mockCreateBookErr)
			}

			// Tạo router của Gin
			router := gin.Default()

			// Tạo BookHandler với mockBookService
			bookHandler := BookHandler{
				IBookService: mockBookService,
			}

			// Đăng ký endpoint cho router
			router.POST("/books", bookHandler.CreateBook)

			// Chuyển request body thành JSON
			requestBody, err := json.Marshal(testCase.requestBody)
			require.NoError(t, err)

			// Tạo request POST với router của Gin
			req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBody))
			require.NoError(t, err)

			// Tạo response recorder
			responseRecorder := httptest.NewRecorder()

			// Gọi handler qua router của Gin
			router.ServeHTTP(responseRecorder, req)

			// Kiểm tra status code
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)

			// Kiểm tra nội dung response body
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())

			// Xác thực rằng mock service được gọi chính xác
			if testCase.mockCreateBookInput != nil {
				mockBookService.AssertCalled(t, "CreateBook", testCase.mockCreateBookInput)
			}
		})
	}
}
