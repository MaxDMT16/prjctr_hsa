package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan struct{})

	go listenToInterruptionSig(done)

	go fetchCurrencyRateAndPushToGA()

	<-done
	log.Println("Done")
}

/*
{ 
"r030":978,"txt":"Євро","rate":26.9789,"cc":"EUR","exchangedate":"02.03.2020"
 }
 */
type CurrencyRateResponse struct {
	Rate float64 `json:"rate"`
	Date string `json:"exchangedate"`
	Currency string `json:"cc"`	
}

func fetchCurrencyRateAndPushToGA() {
	log.Println("Fetching currency rate and push to GA...")

	// https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=EUR&date=20200302&json

	resp, err := http.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=EUR&date=20200302&json")
	if err != nil {
		log.Println(fmt.Errorf("get currency rate from bank.gov.ua: %w", err))
		return
	}

	defer resp.Body.Close()

	var currencyRateResponse []CurrencyRateResponse
	err = json.NewDecoder(resp.Body).Decode(&currencyRateResponse)
	if err != nil {
		log.Println(fmt.Errorf("unmarshal response: %w", err))
		return
	}

	log.Println("Currency rates: ", currencyRateResponse)
}

func listenToInterruptionSig(done chan<- struct{}) {
	log.Println("Listening signals...")
	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	<-c
	close(done)
}
