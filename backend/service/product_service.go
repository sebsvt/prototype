package service

import "github.com/sebsvt/prototype/repository"

type productService struct {
	product_repo repository.ProductRepository
}

func NewProductService(product_repo repository.ProductRepository) ProductService {
	return productService{product_repo: product_repo}
}

// CreateNewProduct implements ProductService.
func (srv productService) CreateNewProduct(entity ProductRequest) (int, error) {
	query_product := repository.Product{
		Sku:         entity.Sku,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		IsAvailable: true,
	}
	product_id, err := srv.product_repo.Create(query_product)
	if err != nil {
		return 0, err
	}
	return product_id, nil
}

// GetAllProducts implements ProductService.
func (srv productService) GetAllProducts() ([]ProductResponse, error) {
	var products []ProductResponse
	query_products, err := srv.product_repo.List()
	if err != nil {
		return nil, err
	}
	for _, pro := range query_products {
		products = append(products, ProductResponse(pro))
	}
	return products, nil
}

// GetProductByID implements ProductService.
func (srv productService) GetProductByID(product_id int) (*ProductResponse, error) {
	product, err := srv.product_repo.FromID(product_id)
	if err != nil {
		return nil, err
	}
	return (*ProductResponse)(product), nil
}
