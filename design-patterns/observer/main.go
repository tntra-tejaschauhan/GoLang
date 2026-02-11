package main

import "observer/order"

func main() {
	orderService := order.NewOrder()
	orderService.Register(&order.EmailNotifier{})
	orderService.Register(&order.SMSNotifier{})
	orderService.Register(&order.AnalyticsService{})

	orderService.Create("order-123")
}
