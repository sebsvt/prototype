package service

type ProductResponse struct {
	ProductID   int     `json:"product_id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"is_available"`
}

type ProductRequest struct {
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductService interface {
	CreateNewProduct(ProductRequest) (int, error)
	GetProductByID(product_id int) (*ProductResponse, error)
	GetAllProducts() ([]ProductResponse, error)
}
