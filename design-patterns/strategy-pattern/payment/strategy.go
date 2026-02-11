package payment

// strategy interface
type PaymentStrategy interface{
	Pay(amount float64) error
}
