package payment

import "fmt"

//Card payment

type CardPayment struct{}

func (c *CardPayment) Pay(amount float64) error {
	fmt.Printf("amount %.2f pay with card\n", amount)
	return nil
}
