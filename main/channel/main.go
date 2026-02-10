package main

import (
	"fmt"
	"time"
)

func add() {
	time.Sleep(time.Second * 21)
	fmt.Println("add")
}
func main() {
	message := make(chan string)
	go func() { // sending goroutine
		time.Sleep(time.Second * 3 / 10)
		message <- "Hello from goroutine!" // send message to channel
		fmt.Println("sent the message......")
	}()
	// add()
	fmt.Println("waiting for message......")
	msg, err := <-message // receiving in main goroutine
	fmt.Println(msg, err)
	fmt.Println("main endeed......")
}
