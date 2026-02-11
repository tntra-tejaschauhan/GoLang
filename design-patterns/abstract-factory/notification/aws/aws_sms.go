package aws

import "fmt"

type SMSSender struct{}

func (a *SMSSender) SendSMS(to string, messsage string) error{
	fmt.Println("sending sms to %s : %s", to,messsage)
	return nil
}