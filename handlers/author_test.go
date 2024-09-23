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

// Mock CreateAuthor method
func (m *mockAuthorService) CreateAuthor(author *models.Author) (int64, error) {
	args := m.Called(author)
	return args.Get(0).(int64), args.Error(1)
}

// Testcase for CreateAuthor with Gin
func TestAuthorHandler_CreateAuthor(t *testing.T) {
	type mockCreateAuthorService struct {
		input *models.Author
		id    int64
		err   error
	}

	testCases := []struct {
		name                  string
		requestBody           interface{}
		expectedResponseBody  string
		expectedStatus        int
		mockCreateAuthorInput *mockCreateAuthorService
	}{
		{
			name: "Invalid request body",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody:  `{"error":"Invalid input"}` + "\n",
			expectedStatus:        http.StatusBadRequest,
			mockCreateAuthorInput: nil,
		},
		{
			name: "Author creation failed",
			requestBody: map[string]interface{}{
				"name": "Arthur",
			},
			expectedResponseBody: `{"error":"failed to create author"}` + "\n",
			expectedStatus:       http.StatusInternalServerError,
			mockCreateAuthorInput: &mockCreateAuthorService{
				input: &models.Author{Name: "Arthur"},
				id:    0,
				err:   errors.New("failed to create author"),
			},
		},
		{
			name: "Successful author creation",
			requestBody: map[string]interface{}{
				"name": "Arthur",
			},
			expectedResponseBody: `{"message":"Author created successfully","id":1}` + "\n",
			expectedStatus:       http.StatusCreated,
			mockCreateAuthorInput: &mockCreateAuthorService{
				input: &models.Author{Name: "Arthur"},
				id:    1,
				err:   nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Mocking AuthorService
			mockAuthorService := new(mockAuthorService)

			if testCase.mockCreateAuthorInput != nil {
				mockAuthorService.On("CreateAuthor", testCase.mockCreateAuthorInput.input).
					Return(testCase.mockCreateAuthorInput.id, testCase.mockCreateAuthorInput.err)
			}

			// Create Gin's router
			router := gin.Default()

			// Create AuthorHandler with mockAuthorService
			authorHandler := AuthorHandler{
				IAuthorService: mockAuthorService,
			}

			// Register endpoint for router
			router.POST("/authors", authorHandler.CreateAuthor)

			// Convert request body to JSON
			requestBody, err := json.Marshal(testCase.requestBody)
			require.NoError(t, err)

			// Create request POST with Gin's router
			req, err := http.NewRequest(http.MethodPost, "/authors", bytes.NewBuffer(requestBody))
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
			if testCase.mockCreateAuthorInput != nil {
				mockAuthorService.AssertCalled(t, "CreateAuthor", testCase.mockCreateAuthorInput.input)
			}
		})
	}
}
