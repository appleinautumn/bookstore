// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "gotu/bookstore/internal/types"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: ctx, o
func (_m *OrderRepository) CreateOrder(ctx context.Context, o *types.Order) (*types.Order, error) {
	ret := _m.Called(ctx, o)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *types.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Order) (*types.Order, error)); ok {
		return rf(ctx, o)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Order) *types.Order); ok {
		r0 = rf(ctx, o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Order) error); ok {
		r1 = rf(ctx, o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrderItem provides a mock function with given fields: ctx, oi
func (_m *OrderRepository) CreateOrderItem(ctx context.Context, oi *types.OrderItem) error {
	ret := _m.Called(ctx, oi)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrderItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.OrderItem) error); ok {
		r0 = rf(ctx, oi)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListOrdersByUserId provides a mock function with given fields: ctx, userID
func (_m *OrderRepository) ListOrdersByUserId(ctx context.Context, userID int64) ([]*types.Order, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for ListOrdersByUserId")
	}

	var r0 []*types.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*types.Order, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*types.Order); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
