package user

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(ctx context.Context, userID int) (*User, error) {
	query := `
		SELECT
			"UserId",
			"Username",
			"Firstname",
			"Lastname",
			"Email",
			"Phone",
			"Password",
			"Active",
			"CreatedAt",
			"UpdatedAt",
			"DeletedAt"
		FROM public."User"
		WHERE "UserId" = $1
		  AND "DeletedAt" IS NULL
		  AND "Active" = True
	`

	var user User

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&user.UserID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("User not found with ID: %d", userID)
			return nil, errors.New("user not found")
		}
		// Log the actual error for debugging
		log.Printf("Database error in GetByID: %v", err)
		return nil, err
	}

	return &user, nil
}
