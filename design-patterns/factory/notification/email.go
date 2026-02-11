package notification

import "fmt"

// Concrete Implementations
type EmailNotification struct{
	Email string
	
}

func (e *EmailNotification) Send(message string) error{
fmt.Println("Sending email to %s:  %s\n",e.Email,message)
return nil
}