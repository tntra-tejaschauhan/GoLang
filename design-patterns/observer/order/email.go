package order

import "fmt"

type EmailNotifier struct{}

func(e *EmailNotifier) Update(orderID string){
	fmt.Println("Email sent for order: ",orderID)
}
