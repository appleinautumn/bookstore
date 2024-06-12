// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "gotu/bookstore/internal/types"
)

// BookService is an autogenerated mock type for the BookService type
type BookService struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx
func (_m *BookService) List(ctx context.Context) ([]*types.Book, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*types.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*types.Book, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*types.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBookService creates a new instance of BookService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookService(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookService {
	mock := &BookService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}