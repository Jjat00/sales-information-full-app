/*
@File: server.go
@Author: Jaime Aza
@email: userjjat00@gmail.com
*/

package main

import (
	"fmt"
	"net/http"
	"sales/app/infrastructure/storage"
	"sales/app/registry"
	"sales/config"
	"sales/handlers"

	"github.com/go-chi/chi"
)

func main() {

	db := storage.NewStorage(config.GetConfig().DB)
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	cnt := registry.NewContainer()
	buyerServices := cnt.BuidBuyer(db)
	productServices := cnt.BuidProduct(db)
	transactionServices := cnt.BuidTransaction(db)
	consultBuyerServices := cnt.BuidConsultBuyer(db)

	handler := handlers.NewHandlers(buyerServices, productServices,
		transactionServices, consultBuyerServices)

	r := chi.NewRouter()
	r.Post("/loadData", handler.LoadData)
	r.Get("/buyers", handler.ListAllBuyers)
	r.Get("/consultBuyer/{buyerId}", handler.ConsultBuyer)
	r.Delete("/delete", handler.DeleteData)

	port := config.GetConfig().Server.Port
	err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println(err)
	}
}
