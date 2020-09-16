package acquisition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sales/app/domain/model"
	"strconv"
	"strings"
)

type Acquisition struct {
	timestamp string
}

func NewAcquisition(timestamp string) *Acquisition {
	acq := &Acquisition{
		timestamp: timestamp,
	}
	return acq
}

func (acq *Acquisition) GetBuyers() ([]*model.Buyer, error) {
	response := getData("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date=" + acq.timestamp)
	buyers := []*model.Buyer{}
	err := json.Unmarshal(response, &buyers)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error get buyers")
	}
	date := UnixTimeToDateString(acq.timestamp)
	for _, buyer := range buyers {
		buyer.SetCreateAt(date)
	}
	return buyers, nil
}

func (acq *Acquisition) GetProducts() []*model.Product {
	response := getData("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products?date=" + acq.timestamp)
	date := UnixTimeToDateString(acq.timestamp)
	fixNames := strings.ReplaceAll(string(response), "'s", "")
	splitNewLine := strings.Split(fixNames, "\n")
	var products []*model.Product
	for i := 0; i < len(splitNewLine)-1; i++ {
		idNamePrice := strings.Split(splitNewLine[i], "'")
		name := cleanText(idNamePrice[1])
		name = strings.ReplaceAll(name, "\"", "")
		price, _ := strconv.Atoi(idNamePrice[len(idNamePrice)-1])
		product := model.NewProduct(idNamePrice[0], name, uint16(price), date)
		products = append(products, product)
	}
	return products
}

func (acq *Acquisition) GetTransactions() []*model.Transaction {
	response := getData("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date=" + acq.timestamp)
	date := UnixTimeToDateString(acq.timestamp)
	split := strings.Split(string(response), "#")
	var transactions []*model.Transaction
	for _, each := range split {
		re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
		ipAddress := re.FindString(each)
		if re.MatchString(each) {
			idsDeviceProductIds := strings.Split(each, ipAddress)
			transactionID := idsDeviceProductIds[0][0:12]
			buyerID := cleanText(idsDeviceProductIds[0][12:])
			deviceAndProducts := strings.Split(idsDeviceProductIds[1], "(")
			device := cleanText(deviceAndProducts[0])
			productIDs := strings.Split(strings.ReplaceAll(deviceAndProducts[1], ")", ""), ",")
			productIDs = cleanArrayText(productIDs)
			transaction := model.NewTransaction(transactionID, buyerID, ipAddress, device, productIDs, date)
			transactions = append(transactions, transaction)
		}
	}
	return transactions
}

func (acq *Acquisition) GetDate() string {
	date := UnixTimeToDateString(acq.timestamp)
	return date
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}
	return response
}

func cleanText(text string) string {
	newtext := strings.Trim(text, "\x00")
	return newtext
}

func cleanArrayText(text []string) []string {
	var newArrayText []string
	for _, each := range text {
		newArrayText = append(newArrayText, cleanText(each))
	}
	return newArrayText
}
