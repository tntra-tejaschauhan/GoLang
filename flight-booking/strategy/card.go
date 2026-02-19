package strategy

import "fmt"

type Card struct{}

func (c *Card)Pay (amount float64) bool{
	fmt.Println("processing payment with card",amount)
	return true
} 
