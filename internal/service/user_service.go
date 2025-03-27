package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yanshuy/http-web-server/internal/domain"
	"github.com/yanshuy/http-web-server/internal/repository"
)

type UserService interface {
	GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
}

type userService struct {
	repo *repository.Queries
}

func NewUserService(repo *repository.Queries) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {

	// Map from repository model to domain model
	time := time.Now()
	return &domain.User{
		ID:        id,
		Name:      "John Doe",
		Email:     "john@doe",
		Password:  "password",
		CreatedAt: time,
		UpdatedAt: time,
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	// Business validation
	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	// Create user in database
	now := time.Now()
	created, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		Name: user.Name,

		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return nil, err
	}

	// Map back to domain model
	return &domain.User{
		ID:   created.ID,
		Name: created.Name,

		CreatedAt: created.CreatedAt,
	}, nil
}
