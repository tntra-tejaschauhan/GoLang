package main

import (
	"log"
	"strategy-pattern/payment"
)

func main() {
	card := &payment.CardPayment{}
	service := payment.NewPaymentService(card)
	service.Process(1000)
	upi := &payment.UpiPayment{}
	service.SetStrategy(upi)
	err := service.Process(2000)
	if err != nil {
		log.Fatal(err)
	}

}
