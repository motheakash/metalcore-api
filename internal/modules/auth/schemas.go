package auth

// LoginRequest represents the HTTP request structure for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the HTTP request structure for user registration
type RegisterRequest struct {
	Username  string  `json:"username" binding:"required,min=3,max=50"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=8"`
	FirstName *string `json:"first_name" binding:"omitempty,max=100"`
	LastName  *string `json:"last_name" binding:"omitempty,max=100"`
	Phone     *string `json:"phone" binding:"omitempty"`
}

// AuthResponse represents the HTTP response structure for authentication
type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in"` // in seconds
	TokenType    string `json:"token_type"` // e.g., "Bearer"
}

// RefreshTokenRequest represents the HTTP request structure for token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
