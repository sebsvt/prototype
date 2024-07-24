package adapters

import (
	"github.com/jmoiron/sqlx"
	"github.com/sebsvt/prototype/orchestra/domain"
)

type membershipRepositoryPSQLDB struct {
	db *sqlx.DB
}

func NewMembershipRepositoryPSQLDB(db *sqlx.DB) domain.MembershipRepository {
	return membershipRepositoryPSQLDB{db: db}
}

// Add implements domain.MemebershipRepository.
func (repo membershipRepositoryPSQLDB) Add(entity domain.Membership) error {
	query := "insert into memberships (member_id, studio_id, role) values ($1, $2, $3)"
	if _, err := repo.db.Exec(query, entity.MemberID, entity.StudioID, entity.Role); err != nil {
		return err
	}
	return nil
}

// Find implements domain.MemebershipRepository.
func (repo membershipRepositoryPSQLDB) Find(member_id int, studio_id int) (*domain.Membership, error) {
	var membership domain.Membership
	query := "select member_id, studio_id, role from memberships where member_id=$1 and studio_id=$2"
	if err := repo.db.Get(membership, query, member_id, studio_id); err != nil {
		return nil, err
	}
	return &membership, nil
}

// Remove implements domain.MemebershipRepository.
func (repo membershipRepositoryPSQLDB) Remove(member_id int, studio_id int) error {
	query := "delete from memberships where studio_id = $1 and user_id = $2"
	if _, err := repo.db.Exec(query, member_id, studio_id); err != nil {
		return err
	}
	return nil
}

func (repo membershipRepositoryPSQLDB) All(member_id int) ([]domain.Membership, error) {
	var list_of_members []domain.Membership
	query := "select member_id, studio_id, role from memberships where member_id=$1"
	if err := repo.db.Select(&list_of_members, query, member_id); err != nil {
		return nil, err
	}
	return list_of_members, nil
}
