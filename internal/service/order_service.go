package service

import (
	"context"
	"errors"
	"log/slog"

	"gotu/bookstore/internal/repository"
	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/types"
)

type OrderService struct {
	repository *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{
		repository: repo,
	}
}

func (s *OrderService) ListOrdersByUserId(ctx context.Context, userID int64) ([]*types.Order, error) {
	orders, err := s.repository.ListOrdersByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, req *request.OrderRequest) (*types.Order, error) {
	order := &types.Order{
		UserID: req.UserID,
	}
	slog.Info("Order created", slog.Int("req.Orders", len(req.Orders)))

	if len(req.Orders) == 0 {
		return nil, errors.New("item is empty")
	}

	// insert order
	res, err := s.repository.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// insert order items
	for _, o := range req.Orders {
		if err := s.repository.CreateOrderItem(ctx, &types.OrderItem{
			OrderID:  res.ID,
			BookID:   int64(o.BookID),
			Quantity: int32(o.Quantity),
		}); err != nil {
			return nil, err
		}
	}
	slog.Info("Order items created", slog.Any("res", res))

	return res, nil
}
