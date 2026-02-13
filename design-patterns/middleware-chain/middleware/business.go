package middleware

import "fmt"

type BusinessHandler struct{
	next Handler
}

func(b *BusinessHandler) SetNext(handler Handler){
	b.next= handler
}

func(b *BusinessHandler) Handle(request string){
	fmt.Println("Processing business logic for: ",request)
}
