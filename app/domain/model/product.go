package model

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ProductId  string `json:"productId,omitempty"`
	Name       string `json:"name,omitempty"`
	Price      uint16 `json:"price,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

func NewProduct(id string, name string, price uint16, date string) *Product {
	return &Product{
		ProductId:  id,
		Name:       name,
		Price:      price,
		Created_at: date,
	}
}

func NewProducts(listProducts []byte) ([]*Product, error) {
	var data struct {
		Products []*Product `json:"products,omitempty"`
	}
	if err := json.Unmarshal(listProducts, &data); err != nil {
		fmt.Println(err)
		return data.Products, fmt.Errorf("error create products")
	}
	return data.Products, nil
}

func (product *Product) SetCreateAt(date string) {
	product.Created_at = date
}

func (product *Product) GetProductId() string {
	return product.ProductId
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) GetPrice() uint16 {
	return product.Price
}

func (product *Product) GetCreate() string {
	return product.Created_at
}
