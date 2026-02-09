package main

import (
	"fmt"
	"net"
	
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Read responses in a goroutine
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Server closed:", err)
				return
			}
			fmt.Println("Received from server:", string(buffer[:n]))
			
		}
	}()

	// Write to server
	for {
		var input string
		fmt.Println("Enter message: ")
		fmt.Scanln(&input)
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}
