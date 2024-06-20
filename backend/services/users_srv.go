package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sebsvt/prototype/logs"
	"github.com/sebsvt/prototype/repositories"
)

var (
	ErrUserEmailAlreadyInUse = errors.New("email already in use")
	ErrUserDoesNotExists     = errors.New("this user account does not exists")
)

type userService struct {
	userRepo repositories.UserRepository
	authSrv  AuthService
}

func NewUserService(userRepo repositories.UserRepository, authSrv AuthService) UserService {
	return userService{userRepo: userRepo, authSrv: authSrv}
}

func (srv userService) CreateUserAccount(newUser CreateNewUserRequest) (*UserResposne, error) {

	// checking email is already in used
	existsUser, err := srv.userRepo.FromEmail(newUser.Email)

	if existsUser != nil {
		logs.Error(ErrUserEmailAlreadyInUse)
		return nil, ErrUserEmailAlreadyInUse
	}

	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return nil, err
	}

	// hashing password
	hashed, err := srv.authSrv.HashPassword(newUser.Password)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newUser.Password = hashed

	// create new user
	user, err := srv.userRepo.CreateNewUser(repositories.User{
		Email:          newUser.Email,
		FirstName:      newUser.FirstName,
		LastName:       newUser.LastName,
		HashedPassword: newUser.Password,
		CreatedAt:      time.Now().Format("2006-1-2 15:04:05"),
	})

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// returing data
	return &UserResposne{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
func (srv userService) GetUser(id int) (*UserResposne, error) {
	user, err := srv.userRepo.FromID(id)
	if err != nil {
		logs.Error(err)
		if err == sql.ErrNoRows {
			return nil, ErrUserDoesNotExists
		}
		return nil, err
	}

	return &UserResposne{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (srv userService) SignIn(email, password string) (string, error) {
	user, err := srv.userRepo.FromEmail(email)
	if err != nil {
		return "", ErrInvalidCredential
	}

	if !srv.authSrv.VerifyPassword(password, user.HashedPassword) {
		return "", ErrInvalidCredential
	}

	accessToken, err := srv.authSrv.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", ErrInvalidCredential
	}

	return accessToken, nil
}
