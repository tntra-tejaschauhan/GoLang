package middleware

import "fmt"

type AuthorizationHandler struct{
	next Handler
}

func(a *AuthorizationHandler) SetNext(handler Handler){
	a.next=handler
}

func(a *AuthorizationHandler) Handle(request string){
	fmt.Println("Authorization passed")

	if a.next!=nil{
		a.next.Handle(request)
	}
}