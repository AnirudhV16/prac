package types

import "fmt"

/*
You're building a payment system. Your system expects:
gotype PaymentGateway interface {
    Pay(amount float64) string
}
But you have a third-party Razorpay SDK with this interface:
gotype RazorpaySDK struct{}

func (r *RazorpaySDK) ProcessPayment(amountInPaise int, currency string) string {
    return fmt.Sprintf("Razorpay processed %d paise in %s", amountInPaise, currency)
}
The interfaces don't match — your system uses float64 rupees, Razorpay uses int paise and needs a currency string.
Write an adapter that makes RazorpaySDK work with your PaymentGateway interface. The adapter should convert rupees to paise (multiply by 100) and hardcode currency as "INR".
*/

type PaymentGateway interface {
	Pay(amount float64) string
}

type RazorpaySDK struct{}

func (r *RazorpaySDK) ProcessPayment(amountInPaise int, currency string) string {
	return fmt.Sprintf("Razorpay processed %d paise in %s", amountInPaise, currency)
}

type adapter struct {
	r RazorpaySDK
}

func (a *adapter) Pay(amount float64) string {
	return a.r.ProcessPayment(int(amount), "INR")
}
