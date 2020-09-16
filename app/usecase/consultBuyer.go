package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"sales/app/domain/repository"
)

var (
	errorGetPurchase     = errors.New("error get purchase histoy")
	errorBuyerSameIp     = errors.New("error get buyers same ip address")
	errorRecommendations = errors.New("error get product recomendations")
)

type ConsultBuyerUsecase interface {
	GetPurchaseHistory(buyerId string) ([]*Product, error)
	GetBuyersSameIP(buyerId string) ([]*Buyer, error)
	GetRecommendations(buyers []*Buyer) ([]*Product, error)
	GetBuyerInformation(buyerId string) *buyerInformation
}

type consultbuyerUsecase struct {
	repoBuyer       repository.BuyerRepository
	repoProduct     repository.ProductRepository
	repoTransaction repository.TransactionRepository
}

func NewConsultBuyer(rb repository.BuyerRepository, rp repository.ProductRepository, rt repository.TransactionRepository) *consultbuyerUsecase {
	return &consultbuyerUsecase{
		repoBuyer:       rb,
		repoProduct:     rp,
		repoTransaction: rt,
	}
}

func (cb *consultbuyerUsecase) GetPurchaseHistory(buyerId string) ([]*Product, error) {
	var purchasedProducts []*Product
	transactions, err := cb.repoTransaction.GetTransactionsByBuyerId(buyerId)
	if err != nil {
		fmt.Println(err)
		return nil, errorGetPurchase
	}
	for _, eachTransaction := range transactions {
		for _, eachProductId := range eachTransaction.PurchasedProductosIds {
			auxProduct, err := cb.repoProduct.GetProductByProductsId(eachProductId)
			if err != nil {
				fmt.Println(err)
			}
			product := toProduct(auxProduct)[0]
			purchasedProducts = append(purchasedProducts, product)
		}
	}
	return purchasedProducts, nil
}

func (cb *consultbuyerUsecase) GetBuyersSameIP(buyerId string) ([]*Buyer, error) {
	var buyers []*Buyer
	transactions, err := cb.repoTransaction.GetTransactionsByBuyerId(buyerId)
	if err != nil {
		fmt.Println(err)
		return nil, errorBuyerSameIp
	}
	NoTransaction := rand.Intn(len(transactions))
	transaction := transactions[NoTransaction]
	ipAddress := transaction.GetIpAddress()
	auxTransactions, err := cb.repoTransaction.GetTransactionsByIpAddress(ipAddress)
	if err != nil {
		fmt.Println(err)
		return nil, errorBuyerSameIp
	}
	for _, eachAuxTransaction := range auxTransactions {
		buyerId := eachAuxTransaction.GetBuyerId()
		auxBuyer, err := cb.repoBuyer.GetBuyersByBuyerId(buyerId)
		if err != nil {
			fmt.Println(err)
			return nil, errorBuyerSameIp
		}
		buyer := toBuyer(auxBuyer)[0]
		buyers = append(buyers, buyer)
	}
	return buyers, nil
}

func (cb *consultbuyerUsecase) GetRecommendations(buyers []*Buyer) ([]*Product, error) {
	var recommendedProducts []*Product
	for index := 0; index < 2; index++ {
		buyerProductHistory, err := cb.GetPurchaseHistory(buyers[index].BuyerId)
		if err != nil {
			fmt.Println(err)
			return nil, errorRecommendations
		}
		for _, eachProduct := range buyerProductHistory {
			recommendedProducts = append(recommendedProducts, eachProduct)
		}
	}
	return recommendedProducts, nil
}

type buyerInformation struct {
	Buyer *Buyer `json:"buyer,omitempty"`
	PurchaseHistory []*Product `json:"purchaseHistory,omitempty"`
	OtherBuyers     []*Buyer   `json:"otherBuyers,omitempty"`
	Recomendations  []*Product `json:"recomendations,omitempty"`
}

func (cb *consultbuyerUsecase) GetBuyerInformation(buyerId string) *buyerInformation {
	buyerInformation := &buyerInformation{}
	buyer, _ := cb.repoBuyer.GetBuyersByBuyerId(buyerId)
	buyerInformation.Buyer = toBuyer(buyer)[0]
	buyerInformation.PurchaseHistory, _ = cb.GetPurchaseHistory(buyerId)
	buyerInformation.OtherBuyers, _ = cb.GetBuyersSameIP(buyerId)
	buyerInformation.Recomendations, _ = cb.GetRecommendations(buyerInformation.OtherBuyers)
	return buyerInformation
}
