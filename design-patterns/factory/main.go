package main

import (
	"log"
	"factory/notification"
)

func main(){
	notifier, err := notification.NewNotification(notification.SMS,"tejas1@gmail.com")
	if err !=nil{
		log.Fatal(err)
	}
	err = notifier.Send("Welcome to our platform")
	if err!=nil{
		log.Fatal(err)
	}
}

