package service

import (
	"errors"
	"log"
	"time"

	"github.com/sebsvt/prototype/repository"
)

type userService struct {
	user_repo repository.UserRepository
	auth_serv AuthService
}

func NewUserService(user_repo repository.UserRepository, auth_srv AuthService) UserService {
	return userService{user_repo: user_repo, auth_serv: auth_srv}
}

// CreateNewUser implements UserService.
func (srv userService) CreateNewUser(entity UserRequest) (int, error) {
	hashed_password, err := srv.auth_serv.HashPassword(entity.Password)
	if err != nil {
		return 0, err
	}
	user_id, err := srv.user_repo.Create(repository.User{
		FirstName:      entity.Firstname,
		LastName:       entity.Lastname,
		Email:          entity.Email,
		HashedPassword: hashed_password,
		CreatedAt:      time.Now(),
	})
	if err != nil {
		return 0, err
	}
	return user_id, nil
}

// GetUserFromEmail implements UserService.
func (srv userService) GetUserFromEmail(email string) (*UserResponse, error) {
	user, err := srv.user_repo.FromEmail(email)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		UserID:    user.UserID,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

// GetUserFromID implements UserService.
func (srv userService) GetUserFromID(user_id int) (*UserResponse, error) {
	user, err := srv.user_repo.FromID(user_id)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		UserID:    user.UserID,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (srv userService) SignIn(crendetial UserLogin) (string, error) {
	user, err := srv.user_repo.FromEmail(crendetial.Email)
	if err != nil {
		return "", err
	}
	if !(srv.auth_serv.VerifyPassword(crendetial.Password, user.HashedPassword)) {
		return "", errors.New("invalid credential")
	}
	token, err := srv.auth_serv.GenerateToken(user.UserID, crendetial.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (srv userService) GetCurrentUser(access_token string) (*UserResponse, error) {
	user_id, err := srv.auth_serv.ValidateToken(access_token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user, err := srv.user_repo.FromID(user_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &UserResponse{
		UserID:    user_id,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
