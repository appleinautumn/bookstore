package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"gotu/bookstore/internal/service/mocks"
	"gotu/bookstore/internal/types"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Response struct {
	Data *types.User `json:"data"`
}

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	os.Exit(m.Run())
}

func TestSignUp(t *testing.T) {
	bookService := mocks.NewBookService(t)
	orderService := mocks.NewOrderService(t)
	userService := mocks.NewUserService(t)
	handler := NewApiHandler(bookService, userService, orderService)

	t.Run("success", func(t *testing.T) {
		// load location for New York
		location, err := time.LoadLocation("America/New_York")
		if err != nil {
			t.Fatalf("Error loading location: %v", err)
		}

		// mock user
		user := &types.User{
			ID:        77,
			Email:     faker.Email(),
			Name:      faker.Name(),
			Password:  faker.Password(),
			CreatedAt: time.Date(2024, time.January, 1, 10, 30, 0, 0, location),
			UpdatedAt: time.Date(2024, time.January, 1, 10, 30, 0, 0, location),
		}

		// set payload
		payload := map[string]interface{}{
			"email":    user.Email,
			"name":     user.Name,
			"password": user.Password,
		}
		payloadJson, _ := json.Marshal(payload)

		// mock CreateUser
		userService.On("CreateUser", mock.Anything, mock.Anything).Return(user, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(payloadJson))
		res := httptest.NewRecorder()

		// list
		handler.SignUp(res, req)
		defer res.Result().Body.Close()

		if status := res.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// get body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading body: %v", err)
		}

		// Create a variable of the wrapper struct type
		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		user2 := response.Data

		assert.NoError(t, err)
		assert.Equal(t, res.Code, http.StatusOK)
		assert.Equal(t, user.ID, user2.ID)
		assert.Equal(t, user.Email, user2.Email)
		assert.Equal(t, user.Password, user2.Password)
	})

	t.Run("error - decoding payload", func(t *testing.T) {
		// set invalid payload
		payload := "abc"
		payloadJson, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(payloadJson))
		res := httptest.NewRecorder()

		// list
		handler.SignUp(res, req)
		defer res.Result().Body.Close()

		// assert
		assert.Equal(t, res.Code, http.StatusBadRequest)
	})

	t.Run("error - missing payloads", func(t *testing.T) {
		// set invalid payload
		tests := map[string]struct {
			params     map[string]interface{}
			statusCode int
		}{
			"missing email": {
				map[string]interface{}{
					"name":     "成龍",
					"password": "123456",
				},
				http.StatusBadRequest,
			},
			"missing name": {
				map[string]interface{}{
					"email":    "cheng@long.com",
					"password": "123456",
				},
				http.StatusBadRequest,
			},
			"missing password": {
				map[string]interface{}{
					"email": "cheng@long.com",
					"name":  "成龍",
				},
				http.StatusBadRequest,
			},
			"without params": {
				map[string]interface{}{},
				http.StatusBadRequest,
			},
		}

		for _, tt := range tests {
			payloadJson, _ := json.Marshal(tt.params)

			req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(payloadJson))
			res := httptest.NewRecorder()

			// list
			handler.SignUp(res, req)
			defer res.Result().Body.Close()

			// assert
			assert.Equal(t, res.Code, tt.statusCode)
		}
	})

	t.Run("error - create error", func(t *testing.T) {
		// mock CreateUser to throw error
		userService.On("CreateUser", mock.Anything, mock.Anything).Return(nil, errors.New("anything")).Once()

		// set payload
		payload := map[string]interface{}{
			"email":    "cheng@long.com",
			"name":     "成龍",
			"password": "123456",
		}
		payloadJson, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(payloadJson))
		res := httptest.NewRecorder()

		// list
		handler.SignUp(res, req)
		defer res.Result().Body.Close()

		// assert
		assert.Equal(t, res.Code, http.StatusInternalServerError)
	})
}
