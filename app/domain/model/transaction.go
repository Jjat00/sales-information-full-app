package model

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	TransactionId         string   `json:"transactionId,omitempty"`
	BuyerId               string   `json:"buyerId,omitempty"`
	IpAddress             string   `json:"ipAddress,omitempty"`
	Device                string   `json:"device,omitempty"`
	PurchasedProductosIds []string `json:"purchasedProductosIds,omitempty"`
	Created_at            string   `json:"created_at,omitempty"`
}

func NewTransaction(tId string, bId string, ip string, dev string, pIds []string, date string) *Transaction {
	return &Transaction{
		TransactionId:         tId,
		BuyerId:               bId,
		IpAddress:             ip,
		Device:                dev,
		PurchasedProductosIds: pIds,
		Created_at:            date,
	}
}

func NewTransactions(listProducts []byte) ([]*Transaction, error) {
	var data struct {
		Transactions []*Transaction `json:"transactions,omitempty"`
	}
	if err := json.Unmarshal(listProducts, &data); err != nil {
		fmt.Println("error create transactions", err)
		return data.Transactions, fmt.Errorf("error create transactions")
	}
	return data.Transactions, nil
}

func (transaction *Transaction) SetCreateAt(date string) {
	transaction.Created_at = date
}

func (transaction *Transaction) GetTransactionId() string {
	return transaction.TransactionId
}
func (transaction *Transaction) GetBuyerId() string {
	return transaction.BuyerId
}

func (transaction *Transaction) GetIpAddress() string {
	return transaction.IpAddress
}

func (transaction *Transaction) GetDevice() string {
	return transaction.Device
}

func (transaction *Transaction) GetPurchasedProductosIds() []string {
	return transaction.PurchasedProductosIds
}

func (transaction *Transaction) GetDate() string {
	return transaction.Created_at
}
