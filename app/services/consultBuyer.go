package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sales/app/usecase"
	"sync"
)

var (
	errorPurchaseHistory = errors.New("error get purchase history")
)

type ConsultBuyerService struct {
	consultBuyerUc usecase.ConsultBuyerUsecase
}

func NewConsultBuyerService(consultBuyerUc usecase.ConsultBuyerUsecase) *ConsultBuyerService {
	return &ConsultBuyerService{
		consultBuyerUc: consultBuyerUc,
	}
}

func (cbs *ConsultBuyerService) GetPurchaseHistory(buyerId string, wg *sync.WaitGroup) ([]byte, error) {
	defer wg.Done()
	products, err := cbs.consultBuyerUc.GetPurchaseHistory(buyerId)
	if err != nil {
		fmt.Println(err)
		return nil, errorPurchaseHistory
	}
	response, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		return nil, errorPurchaseHistory
	}
	return response, nil
}

func (cbs *ConsultBuyerService) GetBuyersSameIP(buyerId string, wg *sync.WaitGroup) ([]byte, error) {
	defer wg.Done()
	buyers, err := cbs.consultBuyerUc.GetBuyersSameIP(buyerId)
	if err != nil {
		fmt.Println(err)
		return nil, errorPurchaseHistory
	}
	response, err := json.Marshal(buyers)
	if err != nil {
		fmt.Println(err)
		return nil, errorPurchaseHistory
	}
	return response, nil
}

func (cbs *ConsultBuyerService) GetBuyerInformation(buyerId string) []byte {
	buyerInformation := cbs.consultBuyerUc.GetBuyerInformation(buyerId)
	out, err := json.Marshal(buyerInformation)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
