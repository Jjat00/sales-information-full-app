package usecase

import (
	"errors"
	"fmt"
	"sales/app/domain/model"
	"sales/app/domain/repository"
	"sales/app/domain/service"
	"sales/app/infrastructure/acquisition"
)

var (
	errorListBuyers     = errors.New("error list buyers")
	errorRegisterBuyers = errors.New("error register buyers")
)

type BuyerUsecase interface {
	ListBuyers() ([]*Buyer, error)
	RegisterBuyers(acq *acquisition.Acquisition) error
	DeleteBuyers() error
}

type buyerUsecase struct {
	repository repository.BuyerRepository
	service    *service.BuyerService
}

func NewBuyerUsecase(r repository.BuyerRepository, s *service.BuyerService) *buyerUsecase {
	return &buyerUsecase{
		repository: r,
		service:    s,
	}
}

func (b *buyerUsecase) ListBuyers() ([]*Buyer, error) {
	buyers, err := b.repository.GetBuyers()
	if err != nil {
		fmt.Println(err)
		return nil, errorListBuyers
	}
	newBuyers := toBuyer(buyers)
	fmt.Println("list buyers ok")
	return newBuyers, nil
}

func (b *buyerUsecase) RegisterBuyers(acq *acquisition.Acquisition) error {
	res, err := b.service.IsDuplicated(acq.GetDate())
	if err != nil {
		fmt.Println(err)
		return errorRegisterBuyers
	}
	if res {
		return errorRegisterBuyers
	}
	buyers, _ := acq.GetBuyers()
	err = b.repository.AddBuyers(buyers)
	if err != nil {
		fmt.Println(err)
		return errorRegisterBuyers
	}
	fmt.Println("register buyers ok")
	return nil
}

func (b *buyerUsecase) DeleteBuyers() error {
	err := b.repository.DeleteBuyers()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

type Buyer struct {
	BuyerId    string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Age        uint8  `json:"age,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

func toBuyer(buyers []*model.Buyer) []*Buyer {
	res := make([]*Buyer, len(buyers))
	for i, buyer := range buyers {
		res[i] = &Buyer{
			BuyerId:    buyer.GetBuyerId(),
			Name:       buyer.GetName(),
			Age:        buyer.GetAge(),
			Created_at: buyer.GetCreate(),
		}
	}
	return res
}
