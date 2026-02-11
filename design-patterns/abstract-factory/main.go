package main

import (
	"abstract-factory/notification"
	"log"
)

func main() {
	factory, err := notification.NewFactory(notification.AWS)
	if err != nil {
		log.Fatal(err)
	}
	email := factory.CreateEmailSender()
	email.SendEmail("tejas@gmail.com", "hi there")
	sms := factory.CreateSMSSender()
	sms.SendSMS("23463543535", "tejas here")
}
