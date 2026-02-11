package order

type Order struct{
	observers []Observer
}

func NewOrder() *Order{
	return &Order{}
}

func(o *Order) Register(observer Observer){
	o.observers=append(o.observers, observer)
}

func(o *Order) Notify(orderID string){
	for _,observer := range o.observers{
		observer.Update(orderID)
	}
}

func (o *Order) Create(orderID string){
	o.Notify(orderID)
}