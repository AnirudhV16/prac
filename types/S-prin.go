package types

import "fmt"

//core order
type Order struct {
	items []string
	total float64
}

func (o *Order) GetItems() []string {
	return o.items
}

func (o *Order) GetTotal() float64 {
	return o.total
}

//core logic
func (o *Order) AddItem(item string, price float64) {
	o.items = append(o.items, item)
	o.total += price
}

//db logic
type OrderStore struct {
	database string
}

func (o *OrderStore) SaveToDatabase() error {
	// db logic
	return nil
}

type InvoiceGenerator struct {
}

func (i *InvoiceGenerator) GenerateInvoice(o *Order) string {
	return fmt.Sprintf("Items: %v, Total: %.2f", o.GetItems(), o.GetTotal())
}

type MailInvoice struct {
}

func (o *MailInvoice) SendInvoiceEmail(email string) error {
	// email logic
	return nil
}
