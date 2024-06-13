package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/service/mocks"
	"gotu/bookstore/internal/types"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	os.Exit(m.Run())
}

func TestListBooks(t *testing.T) {
	bookService := mocks.NewBookService(t)
	orderService := mocks.NewOrderService(t)
	userService := mocks.NewUserService(t)
	handler := NewApiHandler(bookService, userService, orderService)

	t.Run("success", func(t *testing.T) {
		// mock book 1
		var b1 types.Book
		if err := faker.FakeData(&b1); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock book 2
		var b2 types.Book
		if err := faker.FakeData(&b2); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock books
		books := []*types.Book{
			&b1,
			&b2,
		}

		// mock List
		bookService.On("List", mock.Anything, mock.Anything).Return(books, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		res := httptest.NewRecorder()

		// list books
		handler.ListBooks(res, req)
		defer res.Result().Body.Close()

		// get body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading body: %v", err)
		}

		// Create a variable of the wrapper with literal struct
		var response struct {
			Data []*types.Book `json:"data"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		data := response.Data

		// assert
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, len(data), len(books))
	})

	t.Run("error - list error", func(t *testing.T) {
		// mock List to throw error
		bookService.On("List", mock.Anything, mock.Anything).Return(nil, errors.New("anything")).Once()

		req := httptest.NewRequest(http.MethodPost, "/books", nil)
		res := httptest.NewRecorder()

		// list books
		handler.ListBooks(res, req)
		defer res.Result().Body.Close()

		// assert
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
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

		// get body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading body: %v", err)
		}

		// Create a variable of the wrapper with literal struct
		var response struct {
			Data *types.User `json:"data"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		user2 := response.Data

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
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
		assert.Equal(t, http.StatusBadRequest, res.Code)
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
			assert.Equal(t, tt.statusCode, res.Code)
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
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestCreateOrder(t *testing.T) {
	bookService := mocks.NewBookService(t)
	orderService := mocks.NewOrderService(t)
	userService := mocks.NewUserService(t)
	handler := NewApiHandler(bookService, userService, orderService)

	t.Run("success", func(t *testing.T) {
		// set payload
		payload := map[string]interface{}{
			"orders": []map[string]interface{}{
				{
					"book_id":  1,
					"quantity": 3,
				}, {
					"book_id":  2,
					"quantity": 1,
				},
			},
		}

		// set user ID
		userID := 5

		// mock Order to return
		order1 := &types.Order{
			ID:     rand.Int64(),
			UserID: int64(userID),
		}

		payloadJson, _ := json.Marshal(payload)

		// mock CreateOrder
		orderService.On("CreateOrder", mock.Anything, mock.Anything).Return(order1, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(payloadJson))
		res := httptest.NewRecorder()

		// set user ID in header
		req.Header.Set("user_id", strconv.Itoa(userID))

		// call CreateOrder
		handler.CreateOrder(res, req)
		defer res.Result().Body.Close()

		// get body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading body: %v", err)
		}

		// Create a variable of the wrapper with literal struct
		var response struct {
			Data *types.Order `json:"data"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		orderRes := response.Data

		// assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, order1.ID, orderRes.ID)
		assert.Equal(t, order1.UserID, orderRes.UserID)
	})

	t.Run("error - converting userID in header", func(t *testing.T) {
		// mock order item request 1
		var item1 request.OrderItem
		if err := faker.FakeData(&item1); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock order item request 2
		var item2 request.OrderItem
		if err := faker.FakeData(&item2); err != nil {
			t.Errorf("err: %v", err)
		}

		// set payload
		payload := request.OrderRequest{
			Orders: []*request.OrderItem{
				&item1,
				&item2,
			},
		}

		payloadJson, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(payloadJson))
		res := httptest.NewRecorder()

		// set invalid user ID in header
		req.Header.Set("user_id", "abc")

		// call CreateOrder
		handler.CreateOrder(res, req)
		defer res.Result().Body.Close()

		// assert bad request
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("error - decoding payload", func(t *testing.T) {
		// set invalid payload
		payload := "abc"
		payloadJson, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(payloadJson))
		res := httptest.NewRecorder()

		// set user ID in header
		req.Header.Set("user_id", strconv.Itoa(5))

		// call CreateOrder
		handler.CreateOrder(res, req)
		defer res.Result().Body.Close()

		// assert http status bad request
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("error - missing payloads", func(t *testing.T) {
		// set user ID
		userID := 5

		// set invalid payload
		tests := map[string]struct {
			params     map[string]interface{}
			statusCode int
		}{
			"missing orders": {
				map[string]interface{}{
					"not_orders": "111",
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

			req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(payloadJson))
			res := httptest.NewRecorder()

			// set user ID in header
			req.Header.Set("user_id", strconv.Itoa(userID))

			// call CreateOrder
			handler.CreateOrder(res, req)
			defer res.Result().Body.Close()

			// assert
			assert.Equal(t, tt.statusCode, res.Code)
		}
	})

	t.Run("error - CreateOrder error", func(t *testing.T) {
		// set payload
		payload := map[string]interface{}{
			"orders": []map[string]interface{}{
				{
					"book_id":  1,
					"quantity": 3,
				}, {
					"book_id":  2,
					"quantity": 1,
				},
			},
		}

		payloadJson, _ := json.Marshal(payload)

		// mock CreateOrder
		orderService.On("CreateOrder", mock.Anything, mock.Anything).Return(nil, errors.New("anything")).Once()

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(payloadJson))
		res := httptest.NewRecorder()

		// set user ID in header
		req.Header.Set("user_id", strconv.Itoa(5))

		// call CreateOrder
		handler.CreateOrder(res, req)
		defer res.Result().Body.Close()

		// assert
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListOrders(t *testing.T) {
	bookService := mocks.NewBookService(t)
	orderService := mocks.NewOrderService(t)
	userService := mocks.NewUserService(t)
	handler := NewApiHandler(bookService, userService, orderService)

	t.Run("success", func(t *testing.T) {
		// set user ID
		userID := 5

		// mock Orders to return
		orders := []*types.Order{
			{
				ID:     rand.Int64(),
				UserID: int64(userID),
			},
			{
				ID:     rand.Int64(),
				UserID: int64(userID),
			},
		}

		// mock ListOrdersByUserId
		orderService.On("ListOrdersByUserId", mock.Anything, mock.Anything).Return(orders, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/orders", nil)
		res := httptest.NewRecorder()

		// set user ID in header
		req.Header.Set("user_id", strconv.Itoa(userID))

		// call ListOrders
		handler.ListOrders(res, req)
		defer res.Result().Body.Close()

		// get body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading body: %v", err)
		}

		// Create a variable of the wrapper with literal struct
		var response struct {
			Data []*types.Order `json:"data"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		orderRes := response.Data

		// assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, orderRes[0].ID, orders[0].ID)
		assert.Equal(t, orderRes[0].UserID, orders[0].UserID)
		assert.Equal(t, orderRes[1].ID, orders[1].ID)
		assert.Equal(t, orderRes[1].UserID, orders[1].UserID)
	})

	t.Run("error - converting userID in header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/orders", nil)
		res := httptest.NewRecorder()

		// set invalid user ID in header
		req.Header.Set("user_id", "abc")

		// call ListOrders
		handler.ListOrders(res, req)
		defer res.Result().Body.Close()

		// assert bad request
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("error - ListOrdersByUserId error", func(t *testing.T) {
		// mock ListOrdersByUserId
		orderService.On("ListOrdersByUserId", mock.Anything, mock.Anything).Return(nil, errors.New("anything")).Once()

		req := httptest.NewRequest(http.MethodPost, "/orders", nil)
		res := httptest.NewRecorder()

		// set user ID in header
		req.Header.Set("user_id", strconv.Itoa(5))

		// call ListOrders
		handler.ListOrders(res, req)
		defer res.Result().Body.Close()

		// assert
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
