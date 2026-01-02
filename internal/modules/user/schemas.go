package user

import "time"

// UserResponse represents the HTTP response structure for a user
// Excludes sensitive fields like password
type UserResponse struct {
	UserID    int        `json:"user_id"`
	Username  string     `json:"username"`
	FirstName *string    `json:"first_name,omitempty"`
	LastName  *string    `json:"last_name,omitempty"`
	Email     string     `json:"email"`
	Phone     *string    `json:"phone,omitempty"`
	Active    bool       `json:"active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// CreateUserRequest represents the HTTP request structure for creating a user
type CreateUserRequest struct {
	Username  string  `json:"username" binding:"required,min=3,max=50"`
	FirstName *string `json:"first_name" binding:"omitempty,max=100"`
	LastName  *string `json:"last_name" binding:"omitempty,max=100"`
	Email     string  `json:"email" binding:"required,email"`
	Phone     *string `json:"phone" binding:"required,min=10,max=13"`
	Password  string  `json:"password" binding:"required,min=8"`
}

// UpdateUserRequest represents the HTTP request structure for updating a user
type UpdateUserRequest struct {
	FirstName *string `json:"first_name" binding:"omitempty,max=100"`
	LastName  *string `json:"last_name" binding:"omitempty,max=100"`
	Email     *string `json:"email" binding:"omitempty,email"`
	Phone     *string `json:"phone" binding:"omitempty"`
	Active    *bool   `json:"active" binding:"omitempty"`
}

// ToUserResponse converts a User model to UserResponse schema
func ToUserResponse(user *User) *UserResponse {
	if user == nil {
		return nil
	}

	return &UserResponse{
		UserID:    user.UserID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToUserListResponse converts a slice of User models to UserResponse schemas
func ToUserListResponse(users []User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, user := range users {
		responses[i] = UserResponse{
			UserID:    user.UserID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Active:    user.Active,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return responses
}
