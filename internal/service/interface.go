package service

import (
	"context"

	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/types"
)

type BookService interface {
	List(ctx context.Context) ([]*types.Book, error)
}

type OrderService interface {
	ListOrdersByUserId(ctx context.Context, userID int64) ([]*types.Order, error)
	CreateOrder(ctx context.Context, req *request.OrderRequest) (*types.Order, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req *request.SignUpRequest) (*types.User, error)
}
