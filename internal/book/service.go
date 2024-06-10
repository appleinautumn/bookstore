package book

import (
	"context"

	"gotu/bookstore/internal/types"
)

type BookService struct {
	repository *BookRepository
}

func NewService(repo *BookRepository) *BookService {
	return &BookService{
		repository: repo,
	}
}

func (s *BookService) List(ctx context.Context) ([]*types.Book, error) {
	// get books
	books, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
