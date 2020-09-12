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
	errorListTransactions     = errors.New("error list transactions")
	errorRegisterTransactions = errors.New("error register transactions")
)

type TransactionUsecase interface {
	ListTransactions() ([]*Transaction, error)
	RegisterTransactions(acq *acquisition.Acquisition) error
}

type transactionUsecase struct {
	repository repository.TransactionRepository
	service    *service.TransactionService
}

func NewTransactionUsecase(r repository.TransactionRepository, s *service.TransactionService) *transactionUsecase {
	return &transactionUsecase{
		repository: r,
		service:    s,
	}
}

func (t *transactionUsecase) ListTransactions() ([]*Transaction, error) {
	transactions, err := t.repository.GetTransactions()
	if err != nil {
		fmt.Println(err)
		return nil, errorListTransactions
	}
	newTransactions := toTransaction(transactions)
	fmt.Println("list transactions ok")
	return newTransactions, nil
}

func (t *transactionUsecase) RegisterTransactions(acq *acquisition.Acquisition) error {
	res, err := t.service.IsDuplicated(acq.GetDate())
	if err != nil {
		fmt.Println(err)
		return errorRegisterTransactions
	}
	if res {
		return errorRegisterTransactions
	}
	transactions := acq.GetTransactions()
	if err := t.repository.AddTransactions(transactions); err != nil {
		fmt.Println(err)
		return errorRegisterTransactions
	}
	fmt.Println("register transactions Ok")
	return nil
}

type Transaction struct {
	TransactionId         string   `json:"transactionId,omitempty"`
	BuyerId               string   `json:"buyerId,omitempty"`
	IpAddress             string   `json:"ipAddress,omitempty"`
	Device                string   `json:"device,omitempty"`
	PurchasedProductosIds []string `json:"purchasedProductosIds,omitempty"`
	Created_at            string   `json:"created_at,omitempty"`
}

func toTransaction(transactions []*model.Transaction) []*Transaction {
	res := make([]*Transaction, len(transactions))
	for i, transaction := range transactions {
		res[i] = &Transaction{
			TransactionId:         transaction.GetTransactionId(),
			BuyerId:               transaction.GetBuyerId(),
			IpAddress:             transaction.GetIpAddress(),
			Device:                transaction.GetDevice(),
			PurchasedProductosIds: transaction.GetPurchasedProductosIds(),
			Created_at:            transaction.GetDate(),
		}
	}
	return res
}
