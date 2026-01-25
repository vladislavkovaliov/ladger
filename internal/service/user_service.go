package service

import (
	"context"

	"github.com/vladislavkovaliov/ledger/internal/domain/user"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, u *user.User) error {
	return s.repo.Save(ctx, u)
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *UserService) List(ctx context.Context) ([]*user.UserResponse, error) {
	return s.repo.List(ctx)
}
