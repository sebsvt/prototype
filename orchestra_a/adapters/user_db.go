package adapters

import (
	"github.com/jmoiron/sqlx"
	"github.com/sebsvt/prototype/orchestra/domain"
)

type userRepositoryPSQLDB struct {
	db *sqlx.DB
}

func NewUserRepositoryPSQLDB(db *sqlx.DB) domain.UserRepository {
	return userRepositoryPSQLDB{db: db}
}

// Create implements domain.UserRepository.
func (repo userRepositoryPSQLDB) Create(entity domain.User) (int, error) {
	var userID int
	query := "insert into users (email, hashed_password, is_active, is_verified, last_signed_in, created_at, updated_at, deleted_at) values ($1, $2, $3, $4, $5, $6, $7, $8) returning user_id"
	row := repo.db.QueryRow(
		query,
		entity.Email,
		entity.HashedPassword,
		entity.IsActive,
		entity.IsVerified,
		entity.LastSignedIn,
		entity.CreatedAt,
		entity.UpdatedAt,
		entity.DeletedAt,
	)
	if err := row.Scan(&userID); err != nil {
		return 0, err
	}
	return userID, nil
}

// FromEmail implements domain.UserRepository.
func (repo userRepositoryPSQLDB) FromEmail(email string) (*domain.User, error) {
	var user domain.User
	query := "select user_id, email, hashed_password, is_active, is_verified, last_signed_in, created_at, updated_at, deleted_at from users where email=$1"
	if err := repo.db.Get(&user, query, email); err != nil {
		return nil, err
	}
	return &user, nil
}

// FromID implements domain.UserRepository.
func (repo userRepositoryPSQLDB) FromID(id int) (*domain.User, error) {
	var user domain.User
	query := "select user_id, email, hashed_password, is_active, is_verified, last_signed_in, created_at, updated_at, deleted_at from users where user_id=$1"
	if err := repo.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

// Update implements domain.UserRepository.
func (repo userRepositoryPSQLDB) Update(entity domain.User) error {
	query := `
    UPDATE users
    SET
        email = :email,
        hashed_password = :hashed_password,
        is_active = :is_active,
        is_verified = :is_verified,
        last_signed_in = :last_signed_in,
        updated_at = :updated_at,
        deleted_at = :deleted_at
    WHERE user_id = :user_id
    `

	_, err := repo.db.NamedExec(query, entity)
	return err
}
