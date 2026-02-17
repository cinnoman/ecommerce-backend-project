package product

import (
	"ecommerce/domain"
	prdctHandler "ecommerce/rest/handlers/product"
)

type Service interface {
	prdctHandler.Service
}

type ProductRepo interface {
	Get(pId int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Create(newProduct domain.Product) (*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(pId int) error
}
