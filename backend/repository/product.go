package repository

type Product struct {
	ProductID   int     `db:"product_id"`
	Sku         string  `db:"sku"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	IsAvailable bool    `db:"is_available"`
}

type ProductRepository interface {
	Create(Product) (int, error)
	List() ([]Product, error)
	FromID(product_id int) (*Product, error)
}
