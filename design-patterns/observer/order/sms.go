package order

import "fmt"

type SMSNotifier struct{}

func (s *SMSNotifier) Update(orderID string){
	fmt.Println("sms sent for order: ",orderID)
}