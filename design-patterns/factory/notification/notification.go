package notification

// notification default behavior (contract)
type Notification interface{
	Send(message string) error
}