package middleware

import "fmt"

type AuthHandler struct{
	next Handler
}

func(a *AuthHandler) SetNext(handler Handler){
	a.next=handler
}

func(a *AuthHandler) Handle(request string){
	if request=="invalid"{
		fmt.Println("Authentication failed")
		return
	}
	fmt.Println("Authentication Successful")
	if a.next !=nil{
		a.next.Handle(request)
	}
}