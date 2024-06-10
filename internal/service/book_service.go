package service

import (
	"context"

	"gotu/bookstore/internal/repository"
	"gotu/bookstore/internal/types"
)

type BookService struct {
	repository *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{
		repository: repo,
	}
}

func (s *BookService) List(ctx context.Context) ([]*types.Book, error) {
	books, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
