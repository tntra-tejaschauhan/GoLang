package aws

import "fmt"

type EmailSender struct{}

func(a *EmailSender) SendEmail(to string,message string) error{
	fmt.Println("AWS email send to %s : %s",to,message)
	return nil
}

