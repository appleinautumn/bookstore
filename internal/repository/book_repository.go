package repository

import (
	"context"
	"database/sql"
	"strings"

	"gotu/bookstore/internal/types"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) List(ctx context.Context) (res []*types.Book, err error) {
	var sb strings.Builder

	sb.WriteString("SELECT id, title, author, description, created_at, updated_at FROM books")
	sb.WriteString(";")

	rows, err := r.db.Query(sb.String())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book types.Book
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.CreatedAt,
			&book.UpdatedAt); err != nil {
			return nil, err
		}

		res = append(res, &book)
	}

	return res, nil
}
