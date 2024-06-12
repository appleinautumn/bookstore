package service

import (
	"context"

	"gotu/bookstore/internal/repository"
	"gotu/bookstore/internal/types"
)

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *bookService {
	return &bookService{
		repository: repo,
	}
}

func (s *bookService) List(ctx context.Context) ([]*types.Book, error) {
	books, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
