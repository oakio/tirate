package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shopspring/decimal"
)

type envelope struct {
	ResultCode string
	Payload    payload
}

type payload struct {
	Rates []rate
}

type rate struct {
	Category     string
	FromCurrency currency
	ToCurrency   currency
	Buy          decimal.Decimal
	Sell         decimal.Decimal
}

type currency struct {
	Name string
}

func main() {
	httpResp, httpErr := http.Get("https://api.tinkoff.ru/v1/currency_rates")
	checkErr(httpErr)

	httpBody := httpResp.Body
	defer httpBody.Close()

	envelop := new(envelope)

	decoder := json.NewDecoder(httpBody)
	decodeErr := decoder.Decode(envelop)
	checkErr(decodeErr)

	if envelop.ResultCode != "OK" {
		log.Fatal("Rates loading failed. ResultCode = ", envelop.ResultCode)
	}

	for _, rate := range envelop.Payload.Rates {

		fmt.Printf("%s%s %v %v %s\n",
			rate.FromCurrency.Name,
			rate.ToCurrency.Name,
			rate.Buy.StringFixed(2),
			rate.Sell.StringFixed(2),
			rate.Category)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
