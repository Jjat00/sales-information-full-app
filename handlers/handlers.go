package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sales/app/services"
	"sync"

	"github.com/go-chi/chi"
)

type Handlers struct {
	buyerService       *services.BuyerService
	productService     *services.ProductService
	transactionService *services.TransactionService
	consultBuyer       *services.ConsultBuyerService
}

func NewHandlers(bs *services.BuyerService, ps *services.ProductService,
	ts *services.TransactionService, cb *services.ConsultBuyerService) *Handlers {
	return &Handlers{
		buyerService:       bs,
		productService:     ps,
		transactionService: ts,
		consultBuyer:       cb,
	}
}

type RegisterData struct {
	BuyerResponse       string
	ProductResponse     string
	TransactionResponse string
}

func (h *Handlers) ListAllBuyers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	buyers, err := h.buyerService.ListBuyers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(buyers))
}

func (h *Handlers) LoadData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var wg sync.WaitGroup
	wg.Add(3)
	timestamp := getDataFromTimestamp(r)
	reg := &RegisterData{}
	//register buyers data
	go func() {
		err := h.buyerService.RegisterBuyers(timestamp, &wg)
		if err != nil {
			fmt.Println(err)
			reg.BuyerResponse = "fail buyer registration"
		} else {
			fmt.Println("successfull buyer registration")
			reg.BuyerResponse = "successfull buyer registration"
		}
	}()
	//register products data
	go func() {
		err := h.productService.RegisterProducts(timestamp, &wg)
		if err != nil {
			fmt.Println(err)
			reg.ProductResponse = "fail product registration"
		} else {
			fmt.Println("successfull product registration")
			reg.ProductResponse = "successfull product registration"
		}
	}()
	//regiter transactions data
	go func() {
		err := h.transactionService.RegisterTransactions(timestamp, &wg)
		if err != nil {
			fmt.Println(err)
			reg.TransactionResponse = "fail transaction registration"
		} else {
			fmt.Println("successfull transaction registration")
			reg.TransactionResponse = "successfull transaction registration"
		}
	}()
	wg.Wait()
	response, err := json.Marshal(reg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(response))
}

func (h *Handlers) ConsultBuyer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	buyerId := string(chi.URLParam(r, "buyerId"))
	history := h.consultBuyer.GetBuyerInformation(buyerId)
	fmt.Fprintf(w, string(history))
}

func (h *Handlers) DeleteData(w http.ResponseWriter, r *http.Request) {
	err := h.buyerService.DeleteBuyers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string("successfull delete"))
}

func getDataFromTimestamp(r *http.Request) string {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print(err)
	}
	var timeUnix struct {
		Date string `json:"date,omitempty"`
	}
	err = json.Unmarshal(request, &timeUnix)
	if err != nil {
		fmt.Print(err)
	}
	return timeUnix.Date
}

/* 	var wg sync.WaitGroup
wg.Add(2)
go func() {
	history, err := h.consultBuyer.GetPurchaseHistory(buyerId, &wg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(history))
}()
var sameIps []byte
go func() {
	sameIps, _ = h.consultBuyer.GetBuyersSameIP(buyerId, &wg)
}()
wg.Wait()
fmt.Fprintf(w, "\n\n\n\n\n"+string(sameIps)) */
