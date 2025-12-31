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

func (s *Service) GetAll(ctx context.Context, page, page_size int) ([]User, int64, error) {
	users, total_count, err := s.repo.GetAll(ctx, page, page_size)
	if err != nil {
		return nil, 0, err
	}

	// Additional business logic can be added here
	// For example: filtering, sorting, enrichment, etc.
	return users, total_count, nil
}
