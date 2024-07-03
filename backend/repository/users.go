package repository

import "time"

type User struct {
	UserID         int       `db:"user_id"`
	FirstName      string    `db:"firstname"`
	LastName       string    `db:"lastname"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
	CreatedAt      time.Time `db:"created_at"`
}

type UserRepository interface {
	Create(User) (int, error)
	FromID(user_id int) (*User, error)
	FromEmail(email string) (*User, error)
	Save(User) error
}
