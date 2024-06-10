package types

import (
	"database/sql"
	"time"
)

type Book struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Author      sql.NullString `json:"author"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
