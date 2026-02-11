package notification

// This is the core Factory logic

import "errors"

type NotificationType string

const (
	EMAIL NotificationType ="email"
	SMS NotificationType = "sms"
	PUSH NotificationType = "push"
)

// Fectory method

func NewNotification(ntype NotificationType , target string) (Notification, error){

	switch ntype{
	case EMAIL:
		return &EmailNotification{Email: target},nil
	case SMS:
		return &SmsNOtification{Phone:target}, nil
	case PUSH:
		return &PushNotification{DeviceID:target},nil
	default:
		return nil, errors.New("unsupported notification type")

	}



}