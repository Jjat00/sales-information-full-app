package repository

import "sales/app/domain/model"

type ProductRepository interface {
	GetProducts() ([]*model.Product, error)
	FindProductsByDate(date string) ([]*model.Product, error)
	AddProducts(product []*model.Product) error
	GetProductByProductsId(productId string) ([]*model.Product, error)
}
