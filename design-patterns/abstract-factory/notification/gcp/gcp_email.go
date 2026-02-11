package gcp

import "fmt"

type EmailSender struct{}

func (g *EmailSender) SendEmail(to string,message string) error{
	fmt.Println("GCP email send to %s : %s",to,message)
	return nil
}

