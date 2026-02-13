package middleware

type Handler interface{
	SetNext(handler Handler)
	Handle(request string)
}