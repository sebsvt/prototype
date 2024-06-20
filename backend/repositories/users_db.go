package repositories

import (
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) FromID(id int) (*User, error) {

	var user User
	query := "select id, firstname, lastname, email, hashed_password, created_at from users where id = $1"

	err := repo.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo userRepository) FromEmail(email string) (*User, error) {

	var user User
	query := "select id, firstname, lastname, email, hashed_password, created_at from users where email = $1"

	err := repo.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (repo userRepository) CreateNewUser(newUser User) (*User, error) {

	query := "insert into users (firstname, lastname, email, hashed_password, created_at) values ($1,$2,$3,$4,$5) RETURNING id"

	err := repo.db.QueryRow(
		query,
		newUser.FirstName,
		newUser.LastName,
		newUser.Email,
		newUser.HashedPassword,
		newUser.CreatedAt,
	).Scan(&newUser.ID)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
