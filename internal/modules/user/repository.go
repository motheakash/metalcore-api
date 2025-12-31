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

func (r *UserRepository) GetAll(ctx context.Context, offset, limit int) ([]User, int64, error) {
	// Get total count
	var totalCount int64
	countQuery := `
		SELECT COUNT(*)
		FROM public."User"
		WHERE "DeletedAt" IS NULL
		  AND "Active" = True
	`

	err := r.db.QueryRow(ctx, countQuery).Scan(&totalCount)
	if err != nil {
		log.Printf("Database error in GetAllRepo (count): %v", err)
		return nil, 0, err
	}

	// Get paginated data
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
		WHERE "DeletedAt" IS NULL
		  AND "Active" = True
		ORDER BY "CreatedAt" DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Printf("Database error in GetAll: %v", err)
		return nil, 0, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
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
			log.Printf("Error scanning user row: %v", err)
			return nil, 0, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, 0, err
	}

	return users, totalCount, nil
}
