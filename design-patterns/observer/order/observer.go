package order

type Observer interface {
	Update(orderID string)
}