package adapter

import (
	"fmt"
	"log"
	"net"
	"tcp-adapter/pkg/handler"
)

// TCPAdapter represents a TCP server adapter
type TCPAdapter struct {
	host     string
	port     int
	listener net.Listener
}

// NewTCPAdapter creates a new TCP adapter instance
func NewTCPAdapter(host string, port int) *TCPAdapter {
	return &TCPAdapter{
		host: host,
		port: port,
	}
}

// Start begins listening for TCP connections
func (a *TCPAdapter) Start() error {
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	
	// Create TCP listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start listener: %w", err)
	}
	
	a.listener = listener
	log.Printf("TCP Adapter listening on %s", address)

	// Accept connections
	for {
		conn, err := a.listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle each connection in a goroutine (concurrent handling)
		go a.handleConnection(conn)
	}
}

// handleConnection processes a single TCP connection
func (a *TCPAdapter) handleConnection(conn net.Conn) {
	h := handler.NewConnectionHandler(conn)
	h.Handle()
}

// Stop gracefully shuts down the adapter
func (a *TCPAdapter) Stop() error {
	if a.listener != nil {
		log.Println("Stopping TCP Adapter...")
		return a.listener.Close()
	}
	return nil
}

// GetAddress returns the current listening address
func (a *TCPAdapter) GetAddress() string {
	if a.listener != nil {
		return a.listener.Addr().String()
	}
	return fmt.Sprintf("%s:%d", a.host, a.port)
}
