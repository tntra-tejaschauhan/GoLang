package main

import (
	"fmt"
	"time"
)

func add() {
	fmt.Println("add herere")
	time.Sleep(time.Second * 2)
	fmt.Println("add")
}
func main() {
	message := make(chan string)
	go func() { // sending goroutine
		fmt.Println("hi tejas")
		time.Sleep(time.Second * 2)
		message <- "Hello from goroutine!" // send message to channel
		fmt.Println("sent the message......")
	}()
	go add()
	fmt.Println("waiting for message......")
	time.Sleep(time.Second * 10)
	msg, err := <-message // receiving in main goroutine
	fmt.Println(msg, err)
	fmt.Println("main endeed......")
}
