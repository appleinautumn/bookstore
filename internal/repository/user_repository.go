package repository

import (
	"context"
	"database/sql"

	"gotu/bookstore/internal/types"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, u *types.User) (res *types.User, err error) {
	sql := `INSERT INTO users (email, name, password)
			VALUES ($1, $2, $3)
			RETURNING id, email, name, created_at, updated_at`

	row := r.db.QueryRow(sql, u.Email, u.Name, u.Password)

	if err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Name,
		&u.CreatedAt,
		&u.UpdatedAt); err != nil {
		return nil, err
	}

	return u, nil
}
