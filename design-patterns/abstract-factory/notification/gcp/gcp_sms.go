package gcp


import "fmt"

type SMSSender struct{}

func(s *SMSSender) SendSMS(to string, message string) error{
	fmt.Println("GCP SMS send to %s : %s ",to,message)
	return nil
}

