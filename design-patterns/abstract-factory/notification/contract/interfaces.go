package contract


// product A
type EmailSender interface{
	SendEmail(to string,message string) error
}

// product B

type SMSSender interface{
	SendSMS(to string, message string) error
}

// Abstract factory interface

type NotificationFactory interface{
	CreateEmailSender() EmailSender
	CreateSMSSender() SMSSender
}
