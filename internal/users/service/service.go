package service

import (
	"context"
	"net/http"
	"uuid"

	customError "github.com/yourusername/go-wallets/internal/errors"
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
	existingUser, _ := s.repo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, customError.NewAppError(http.StatusConflict, "EMAIL_ALREADY_REGISTERED", "this email already registered.")
	}
	user := &model.User{
		Id:           uuid.New().String(),
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: req.Password, //Todo: hash password
	}
	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, customError.ErrInternalServer
	}
	return user, nil
}

func (s *userService) GetProfile(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, customError.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", "user not found")
	}
	return user, nil
}

func (s *userService) UpdateProfile(ctx context.Context, id string, req model.UpdateUserRequest) error {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return customError.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", "user not found")
	}
	user = &model.User{
		Id:       id,
		FullName: req.FullName,
	}
	err = s.repo.Update(ctx, user)
	if err != nil {
		return customError.ErrInternalServer
	}
	return nil
}
