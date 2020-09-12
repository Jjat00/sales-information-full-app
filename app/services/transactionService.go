package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"sales/app/infrastructure/acquisition"
	"sales/app/usecase"
	"sync"
)

var (
	errorListTransactions     = errors.New("error transaction service list Transactions")
	errorRegisterTransactions = errors.New("error transaction service register Transactions")
)

type TransactionService struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionService(transactionUsecase usecase.TransactionUsecase) *TransactionService {
	return &TransactionService{
		transactionUsecase: transactionUsecase,
	}
}

func (transaction *TransactionService) ListTransactions() ([]byte, error) {
	transactions, err := transaction.transactionUsecase.ListTransactions()
	if err != nil {
		fmt.Println(err)
		return nil, errorListTransactions
	}
	response, err := json.Marshal(transactions)
	if err != nil {
		fmt.Println(err)
		return nil, errorListTransactions
	}
	return response, nil
}

func (buyer *TransactionService) RegisterTransactions(timestamp string, wg *sync.WaitGroup) error {
	defer wg.Done()
	acq := acquisition.NewAcquisition(timestamp)
	err := buyer.transactionUsecase.RegisterTransactions(acq)
	if err != nil {
		fmt.Println(err)
		return errorRegisterTransactions
	}
	return nil
}
