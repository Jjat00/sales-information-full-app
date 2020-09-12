package repository

import "sales/app/domain/model"

type BuyerRepository interface {
	GetBuyers() ([]*model.Buyer, error)
	FindBuyersByDate(date string) ([]*model.Buyer, error)
	AddBuyers(buyers []*model.Buyer) error
	GetBuyersByBuyerId(buyerId string) ([]*model.Buyer, error)
	DeleteBuyers() error
}
