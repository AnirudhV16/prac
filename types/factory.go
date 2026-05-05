package types

import "fmt"

/*
Build a notification factory:

Notification interface with Notify(message string) error
Three concrete types — EmailNotification, SMSNotification, PushNotification — each with a relevant field and their own Notify implementation
NewNotification(notifType string) (Notification, error) factory function
In main, create all three using the factory and call Notify on each
*/

type Notification interface {
	Notify(message string) error
}

func NewNotification(notifType string) (Notification, error) {
	switch notifType {
	case "email":
		return &EmailNoitification{}, nil
	case "sms":
		return &SMSNoitification{}, nil
	case "push":
		return &PushNoitification{}, nil
	default:
		return nil, fmt.Errorf("no such service")
	}
}

type EmailNoitification struct{}
type SMSNoitification struct{}
type PushNoitification struct{}

func (e *EmailNoitification) Notify(message string) error {
	fmt.Println(message)
	return nil
}

func (e *SMSNoitification) Notify(message string) error {
	fmt.Println(message)
	return nil
}

func (e *PushNoitification) Notify(message string) error {
	fmt.Println(message)
	return nil
}
