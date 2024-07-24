package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sebsvt/prototype/orchestra/configs"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateAccessToken(email string, user_id int) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(user_id),
			ExpiresAt: time.Now().Add(time.Minute * 21600).Unix(),
		},
	}
	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwt_token.SignedString([]byte(configs.EnvConfig.SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(token_string string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(token_string, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.EnvConfig.SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
