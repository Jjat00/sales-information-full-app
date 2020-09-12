package usecase

import (
	"errors"
	"fmt"
	"sales/app/domain/model"
	"sales/app/domain/repository"
	"sales/app/domain/service"
	"sales/app/infrastructure/acquisition"
)

var (
	errorListProducts     = errors.New("error list products")
	errorRegisterProducts = errors.New("error register products")
)

type ProductUsecase interface {
	ListProducts() ([]*Product, error)
	RegisterProducts(acq *acquisition.Acquisition) error
}

type productUsecase struct {
	repository repository.ProductRepository
	service    *service.ProductService
}

func NewProductUsecase(r repository.ProductRepository, s *service.ProductService) *productUsecase {
	return &productUsecase{
		repository: r,
		service:    s,
	}
}

func (p *productUsecase) ListProducts() ([]*Product, error) {
	products, err := p.repository.GetProducts()
	if err != nil {
		fmt.Println(err)
		return nil, errorListProducts
	}
	newProducts := toProduct(products)
	fmt.Println("list products ok")
	return newProducts, nil
}

func (p *productUsecase) RegisterProducts(acq *acquisition.Acquisition) error {
	res, err := p.service.IsDuplicated(acq.GetDate())
	if err != nil {
		fmt.Println(err)
		return errorRegisterProducts
	}
	if res {
		return errorRegisterProducts
	}
	products := acq.GetProducts()
	err = p.repository.AddProducts(products)
	if err != nil {
		fmt.Println(err)
		return errorRegisterProducts
	}
	fmt.Println("register products ok")
	return nil
}

type Product struct {
	ProductId  string `json:"productId,omitempty"`
	Name       string `json:"name,omitempty"`
	Price      uint16 `json:"price,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

func toProduct(products []*model.Product) []*Product {
	res := make([]*Product, len(products))
	for i, product := range products {
		res[i] = &Product{
			ProductId:  product.GetProductId(),
			Name:       product.GetName(),
			Price:      product.GetPrice(),
			Created_at: product.GetCreate(),
		}
	}
	return res
}
