package repository

import "github.com/jmoiron/sqlx"

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepositoryDB(db *sqlx.DB) ProductRepository {
	return productRepository{db: db}
}

// Create implements ProductRepository.
func (repo productRepository) Create(entity Product) (int, error) {
	query := "insert into products (sku, name, description, price, is_available) values (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(
		query,
		entity.Sku,
		entity.Name,
		entity.Description,
		entity.Price,
		entity.IsAvailable,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// FromID implements ProductRepository.
func (repo productRepository) FromID(product_id int) (*Product, error) {
	var product Product
	query := "select product_id, sku, name, description, price, is_available from products where product_id=?"
	if err := repo.db.Get(&product, query, product_id); err != nil {
		return nil, err
	}
	return &product, nil
}

// List implements ProductRepository.
func (repo productRepository) List() ([]Product, error) {
	var products []Product
	query := "select product_id, sku, name, description, price, is_available from products"
	if err := repo.db.Select(&products, query); err != nil {
		return nil, err
	}
	return products, nil
}
