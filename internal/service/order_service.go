package service

import (
	"context"
	"errors"

	"gotu/bookstore/internal/repository"
	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/types"
)

type orderService struct {
	repository repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *orderService {
	return &orderService{
		repository: repo,
	}
}

func (s *orderService) ListOrdersByUserId(ctx context.Context, userID int64) ([]*types.Order, error) {
	orders, err := s.repository.ListOrdersByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderService) CreateOrder(ctx context.Context, req *request.OrderRequest) (*types.Order, error) {
	order := &types.Order{
		UserID: req.UserID,
	}

	if len(req.Orders) == 0 {
		return nil, errors.New("item is empty")
	}

	// create order
	res, err := s.repository.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// create order items
	for _, o := range req.Orders {
		if err := s.repository.CreateOrderItem(ctx, &types.OrderItem{
			OrderID:  res.ID,
			BookID:   int64(o.BookID),
			Quantity: int32(o.Quantity),
		}); err != nil {
			return nil, err
		}
	}

	return res, nil
}
