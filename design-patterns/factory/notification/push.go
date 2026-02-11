package notification

import "fmt"

// Concrete Implementations
type PushNotification struct{
		DeviceID string
}

func (p *PushNotification) Send(message string)error{
	fmt.Println("sending the push notification to %s : %s",p.DeviceID,message)
	return nil
}