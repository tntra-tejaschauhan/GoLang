package middleware

import "fmt"

type LoggingHandler struct{
  next Handler
}

func (l *LoggingHandler) SetNext(handler Handler){
	l.next= handler
}

func(l *LoggingHandler) Handle(request string){
	fmt.Println("Logging request:",request)

	if l.next !=nil{
		l.next.Handle(request)
	}
}