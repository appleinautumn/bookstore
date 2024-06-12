package handler

import (
	"bytes"
	"encoding/json"
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

		userService.On("CreateUser", mock.Anything, mock.Anything).Return(user, nil).Once()
		handler := NewApiHandler(bookService, userService, orderService)

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
}
