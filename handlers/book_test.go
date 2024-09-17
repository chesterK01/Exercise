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

func TestBookHandler_CreateBook(t *testing.T) {
	type mockCreateBookService struct {
		input *models.Book
		id    int64
		err   error
	}

	testCases := []struct {
		name                  string
		requestBody           interface{}
		expectedResponseBody  string
		expectedStatus        int
		mockCreateBookService *mockCreateBookService
	}{
		{
			name: "Invalid request body",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody: `{"error":"Invalid input"}` + "\n",
			expectedStatus:       http.StatusInternalServerError,
			mockCreateBookService: &mockCreateBookService{
				input: &models.Book{Name: "Math"},
				id:    0,
				err:   errors.New("failed to create book"),
			},
		},
		{
			name: "Successful book creation",
			requestBody: map[string]interface{}{
				"name": "Arthur",
			},
			expectedResponseBody: `{"id":1,"message":"Book created successfully"}` + "\n", // Bao gồm cả message
			expectedStatus:       http.StatusCreated,
			mockCreateBookService: &mockCreateBookService{
				input: &models.Book{Name: "Arthur"},
				id:    1,
				err:   nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockBookService := new(mockBookService)

			if testCase.mockCreateBookService != nil {
				mockBookService.On("CreateBook", testCase.mockCreateBookService.input).
					Return(testCase.mockCreateBookService.id, testCase.mockCreateBookService.err)
			}

			bookHandler := BookHandler{
				IBookService: mockBookService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBody))
			require.NoError(t, err)

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(bookHandler.CreateBook)
			handler.ServeHTTP(responseRecorder, req)

			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
