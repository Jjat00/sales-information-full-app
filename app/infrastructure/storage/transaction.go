package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sales/app/domain/model"
)

var (
	errorGetTransactions  = errors.New("error get transactions database")
	errorAddTransactions  = errors.New("error add transactions database")
	errorFindTransactions = errors.New("error find transactions by date database")
)

type transactionRepository struct {
	db *Storage
}

func NewTransactionRepository(db *Storage) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetTransactions() ([]*model.Transaction, error) {
	query := `
		{
			transactions(func: has(transactionId)) {
				transactionId
				buyerId
				ipAddress
				device
				purchasedProductosIds
			}
		}
	`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorGetTransactions
	}
	transactions, err := model.NewTransactions(response)
	if err != nil {
		log.Println(err)
		return nil, errorGetTransactions
	}
	return transactions, nil
}

func (r *transactionRepository) AddTransactions(transactions []*model.Transaction) error {
	response, err := json.Marshal(&transactions)
	if err != nil {
		fmt.Println(err)
		return errorAddTransactions
	}
	err = r.db.RunMutation(response)
	if err != nil {
		fmt.Println(err)
		return errorAddTransactions
	}
	return nil
}

func (r *transactionRepository) FindTransactionsByDate(date string) ([]*model.Transaction, error) {
	query := `
		{
			transactions(func: eq(created_at,` + date + `)) @filter(has(ipAddress)){
				transactionId
			}
		}
	`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindTransactions
	}
	transactions, err := model.NewTransactions(response)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindTransactions
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransactionsByBuyerId(buyerId string) ([]*model.Transaction, error) {
	query := `
		{		
			transactions(func: eq(buyerId,` + buyerId + `)){
				buyerId
				purchasedProductosIds
				ipAddress
			}
		}
	`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindTransactions
	}
	transactions, err := model.NewTransactions(response)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindTransactions
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransactionsByIpAddress(ipAddress string) ([]*model.Transaction, error) {
	query := `
				{
					transactions(func: eq(ipAddress,` + ipAddress + `)){
						buyerId
						purchasedProductosIds
						ipAddress
          			}
				}			
			`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindTransactions
	}
	transactions, err := model.NewTransactions(response)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindTransactions
	}
	return transactions, nil
}
