package payment


type PaymentService struct{
	strategy PaymentStrategy
}

func NewPaymentService(strategy PaymentStrategy) *PaymentService {
	return &PaymentService{strategy: strategy}
}

func (p *PaymentService) SetStrategy(strategy PaymentStrategy){
	p.strategy=strategy
	
}

func(p *PaymentService) Process(amount float64) error{
	return p.strategy.Pay(amount)
}