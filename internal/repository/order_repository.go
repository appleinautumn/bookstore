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

func (r *orderRepository) ListOrdersByUserId(ctx context.Context, userID int64) (res []*types.OrderView, err error) {
	var sb strings.Builder

	sb.WriteString(`SELECT
						o.id,
						o.user_id,
						ob.book_id,
						b.title AS book_title,
						b.author AS book_author,
						ob.quantity,
						o.created_at,
						o.updated_at
					FROM orders o
					INNER JOIN order_books ob ON ob.order_id = o.id
					INNER JOIN books b ON b.id = ob.book_id
					WHERE user_id=$1`)
	sb.WriteString(";")

	rows, err := r.db.Query(sb.String(), userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ov types.OrderView
		if err := rows.Scan(
			&ov.ID,
			&ov.UserID,
			&ov.BookID,
			&ov.BookTitle,
			&ov.BookAuthor,
			&ov.Quantity,
			&ov.CreatedAt,
			&ov.UpdatedAt); err != nil {
			return nil, err
		}

		res = append(res, &ov)
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
