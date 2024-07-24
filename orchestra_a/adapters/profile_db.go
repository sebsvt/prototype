package adapters

import (
	"github.com/jmoiron/sqlx"
	"github.com/sebsvt/prototype/orchestra/domain"
)

type profileRepositoryPSQLDB struct {
	db *sqlx.DB
}

func NewProfileRepositoryPSQLDB(db *sqlx.DB) domain.ProfileRepository {
	return profileRepositoryPSQLDB{db: db}
}

// Create implements domain.ProfileRepository.
func (repo profileRepositoryPSQLDB) Create(entity domain.Profile) (int, error) {
	var profile_id int
	query := "insert into profiles (avatar, firstname, surname, gender, phone, date_of_birth, user_ref) values ($1, $2, $3, $4, $5, $6, $7) returning profile_id"
	if err := repo.db.QueryRow(
		query,
		entity.Avatar,
		entity.FirstName,
		entity.Surname,
		entity.Gender,
		entity.Phone,
		entity.DateOfBirth,
		entity.UserRef,
	).Scan(&profile_id); err != nil {
		return 0, err
	}
	return profile_id, nil
}

// FromID implements domain.ProfileRepository.
func (repo profileRepositoryPSQLDB) FromUserRef(user_ref int) (*domain.Profile, error) {
	var profile domain.Profile
	query := "select profile_id, avatar, firstname, surname, gender, phone, date_of_birth, user_ref from profiles where user_ref=$1"
	if err := repo.db.Get(&profile, query, user_ref); err != nil {
		return nil, err
	}
	return &profile, nil
}

// Update implements domain.ProfileRepository.
func (repo profileRepositoryPSQLDB) Update(entity domain.Profile) error {
	query := `
    UPDATE profiles
    SET
        avatar = :avatar,
        firstname = :firstname,
        surname = :surname,
        gender = :gender,
        phone = :phone,
        date_of_birth = :date_of_birth
    WHERE user_ref = :user_ref
    `
	// Use NamedExec for easier struct binding
	_, err := repo.db.NamedExec(query, entity)
	return err
}
