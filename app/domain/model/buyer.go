package model

import (
	"encoding/json"
	"fmt"
)

type Buyer struct {
	BuyerId    string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Age        uint8  `json:"age,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

func NewBuyer(id string, name string, age uint8, date string) *Buyer {
	return &Buyer{
		BuyerId:    id,
		Name:       name,
		Age:        age,
		Created_at: date,
	}
}

func NewBuyers(listBuyers []byte) ([]*Buyer, error) {
	var data struct {
		Buyers []*Buyer `json:"buyers,omitempty"`
	}
	if err := json.Unmarshal(listBuyers, &data); err != nil {
		fmt.Println(err)
		return data.Buyers, fmt.Errorf("error create buyers")
	}
	return data.Buyers, nil
}

func (buyer *Buyer) SetCreateAt(date string) {
	buyer.Created_at = date
}

func (buyer *Buyer) GetBuyerId() string {
	return buyer.BuyerId
}

func (buyer *Buyer) GetName() string {
	return buyer.Name
}

func (buyer *Buyer) GetAge() uint8 {
	return buyer.Age
}

func (buyer *Buyer) GetCreate() string {
	return buyer.Created_at
}
