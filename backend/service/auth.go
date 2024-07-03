package service

import "github.com/golang-jwt/jwt/v5"

type TokenData struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type AuthService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hashed_password string) bool
	GenerateToken(user_id int, email string) (string, error)
	ValidateToken(token string) (int, error)
}
