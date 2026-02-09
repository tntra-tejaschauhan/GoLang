package protocol

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Message represents a TCP message with simple structure
type Message struct {
	Command string
	Payload string
}

// Encode converts a Message to string format for transmission
// Format: COMMAND:PAYLOAD\n
func (m *Message) Encode() string {
	return fmt.Sprintf("%s:%s\n", m.Command, m.Payload)
}

// Decode reads and parses a message from a reader
func Decode(reader *bufio.Reader) (*Message, error) {
	// Read until newline
	line, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("connection closed")
		}
		return nil, err
	}

	// Remove trailing newline
	line = strings.TrimSpace(line)

	// Split by colon
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid message format: %s", line)
	}

	return &Message{
		Command: parts[0],
		Payload: parts[1],
	}, nil
}

// NewMessage creates a new message
func NewMessage(command, payload string) *Message {
	return &Message{
		Command: command,
		Payload: payload,
	}
}
