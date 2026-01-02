package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUserInactive   = errors.New("user is inactive")
	ErrUsernameExists = errors.New("username already exists")
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

func (s *Service) Create(ctx context.Context, payload CreateUserRequest) (*User, error) {

	exists, err := s.repo.UsernameExists(ctx, payload.Username)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrUsernameExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username:  payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		Password:  string(hashedPassword),
		Active:    true,
	}

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
