package repository

import (
	"context"
	"database/sql"
	"log/slog"
	"strings"

	"gotu/bookstore/internal/types"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) ListOrdersByUserId(ctx context.Context, userID int64) (res []*types.Order, err error) {
	var sb strings.Builder

	sb.WriteString("SELECT id, user_id, created_at, updated_at FROM orders WHERE user_id = $1")
	sb.WriteString(";")

	rows, err := r.db.Query(sb.String(), userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order types.Order
		if err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.CreatedAt,
			&order.UpdatedAt); err != nil {
			return nil, err
		}

		res = append(res, &order)
	}

	return res, nil
}

func (r *orderRepository) CreateOrder(ctx context.Context, o *types.Order) (res *types.Order, err error) {
	sql := `INSERT INTO orders (user_id)
			VALUES ($1)
			RETURNING id, user_id, created_at, updated_at`

	row := r.db.QueryRow(sql, o.UserID)

	if err := row.Scan(
		&o.ID,
		&o.UserID,
		&o.CreatedAt,
		&o.UpdatedAt); err != nil {
		return nil, err
	}

	return o, nil
}

func (r *orderRepository) CreateOrderItem(ctx context.Context, o *types.OrderItem) (err error) {
	sql := `INSERT INTO order_books (order_id, book_id, quantity)
			VALUES ($1, $2, $3)`

	row := r.db.QueryRow(sql, o.OrderID, o.BookID, o.Quantity)
	slog.Info("Orderitem created", slog.Any("row", row))

	return nil
}
