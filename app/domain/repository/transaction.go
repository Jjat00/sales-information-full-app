package repository

import "sales/app/domain/model"

type TransactionRepository interface {
	GetTransactions() ([]*model.Transaction, error)
	FindTransactionsByDate(date string) ([]*model.Transaction, error)
	AddTransactions(transactions []*model.Transaction) error

	GetTransactionsByBuyerId(buyerId string) ([]*model.Transaction, error)
	GetTransactionsByIpAddress(ipAddress string) ([]*model.Transaction, error)
}
