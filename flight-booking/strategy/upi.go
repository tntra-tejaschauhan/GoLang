package strategy

import (
	"fmt"
	"math/rand"
	"time"
)

type UPI struct{}

func (u *UPI) Pay(amount float64) bool {
	fmt.Println("Processing payment...")

	sleepTime := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(sleepTime)

	fmt.Println("Payment finished in", sleepTime)
	return true
}
