package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredential = errors.New("invalid credential")
)

var secret = []byte("5369d0098d3aa52ae218048702c10624abcf47154ca051d65c7653fe563e8482")

type authService struct {
}

func NewAuthService() AuthService {
	return authService{}
}

func (srv authService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (srv authService) VerifyPassword(password string, hashed_password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (srv authService) GenerateToken(userID int, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_id"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (srv authService) ValidateToken(token string) (int, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidCredential
		}
		return secret, nil
	})
	if err != nil {
		return 0, ErrInvalidCredential
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, ErrInvalidCredential
		}
		return int(userID), nil
	}

	return 0, ErrInvalidCredential
}
