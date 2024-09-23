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

// Testcase for CreateBook with Gin
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

			// Create Gin's router
			router := gin.Default()

			// Create BookHandler with mockBookService
			bookHandler := BookHandler{
				IBookService: mockBookService,
			}

			// Register endpoint for router
			router.POST("/books", bookHandler.CreateBook)

			// Convert request body to JSON
			requestBody, err := json.Marshal(testCase.requestBody)
			require.NoError(t, err)

			// Create POST request with Gin's router
			req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBody))
			require.NoError(t, err)

			// Create response recorder
			responseRecorder := httptest.NewRecorder()

			// Call handler via Gin's router
			router.ServeHTTP(responseRecorder, req)

			// Check status code
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)

			// Check response body
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())

			// Verify that the mock service is called correctly
			if testCase.mockCreateBookInput != nil {
				mockBookService.AssertCalled(t, "CreateBook", testCase.mockCreateBookInput)
			}
		})
	}
}
