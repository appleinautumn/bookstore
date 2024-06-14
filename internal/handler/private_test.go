package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/service/mocks"
	"gotu/bookstore/internal/types"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

		req := httptest.NewRequest(http.MethodPost, "/my/orders", bytes.NewReader(payloadJson))
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

		req := httptest.NewRequest(http.MethodPost, "/my/orders", bytes.NewReader(payloadJson))
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

		req := httptest.NewRequest(http.MethodPost, "/my/orders", bytes.NewReader(payloadJson))
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

			req := httptest.NewRequest(http.MethodPost, "/my/orders", bytes.NewReader(payloadJson))
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

		req := httptest.NewRequest(http.MethodPost, "/my/orders", bytes.NewReader(payloadJson))
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
		orderViews := []*types.OrderView{
			{
				ID:        rand.Int64(),
				UserID:    int64(userID),
				BookID:    rand.Int64(),
				BookTitle: faker.Word(),
				BookAuthor: types.NullString{
					NullString: sql.NullString{
						String: faker.Word(),
						Valid:  true,
					},
				},
				Quantity: rand.Int32(),
			},
			{
				ID:        rand.Int64(),
				UserID:    int64(userID),
				BookID:    rand.Int64(),
				BookTitle: faker.Word(),
				BookAuthor: types.NullString{
					NullString: sql.NullString{
						String: faker.Word(),
						Valid:  true,
					},
				},
				Quantity: rand.Int32(),
			},
		}

		// mock ListOrdersByUserId
		orderService.On("ListOrdersByUserId", mock.Anything, mock.Anything).Return(orderViews, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/my/orders", nil)
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
			Data []*types.OrderView `json:"data"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf("Error unmarshaling JSON: %v", err)
		}

		orderRes := response.Data

		// assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, orderRes[0].ID, orderViews[0].ID)
		assert.Equal(t, orderRes[0].UserID, orderViews[0].UserID)
		assert.Equal(t, orderRes[1].ID, orderViews[1].ID)
		assert.Equal(t, orderRes[1].UserID, orderViews[1].UserID)
	})

	t.Run("error - converting userID in header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/my/orders", nil)
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

		req := httptest.NewRequest(http.MethodPost, "/my/orders", nil)
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
