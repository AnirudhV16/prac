package types

import "fmt"

type Notify interface {
	Send(message string) error
}

type SMSNotifier struct {
	PhoneNumber string
}

type MailNotifier struct {
	Address string
}

type PushNotifier struct {
	Device string
}

//newone
type WhatsappNotifier struct {
	PhoneNumber string
}

func (n *SMSNotifier) Send(message string) error {
	fmt.Printf("%s sent to %s\n", message, n.PhoneNumber)
	return nil
}

func (n *MailNotifier) Send(message string) error {
	fmt.Printf("%s sent to %s\n", message, n.Address)
	return nil
}

func (n *PushNotifier) Send(message string) error {
	fmt.Printf("%s sent to %s\n", message, n.Device)
	return nil
}

func (n *WhatsappNotifier) Send(message string) error {
	fmt.Printf("%s Whatsapp msg sent to %s\n", message, n.PhoneNumber)
	return nil
}
