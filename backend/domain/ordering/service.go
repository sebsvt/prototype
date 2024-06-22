package ordering

type OrderCreated struct {
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
}

type OrderBase struct {
	Reference  string  `json:"reference"`
	CustomerID string  `json:"customer_id"`
	ProductID  string  `json:"product_id"`
	Price      float64 `json:"price"`
	IsPaid     bool    `json:"is_paid"`
	CreatedAt  string  `jons:"created_at"`
}

type OrderService interface {
	CreateNewOrder(new_order OrderCreated) (OrderBase, error)
	GetOrderByOrderReference(reference string) (OrderBase, error)
}
