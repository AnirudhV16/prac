package types

import "fmt"

type Notify interface {
	Send(message string) error
}

type SMSNotifier struct {
	Phonenumber string
}

type MailNotifier struct {
	Address string
}

type PushNotifier struct {
	Device string
}

//newone
type WhatsappNotifier struct {
	Phonenumber string
}

func (n *SMSNotifier) Send(message string) error {
	fmt.Printf("%s sent to %s", message, n.Phonenumber)
	return nil
}

func (n *MailNotifier) Send(message string) error {
	fmt.Printf("%s sent to %s", message, n.Address)
	return nil
}

func (n *PushNotifier) Send(message string) error {
	fmt.Printf("%s sent to %s", message, n.Device)
	return nil
}

func (n *WhatsappNotifier) Send(message string) error {
	fmt.Printf("%s Whatsapp msg sent to %s", message, n.Phonenumber)
	return nil
}
