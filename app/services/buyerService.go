package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"sales/app/infrastructure/acquisition"
	"sales/app/usecase"
	"sync"
)

var (
	errorListBuyers     = errors.New("error buyer service list buyers")
	errorRegisterBuyers = errors.New("error buyer service register buyers")
)

type BuyerService struct {
	buyerUseCase usecase.BuyerUsecase
}

func NewBuyerService(buyerUseCase usecase.BuyerUsecase) *BuyerService {
	return &BuyerService{
		buyerUseCase: buyerUseCase,
	}
}

func (buyer *BuyerService) ListBuyers() ([]byte, error) {
	buyers, err := buyer.buyerUseCase.ListBuyers()
	if err != nil {
		fmt.Println(err)
		return nil, errorListBuyers
	}
	response, err := json.Marshal(buyers)
	if err != nil {
		fmt.Println(err)
		return nil, errorListBuyers
	}
	return response, nil
}

func (buyer *BuyerService) RegisterBuyers(timestamp string, wg *sync.WaitGroup) error {
	defer wg.Done()
	acq := acquisition.NewAcquisition(timestamp)
	err := buyer.buyerUseCase.RegisterBuyers(acq)
	if err != nil {
		fmt.Println(err)
		return errorRegisterBuyers
	}
	return nil
}

func (buyer *BuyerService) DeleteBuyers() error {
	err := buyer.buyerUseCase.DeleteBuyers()
	if err != nil {
		return err
	}
	return nil
}
