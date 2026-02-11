package notification

import "fmt"

// Concrete Implementations
type SmsNOtification struct{
  Phone string
}

func(s *SmsNOtification) Send(message string) error{
	fmt.Println("Sending SMS to %s : %s \n ", s.Phone,message)
	return nil
} 