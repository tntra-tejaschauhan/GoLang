package handler

import (
	"bufio"
	"log"
	"net"
	"strings"
	"tcp-adapter/pkg/protocol"
)

// ConnectionHandler handles individual TCP connections
type ConnectionHandler struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

// NewConnectionHandler creates a new connection handler
func NewConnectionHandler(conn net.Conn) *ConnectionHandler {
	return &ConnectionHandler{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
}

// Handle processes messages from the connection
func (h *ConnectionHandler) Handle() {
	defer h.conn.Close()
	
	clientAddr := h.conn.RemoteAddr().String()
	log.Printf("New connection from: %s", clientAddr)

	// Send welcome message
	welcome := protocol.NewMessage("WELCOME", "Connected to TCP Adapter Server")
	h.SendMessage(welcome)

	// Main message loop
	for {
		msg, err := protocol.Decode(h.reader)
		if err != nil {
			log.Printf("Connection closed from %s: %v", clientAddr, err)
			return
		}

		log.Printf("Received from %s - Command: %s, Payload: %s", 
			clientAddr, msg.Command, msg.Payload)

		// Process the message
		response := h.processMessage(msg)
		h.SendMessage(response)
	}
}

// processMessage handles different message commands
func (h *ConnectionHandler) processMessage(msg *protocol.Message) *protocol.Message {
	switch strings.ToUpper(msg.Command) {
	case "ECHO":
		return protocol.NewMessage("ECHO_RESPONSE", msg.Payload)
	
	case "UPPER":
		return protocol.NewMessage("UPPER_RESPONSE", strings.ToUpper(msg.Payload))
	
	case "LOWER":
		return protocol.NewMessage("LOWER_RESPONSE", strings.ToLower(msg.Payload))
	
	case "REVERSE":
		reversed := reverseString(msg.Payload)
		return protocol.NewMessage("REVERSE_RESPONSE", reversed)
	
	case "PING":
		return protocol.NewMessage("PONG", "alive")
	
	case "QUIT":
		return protocol.NewMessage("BYE", "Goodbye!")
	
	default:
		return protocol.NewMessage("ERROR", "Unknown command: "+msg.Command)
	}
}

// SendMessage sends a message to the client
func (h *ConnectionHandler) SendMessage(msg *protocol.Message) error {
	_, err := h.writer.WriteString(msg.Encode())
	if err != nil {
		return err
	}
	return h.writer.Flush()
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
