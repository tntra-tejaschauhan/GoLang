package order

import "fmt"

type AnalyticsService struct{}

func(a *AnalyticsService) Update(orderID string){
	fmt.Println("Analytics update for order:", orderID)
}