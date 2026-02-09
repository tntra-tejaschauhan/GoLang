package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tcp-adapter/pkg/adapter"
)

func main() {
	// Create TCP adapter on localhost:8080
	tcpAdapter := adapter.NewTCPAdapter("localhost", 8080)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := tcpAdapter.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-sigChan
	log.Println("\nReceived shutdown signal")
	
	if err := tcpAdapter.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}
	
	log.Println("Server stopped gracefully")
}
