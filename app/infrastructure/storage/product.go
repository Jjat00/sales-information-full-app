package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"sales/app/domain/model"
)

var (
	errorGetProducts  = errors.New("error get products database")
	errorAddProducts  = errors.New("error add products database")
	errorFindProducts = errors.New("error find products by date database")
)

type productRepository struct {
	db *Storage
}

func NewProductRepository(db *Storage) *productRepository {
	return &productRepository{
		db: db,
	}
}
func (r *productRepository) GetProducts() ([]*model.Product, error) {
	query := `
		{
			products(func: has(productId)) {
				name
				price
			}
		}
	`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorGetProducts
	}
	products, err := model.NewProducts(response)
	if err != nil {
		fmt.Println(err)
		return nil, errorGetProducts
	}
	return products, nil
}

func (r *productRepository) AddProducts(products []*model.Product) error {
	response, err := json.Marshal(&products)
	if err != nil {
		fmt.Println(err)
		return errorAddProducts
	}
	err = r.db.RunMutation(response)
	if err != nil {
		fmt.Println(err)
		return errorAddProducts
	}
	return nil
}

func (r *productRepository) FindProductsByDate(date string) ([]*model.Product, error) {
	query := `
		{
			products(func: eq(created_at,` + date + `)) @filter(has(productId)){
				name
			}
		}
	`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindProducts
	}
	products, err := model.NewProducts(response)
	if err != nil {
		return nil, errorFindProducts
	}
	return products, nil
}

func (r *productRepository) GetProductByProductsId(productId string) ([]*model.Product, error) {
	query := `
				{
					products(func: eq(productId,` + productId + `)){
						name
						price
          			}
				}			
			`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindProducts
	}
	products, err := model.NewProducts(response)
	if err != nil {
		return nil, errorFindProducts
	}
	return products, nil
}
