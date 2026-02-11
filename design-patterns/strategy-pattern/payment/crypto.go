package payment

import "fmt"

// Crypto payment

type CryptoPayment struct{}

func (c *CryptoPayment) Pay(amount float64) error {
	fmt.Printf("amount %.2f pay with crypto\n", amount)
	return nil
}
