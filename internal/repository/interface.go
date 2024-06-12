package repository

import (
	"context"

	"gotu/bookstore/internal/types"
)

type BookRepository interface {
	List(ctx context.Context) (res []*types.Book, err error)
}

type OrderRepository interface {
	ListOrdersByUserId(ctx context.Context, userID int64) (res []*types.Order, err error)
	CreateOrder(ctx context.Context, o *types.Order) (res *types.Order, err error)
	CreateOrderItem(ctx context.Context, oi *types.OrderItem) (err error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *types.User) (res *types.User, err error)
}
