package repository

import "github.com/jmoiron/sqlx"

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepository{db: db}
}

// Create implements UserRepository.
func (repo userRepository) Create(entity User) (int, error) {
	query := "insert into users(firstname, lastname, email, hashed_password, created_at) values (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(
		query,
		entity.FirstName,
		entity.LastName,
		entity.Email,
		entity.HashedPassword,
		entity.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	user_id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(user_id), nil
}

// FromEmail implements UserRepository.
func (repo userRepository) FromEmail(email string) (*User, error) {
	var user User
	query := "select user_id, firstname, lastname, email, hashed_password, created_at from users where email=?"
	if err := repo.db.Get(&user, query, email); err != nil {
		return nil, err
	}
	return &user, nil
}

// FromID implements UserRepository.
func (repo userRepository) FromID(user_id int) (*User, error) {
	var user User
	query := "select user_id, firstname, lastname, email, hashed_password, created_at from users where user_id=?"
	if err := repo.db.Get(&user, query, user_id); err != nil {
		return nil, err
	}
	return &user, nil
}

// Save implements UserRepository.
func (repo userRepository) Save(entity User) error {
	query := `
		UPDATE users
		SET firstname = ?, lastname = ?, email = ?, hashed_password = ?, created_at = ?
		WHERE user_id = ?`

	_, err := repo.db.Exec(
		query,
		entity.FirstName,
		entity.LastName,
		entity.Email,
		entity.HashedPassword,
		entity.CreatedAt,
		entity.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}
