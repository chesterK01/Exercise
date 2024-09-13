package handlers

import (
	"Exercise1/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
			expectedResponseBody:  `{"error":"Invalid input"}` + "\n", // Định dạng JSON
			expectedStatus:        http.StatusBadRequest,
			mockCreateAuthorInput: nil,
		},
		{
			name: "Author creation failed",
			requestBody: map[string]interface{}{
				"name": "Arthur",
			},
			expectedResponseBody: `{"error":"failed to create author"}` + "\n", // Định dạng JSON
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
			expectedResponseBody: `{"id":1,"message":"Author created successfully"}` + "\n", // Bao gồm cả message
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
			mockAuthorService := new(mockAuthorService)

			if testCase.mockCreateAuthorInput != nil {
				mockAuthorService.On("CreateAuthor", testCase.mockCreateAuthorInput.input).
					Return(testCase.mockCreateAuthorInput.id, testCase.mockCreateAuthorInput.err)
			}

			authorHandler := AuthorHandler{
				IAuthorService: mockAuthorService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/authors", bytes.NewBuffer(requestBody))
			require.NoError(t, err)

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(authorHandler.CreateAuthor)
			handler.ServeHTTP(responseRecorder, req)

			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
