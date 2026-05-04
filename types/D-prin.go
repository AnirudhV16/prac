package types

import "fmt"

//need Email Service along with SMS Service
type MessageService interface {
	Send(message string)
}

type SMSService struct{}

func (s SMSService) Send(m string) {
	fmt.Println(m)
}

type EmailService struct{}

func (s EmailService) Send(m string) {
	fmt.Println(m)
}

type OrderNotifier struct {
	sms MessageService
}

func NewOrderNotifier(sms MessageService) *OrderNotifier {
	return &OrderNotifier{sms: sms}
}

func (o *OrderNotifier) NotifyOrder(message string) {
	o.sms.Send(message)
}
