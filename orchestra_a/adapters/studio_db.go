package adapters

import (
	"github.com/jmoiron/sqlx"
	"github.com/sebsvt/prototype/orchestra/domain"
)

type studioRepositoryPSQLDB struct {
	db *sqlx.DB
}

func NewStudioRepositoryPSQLDB(db *sqlx.DB) domain.StudioRepository {
	return studioRepositoryPSQLDB{db: db}
}

// Create implements domain.StudioRepository.
func (repo studioRepositoryPSQLDB) Create(entity domain.Studio) (int, error) {
	var studio_id int
	query := `
		insert into studios (subdomain, picture, name, description, address, city, zipcode, state, country, owner_id, created_at, updated_at, deleted_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		returning studio_id
	`
	if err := repo.db.QueryRow(
		query,
		entity.SubDomain,
		entity.Picture,
		entity.Name,
		entity.Description,
		entity.Address,
		entity.City,
		entity.ZipCode,
		entity.State,
		entity.Country,
		entity.OwnerID,
		entity.CreatedAt,
		entity.UpdatedAt,
		entity.DeletedAt,
	).Scan(&studio_id); err != nil {
		return 0, nil
	}
	return studio_id, nil
}

// FromSubDomain implements domain.StudioRepository.
func (repo studioRepositoryPSQLDB) FromSubDomain(subdomain string) (*domain.Studio, error) {
	var studio domain.Studio
	query := "select * from studios where subdomain=$1"
	if err := repo.db.Get(&studio, query, subdomain); err != nil {
		return nil, err
	}
	return &studio, nil
}

func (repo studioRepositoryPSQLDB) FromID(studio_id int) (*domain.Studio, error) {
	var studio domain.Studio
	query := "select * from studios where studio_id=$1"
	if err := repo.db.Get(&studio, query, studio_id); err != nil {
		return nil, err
	}
	return &studio, nil
}

// Update implements domain.StudioRepository.
func (repo studioRepositoryPSQLDB) Update(entity domain.Studio) error {
	query := `
        UPDATE studios
        SET subdomain = $1, picture = $2, name = $3, description = $4,
            address = $5, city = $6, zipcode = $7, state = $8, country = $9,
            updated_at = $10
        WHERE studio_id = $11
    `
	_, err := repo.db.Exec(query,
		entity.SubDomain, entity.Picture, entity.Name, entity.Description,
		entity.Address, entity.City, entity.ZipCode, entity.State,
		entity.Country, entity.UpdatedAt, entity.StudioID,
	)
	return err
}
