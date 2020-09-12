package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sales/app/infrastructure/acquisition"
	"sales/app/usecase"
	"sync"
)

var (
	errorListProducts     = errors.New("error product service list products")
	errorRegisterProducts = errors.New("error product service register products")
)

type ProductService struct {
	productUsecase usecase.ProductUsecase
}

func NewProductService(productUsecase usecase.ProductUsecase) *ProductService {
	return &ProductService{
		productUsecase: productUsecase,
	}
}

func (product *ProductService) ListProducts() ([]byte, error) {
	products, err := product.productUsecase.ListProducts()
	if err != nil {
		fmt.Println(err)
		return nil, errorListProducts
	}
	response, err := json.Marshal(products)
	if err != nil {
		log.Fatal(err)
		return nil, errorListProducts
	}
	return response, nil
}

func (buyer *ProductService) RegisterProducts(timestamp string, wg *sync.WaitGroup) error {
	defer wg.Done()
	acq := acquisition.NewAcquisition(timestamp)
	err := buyer.productUsecase.RegisterProducts(acq)
	if err != nil {
		fmt.Println(err)
		return errorRegisterProducts
	}
	return nil
}
