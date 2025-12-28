package user

import "time"

type User struct {
	UserID    int        `db:"UserId" json:"user_id"`
	Username  string     `db:"Username" json:"username"`
	FirstName *string    `db:"Firstname" json:"first_name,omitempty"`
	LastName  *string    `db:"Lastname" json:"last_name,omitempty"`
	Email     string     `db:"Email" json:"email"`
	Phone     *string    `db:"Phone" json:"phone,omitempty"`
	Password  string     `db:"Password" json:"-"` // never expose
	Active    bool       `db:"Active" json:"active"`
	CreatedAt time.Time  `db:"CreatedAt" json:"created_at"`
	UpdatedAt *time.Time `db:"UpdatedAt" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"DeletedAt" json:"deleted_at,omitempty"`
}
