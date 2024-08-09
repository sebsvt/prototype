package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed_password), nil
}

func VerifyPassword(password, hashed_password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password)); err != nil {
		return false
	}
	return true
}
