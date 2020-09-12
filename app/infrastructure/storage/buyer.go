package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"sales/app/domain/model"
)

var (
	errorGetBuyers  = errors.New("error get buyers database")
	errorAddBuyers  = errors.New("error add buyers database")
	errorFindBuyers = errors.New("error find buyers by date database")
)

type buyerRepository struct {
	db *Storage
}

func NewBuyerRepository(db *Storage) *buyerRepository {
	return &buyerRepository{
		db: db,
	}
}

func (r *buyerRepository) GetBuyers() ([]*model.Buyer, error) {
	query := `
		{
			buyers(func: has(id)) {
				id
				age
				name
			}
		}
	`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorGetBuyers
	}
	buyers, err := model.NewBuyers(response)
	if err != nil {
		fmt.Println(err)
		return nil, errorGetBuyers
	}
	return buyers, nil
}

func (r *buyerRepository) AddBuyers(buyers []*model.Buyer) error {
	response, err := json.Marshal(&buyers)
	if err != nil {
		fmt.Println(err)
		return errorAddBuyers
	}
	err = r.db.RunMutation(response)
	if err != nil {
		fmt.Println(err)
		return errorAddBuyers
	}
	return nil
}

func (r *buyerRepository) FindBuyersByDate(date string) ([]*model.Buyer, error) {
	query := `
		{
			buyers(func: eq(created_at,` + date + `)) @filter(has(id)){
				name
			}
		}
	`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindBuyers
	}
	buyers, err := model.NewBuyers(response)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindBuyers
	}
	return buyers, nil
}

func (r *buyerRepository) GetBuyersByBuyerId(buyerId string) ([]*model.Buyer, error) {
	query := `
				{
					buyers(func: eq(id,` + buyerId + `)){
						id
						name
						age
					}
				}			
			`
	response, err := r.db.RunQuery(query)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindBuyers
	}
	buyer, err := model.NewBuyers(response)
	if err != nil {
		fmt.Println(err)
		return nil, errorFindBuyers
	}
	return buyer, nil
}

func (r *buyerRepository) DeleteBuyers() error {
	err := r.db.DeleteAllPredicate()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
