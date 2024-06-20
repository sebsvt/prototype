package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sebsvt/prototype/logs"
	// "github.com/sebsvt/prototype/logs"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredential = errors.New("invalid credentail")
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
		logs.Error(err)
		return "", err
	}
	return string(hashedPassword), nil
}
func (srv authService) VerifyPassword(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		logs.Error(err)
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
	claims["exp"] = time.Now().Add(time.Second * 30).Unix()

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (srv authService) ValidateToken(token string) error {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidCredential
		}
		return []byte(secret), nil
	})
	if err != nil {
		logs.Error(err)
		return ErrInvalidCredential
	}
	if !t.Valid {
		logs.Debug("invalid token")
		return ErrInvalidCredential
	}
	return nil
}
