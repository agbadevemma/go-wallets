package service

import (
	"context"
	"errors"
	"uuid"

	"github.com/yourusername/go-wallets/internal/users/model"
	"github.com/yourusername/go-wallets/internal/users/repository"
)

type UserService interface {
	Register(ctx context.Context, req model.CreateUserRequest) (*model.User, error)
	GetProfile(ctx context.Context, id string) (*model.User, error)
	UpdateProfile(ctx context.Context, id string, req model.UpdateUserRequest) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, req model.CreateUserRequest) (*model.User, error) {
	// Check if the email already exists
	existingUser, err := s.repo.FindByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}
	user := &model.User{
		Id:       uuid.New().String(),
		FullName: req.FullName,
		Email:    req.Email,
		PasswordHash: req.Password,//Todo: hash password
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetProfile(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindById(ctx, id)
}

func (s *userService) UpdateProfile(ctx context.Context, id string, req model.UpdateUserRequest) error {
	user := &model.User{
		Id:       id,
		FullName: req.FullName,
	}
	err := s.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
