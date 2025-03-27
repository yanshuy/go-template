package service

import (
	"github.com/yanshuy/http-web-server/internal/repository"
)

type Service struct {
	Users UserService
}

func New(repo *repository.Queries) *Service {
	return &Service{
		Users: NewUserService(repo),
	}
}
