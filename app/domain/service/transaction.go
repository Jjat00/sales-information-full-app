package service

import (
	"fmt"
	"sales/app/domain/repository"
)

type TransactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repository: repo,
	}
}

func (s *TransactionService) IsDuplicated(date string) (bool, error) {
	transactions, err := s.repository.FindTransactionsByDate(date)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error chek transactions duplicate")
	}
	/* 	transactions1, err := s.repository.GetTransactionsByBuyerId("2483455e")
	   	for i := 0; i < len(transactions1); i++ {
	   		fmt.Println(*transactions1[i])
	   		ipAddress := transactions1[i].GetIpAddress()
	   		fmt.Println(ipAddress)
	   		transaction, _ := s.repository.GetTransactionsByIpAddress(ipAddress)
	   		fmt.Println(*transaction[i])
	   	} */
	if len(transactions) != 0 {
		fmt.Printf("transactions in date: %s already exists\n", date)
		return true, nil
	}
	return false, nil
}
