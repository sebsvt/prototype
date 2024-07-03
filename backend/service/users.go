package service

import "time"

type UserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserResponse struct {
	UserID    int       `json:"user_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateNewUser(entity UserRequest) (int, error)
	GetUserFromEmail(email string) (*UserResponse, error)
	GetUserFromID(user_id int) (*UserResponse, error)
	SignIn(UserLogin) (string, error)
	GetCurrentUser(access_token string) (*UserResponse, error)
}
