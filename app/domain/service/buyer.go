package service

import (
	"fmt"
	"sales/app/domain/repository"
)

type argError struct {
	arg  int
	prob string
}

type BuyerService struct {
	repository repository.BuyerRepository
}

func NewBuyerService(repo repository.BuyerRepository) *BuyerService {
	return &BuyerService{
		repository: repo,
	}
}

func (s *BuyerService) IsDuplicated(date string) (bool, error) {
	buyers, err := s.repository.FindBuyersByDate(date)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error check buyers duplicate")
	}
	if len(buyers) != 0 {
		fmt.Printf("buyers in date: %s already exist\n", date)
		return true, nil
	}
	return false, nil
}
