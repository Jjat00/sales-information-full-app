package registry

import (
	"sales/app/domain/service"
	"sales/app/infrastructure/storage"
	"sales/app/services"
	"sales/app/usecase"
)

type Container struct {
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) BuidBuyer(db *storage.Storage) *services.BuyerService {
	buyerRepo := storage.NewBuyerRepository(db)
	service := service.NewBuyerService(buyerRepo)

	buyerUseCase := usecase.NewBuyerUsecase(buyerRepo, service)
	buyerService := services.NewBuyerService(buyerUseCase)
	return buyerService
}

func (c *Container) BuidProduct(db *storage.Storage) *services.ProductService {
	productRepo := storage.NewProductRepository(db)
	service := service.NewProductService(productRepo)

	productUseCase := usecase.NewProductUsecase(productRepo, service)
	productService := services.NewProductService(productUseCase)

	return productService
}

func (c *Container) BuidTransaction(db *storage.Storage) *services.TransactionService {
	transactionRepo := storage.NewTransactionRepository(db)
	service := service.NewTransactionService(transactionRepo)

	transactionUseCase := usecase.NewTransactionUsecase(transactionRepo, service)
	transactionService := services.NewTransactionService(transactionUseCase)
	return transactionService
}

func (c *Container) BuidConsultBuyer(db *storage.Storage) *services.ConsultBuyerService {
	bRepo := storage.NewBuyerRepository(db)
	pRepo := storage.NewProductRepository(db)
	tRepo := storage.NewTransactionRepository(db)

	consultUsecase := usecase.NewConsultBuyer(bRepo, pRepo, tRepo)
	consultBuyerService := services.NewConsultBuyerService(consultUsecase)
	return consultBuyerService
}
