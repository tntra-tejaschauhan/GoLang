package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"tcp-adapter/pkg/protocol"
)

func main() {
	// Connect to TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to server at localhost:8080")

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	stdinReader := bufio.NewReader(os.Stdin)

	// Read welcome message
	welcomeMsg, err := protocol.Decode(reader)
	if err != nil {
		log.Fatalf("Error reading welcome message: %v", err)
	}
	fmt.Printf("Server: [%s] %s\n\n", welcomeMsg.Command, welcomeMsg.Payload)

	// Display available commands
	printHelp()

	// Main client loop
	for {
		fmt.Print("> ")
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		// Parse command and payload
		parts := strings.SplitN(input, " ", 2)
		command := strings.ToUpper(parts[0])
		payload := ""
		if len(parts) > 1 {
			payload = parts[1]
		}

		// Special handling for local commands
		if command == "HELP" {
			printHelp()
			continue
		}

		// Create and send message
		msg := protocol.NewMessage(command, payload)
		_, err = writer.WriteString(msg.Encode())
		if err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}
		writer.Flush()

		// Read response
		response, err := protocol.Decode(reader)
		if err != nil {
			log.Printf("Error reading response: %v", err)
			break
		}

		fmt.Printf("Server: [%s] %s\n", response.Command, response.Payload)

		// Exit if QUIT command
		if command == "QUIT" {
			log.Println("Disconnecting...")
			break
		}
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  ECHO <text>     - Echo back the text")
	fmt.Println("  UPPER <text>    - Convert text to uppercase")
	fmt.Println("  LOWER <text>    - Convert text to lowercase")
	fmt.Println("  REVERSE <text>  - Reverse the text")
	fmt.Println("  PING            - Check server status")
	fmt.Println("  QUIT            - Disconnect from server")
	fmt.Println("  HELP            - Show this help message")
	fmt.Println()
}
