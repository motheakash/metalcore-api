package user

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserInactive = errors.New("user is inactive")
)

type Service struct {
	repo *UserRepository
}

func NewService(repo *UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByID(ctx context.Context, userID int) (*User, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if !user.Active {
		return nil, ErrUserInactive
	}

	// business rules can grow here
	return user, nil
}
