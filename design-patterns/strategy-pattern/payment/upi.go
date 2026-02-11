package payment

import "fmt"

// upi payment

type UpiPayment struct{}

func (u *UpiPayment) Pay(amount float64) error {
	fmt.Printf("amount %.2f pay with upi\n", amount)
	return nil
}
