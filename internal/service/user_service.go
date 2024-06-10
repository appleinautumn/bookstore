package service

import (
	"context"

	"gotu/bookstore/internal/repository"
	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/types"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *request.SignUpRequest) (*types.User, error) {
	user := &types.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	res, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}
