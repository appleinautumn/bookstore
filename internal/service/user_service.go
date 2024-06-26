package service

import (
	"context"

	"gotu/bookstore/internal/repository"
	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/types"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *request.SignUpRequest) (*types.User, error) {
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
