package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sebsvt/prototype/authentication/domain"
	"github.com/sebsvt/prototype/authentication/logs"
	"github.com/sebsvt/prototype/authentication/utils"
)

var (
	ErrUserEmailAlreadyInUse = errors.New("user's email already in use")
	ErrInvalidCredential     = errors.New("invalid credential")
)

type authService struct {
	user_repo domain.UserRepository
}

func NewAuthService(user_repo domain.UserRepository) AuthService {
	return authService{user_repo: user_repo}
}

// SignIn implements AuthService.
func (srv authService) SignIn(credential SignInRequest) (*TokenData, error) {
	user, err := srv.user_repo.FromEmail(credential.Email)
	if err != nil {
		logs.Error(err)
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredential
		}
		return nil, err
	}
	if is_matched := utils.VerifyPassword(credential.Password, user.HashedPassword); !is_matched {
		return nil, ErrInvalidCredential
	}
	user.LastSignedIn = time.Now()
	if err := srv.user_repo.Update(*user); err != nil {
		logs.Error(err)
		return nil, err
	}
	token, err := utils.GenerateAccessToken(user.Email, user.UserID)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	access_token := TokenData{
		AccessToken: token,
		Type:        "Bearer",
	}
	return &access_token, nil
}

// SignUp implements AuthService.
func (srv authService) SignUp(credential SignUpRequest) (*TokenData, error) {
	// checking is email alreay in use
	exists_user, err := srv.user_repo.FromEmail(credential.Email)
	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return nil, err
	}
	if exists_user != nil {
		logs.Error(ErrUserEmailAlreadyInUse)
		return nil, ErrUserEmailAlreadyInUse
	}
	hashed_password, err := utils.HashPassword(credential.Password)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	// create new user
	new_user := domain.User{
		Email:          credential.Email,
		HashedPassword: hashed_password,
		IsActive:       true,
		IsVerified:     false,
		LastSignedIn:   time.Time{},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
	}
	user_id, err := srv.user_repo.Create(new_user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	// create access token
	token, err := utils.GenerateAccessToken(new_user.Email, user_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	access_token := TokenData{
		AccessToken: token,
		Type:        "Bearer",
	}
	return &access_token, nil
}

// VerityToken implements AuthService.
func (srv authService) VerityToken(access_token string) (*ClaimsData, error) {
	claims, err := utils.VerifyToken(access_token)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &ClaimsData{
		UserID: claims.Issuer,
		Email:  claims.Email,
	}, nil
}
