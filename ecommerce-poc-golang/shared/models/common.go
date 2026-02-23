package models

// Fixed users for the POC
var FixedUsers = map[string]string{
	"user-1": "Alice",
	"user-2": "Bob",
	"user-3": "Charlie",
	"user-4": "Diana",
}

func IsValidUser(userID string) bool {
	_, exists := FixedUsers[userID]
	return exists
}

// Common response structures
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
