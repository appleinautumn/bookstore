package types

import (
	"time"
)

type Book struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Author      NullString `json:"author"`
	Description NullString `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	OrderID  int64 `json:"order_id"`
	BookID   int64 `json:"book_id"`
	Quantity int32 `json:"quantity"`
}

type OrderView struct {
	ID         int64      `json:"id"`
	UserID     int64      `json:"user_id"`
	BookID     int64      `json:"book_id"`
	BookTitle  string     `json:"book_title"`
	BookAuthor NullString `json:"book_author"`
	Quantity   int32      `json:"quantity"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
