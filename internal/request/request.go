package request

type SignUpRequest struct {
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type OrderItem struct {
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`
}

type OrderRequest struct {
	Orders []*OrderItem `json:"orders"`
	UserID int64        `json:"user_id"`
}
