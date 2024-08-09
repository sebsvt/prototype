package domain

import "time"

type User struct {
	UserID         int        `db:"user_id"`
	Email          string     `db:"email"`
	HashedPassword string     `db:"hashed_password"`
	IsActive       bool       `db:"is_active"`
	IsVerified     bool       `db:"is_verified"`
	LastSignedIn   time.Time  `db:"last_signed_in"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"` // Use pointer to allow null values
}

type UserRepository interface {
	Create(entity User) (int, error)
	FromID(id int) (*User, error)
	FromEmail(email string) (*User, error)
	Update(entity User) error
}
