package repository

import (
	"github.com/jmoiron/sqlx"
)

// Order represents the order struct.
type orderRepositoryDB struct {
	db *sqlx.DB
}

func NewOrderRepositoryDB(db *sqlx.DB) OrderRepository {
	return orderRepositoryDB{db: db}
}

// CreateNewOrder implements OrderRepository.
func (repo orderRepositoryDB) CreateNewOrder(entity Order) (*Order, error) {
	query := "INSERT INTO orders (customer_id, product_sku, product_cost, duration, payment_id, is_paid, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(
		query,
		entity.CustomerID,
		entity.ProductSKU,
		entity.ProductCost,
		entity.Duration,
		entity.PaymentID,
		entity.IsPaid,
		entity.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	orderID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	entity.OrderID = int(orderID)
	return &entity, nil
}

// FromOrderID implements OrderRepository.
func (repo orderRepositoryDB) FromOrderID(orderID int) (*Order, error) {
	var order Order
	query := "SELECT order_id, customer_id, product_sku, product_cost, duration, payment_id, is_paid, created_at FROM orders WHERE order_id=?"
	if err := repo.db.Get(&order, query, orderID); err != nil {
		return nil, err
	}
	return &order, nil
}

// GetAllOrder implements OrderRepository.
func (repo orderRepositoryDB) GetAllOrder(user_id int) ([]Order, error) {
	var orders []Order
	query := "SELECT order_id, customer_id, product_sku, product_cost, duration, payment_id, is_paid, created_at FROM orders where customer_id = ?"
	if err := repo.db.Select(&orders, query, user_id); err != nil {
		return nil, err
	}
	return orders, nil
}

func (repo orderRepositoryDB) UpdateOrder(entity Order) error {
	query := `
		UPDATE orders
		SET product_sku=?, product_cost=?, duration=?, payment_id=?, is_paid=?
		WHERE order_id=?
	`
	if _, err := repo.db.Exec(
		query,
		entity.ProductSKU,
		entity.ProductCost,
		entity.Duration,
		entity.PaymentID,
		entity.IsPaid,
		entity.OrderID,
	); err != nil {
		return err
	}
	return nil
}
