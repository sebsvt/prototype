package services

import (
	"database/sql"
	"errors"

	"github.com/sebsvt/prototype/orchestra/domain"
	"github.com/sebsvt/prototype/orchestra/logs"
)

var (
	ErrProfileNotFound = errors.New("user profile not found")
)

type profileService struct {
	profile_repo domain.ProfileRepository
}

func NewProfileService(profile_repo domain.ProfileRepository) ProfileService {
	return profileService{profile_repo: profile_repo}
}

// CreateNewProfile implements ProfileService.
func (repo profileService) CreateNewProfile(entity ProfileRequest, user_ref int) error {
	new_profile := domain.Profile{
		Avatar:      "",
		FirstName:   entity.FirstName,
		Surname:     entity.Surname,
		Gender:      entity.Gender,
		Phone:       entity.Phone,
		DateOfBirth: entity.DateOfBirth,
		UserRef:     user_ref,
	}
	if _, err := repo.profile_repo.Create(new_profile); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

// GetProfileFromUserID implements ProfileService.
func (repo profileService) GetProfileFromUserID(user_id int) (*ProfileResponse, error) {
	profile, err := repo.profile_repo.FromUserRef(user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProfileNotFound
		}
		return nil, err
	}
	return &ProfileResponse{
		UserRef:     profile.UserRef,
		FirstName:   profile.FirstName,
		Surname:     profile.Surname,
		Gender:      profile.Gender,
		Phone:       profile.Phone,
		DateOfBirth: profile.DateOfBirth,
	}, nil
}
