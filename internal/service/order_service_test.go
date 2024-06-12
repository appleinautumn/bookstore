package service

import (
	"context"
	"errors"
	"math/rand/v2"
	"reflect"
	"testing"

	"gotu/bookstore/internal/repository/mocks"
	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/types"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListOrdersByUserId(t *testing.T) {
	repository := mocks.NewOrderRepository(t)

	t.Run("success", func(t *testing.T) {
		// mock order1
		var order1 types.Order
		if err := faker.FakeData(&order1); err != nil {
			t.Errorf("err: %v", err)
		}

		var order2 types.Order
		if err := faker.FakeData(&order2); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock orders
		orders := []*types.Order{
			&order1,
			&order2,
		}

		repository.On("ListOrdersByUserId", mock.Anything, mock.Anything).Return(orders, nil).Once()
		service := NewOrderService(repository)

		userID := int64(rand.IntN(100))

		// create
		res, err := service.ListOrdersByUserId(context.TODO(), userID)

		// assert
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("error - ListOrdersByUserId error", func(t *testing.T) {

		// mock order1
		var order1 types.Order
		if err := faker.FakeData(&order1); err != nil {
			t.Errorf("err: %v", err)
		}

		repository.On("ListOrdersByUserId", mock.Anything, mock.Anything).Return(nil, errors.New("error"))
		service := NewOrderService(repository)

		userID := int64(rand.IntN(100))

		// create
		res, err := service.ListOrdersByUserId(context.TODO(), userID)

		// assert
		assert.Error(t, err)
		assert.Nil(t, res)
	})

}

func TestCreateOrder(t *testing.T) {
	repository := mocks.NewOrderRepository(t)

	t.Run("success", func(t *testing.T) {
		// mock order item 1
		var item1 request.OrderItem
		if err := faker.FakeData(&item1); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock order item 2
		var item2 request.OrderItem
		if err := faker.FakeData(&item2); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock order request
		orderReq := request.OrderRequest{
			Orders: []*request.OrderItem{
				&item1,
				&item2,
			},
			UserID: int64(rand.IntN(100)),
		}

		// mock order1
		var order1 types.Order
		if err := faker.FakeData(&order1); err != nil {
			t.Errorf("err: %v", err)
		}

		repository.On("CreateOrder", mock.Anything, mock.Anything).Return(&order1, nil).Once()
		repository.On("CreateOrderItem", mock.Anything, mock.Anything).Return(nil).Twice()
		service := NewOrderService(repository)

		// create
		res, err := service.CreateOrder(context.TODO(), &orderReq)

		if !reflect.DeepEqual(&order1, res) {
			t.Errorf("res and mockEntity are not deep equal")
		}

		// assert
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("error - empty orders", func(t *testing.T) {
		// mock order request - empty order
		orderReq := request.OrderRequest{
			Orders: []*request.OrderItem{},
			UserID: int64(rand.IntN(100)),
		}

		// mock order1
		var order1 types.Order
		if err := faker.FakeData(&order1); err != nil {
			t.Errorf("err: %v", err)
		}

		service := NewOrderService(repository)

		// create
		res, err := service.CreateOrder(context.TODO(), &orderReq)

		// assert
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - create order error", func(t *testing.T) {
		// mock order item 1
		var item1 request.OrderItem
		if err := faker.FakeData(&item1); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock order item 2
		var item2 request.OrderItem
		if err := faker.FakeData(&item2); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock order request
		orderReq := request.OrderRequest{
			Orders: []*request.OrderItem{
				&item1,
				&item2,
			},
			UserID: int64(rand.IntN(100)),
		}

		// mock order1
		var order1 types.Order
		if err := faker.FakeData(&order1); err != nil {
			t.Errorf("err: %v", err)
		}

		repository.On("CreateOrder", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		service := NewOrderService(repository)

		// create
		res, err := service.CreateOrder(context.TODO(), &orderReq)

		// assert
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - create order item error", func(t *testing.T) {
		// mock order item 1
		var item1 request.OrderItem
		if err := faker.FakeData(&item1); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock order item 2
		var item2 request.OrderItem
		if err := faker.FakeData(&item2); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock order request
		orderReq := request.OrderRequest{
			Orders: []*request.OrderItem{
				&item1,
				&item2,
			},
			UserID: int64(rand.IntN(100)),
		}

		// mock order1
		var order1 types.Order
		if err := faker.FakeData(&order1); err != nil {
			t.Errorf("err: %v", err)
		}

		repository.On("CreateOrder", mock.Anything, mock.Anything).Return(&order1, nil).Once()
		repository.On("CreateOrderItem", mock.Anything, mock.Anything).Return(errors.New("order item insert error"))
		service := NewOrderService(repository)

		// create and throw error
		res, err := service.CreateOrder(context.TODO(), &orderReq)

		// assert
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
