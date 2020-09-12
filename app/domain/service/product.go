package service

import (
	"fmt"
	"sales/app/domain/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

func (s *ProductService) IsDuplicated(date string) (bool, error) {
	products, err := s.repository.FindProductsByDate(date)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error chek products duplicate")
	}
	if len(products) != 0 {
		fmt.Printf("products in date: %s already exists\n", date)
		return true, nil
	}
	return false, nil
}
