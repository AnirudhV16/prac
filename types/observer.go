package types

import "fmt"

/*
Build a stock market alert system:

Observer interface with Update(stock string, price float64)
StockMarket struct that maintains a list of observers and has Subscribe(o Observer), Unsubscribe(o Observer), and SetPrice(stock string, price float64) — SetPrice notifies all observers
Two observers — TraderAlert with a name and PriceLogger
In main, subscribe both, change the price a few times, then unsubscribe TraderAlert and change price again — show it no longer gets notified
*/

//observer interface with update behaviour
type Stocker interface {
	Update(string, float64)
}

//subject observers list not specific(seperate) to the event
type StockMarket struct {
	observers []Stocker
}

func NewStockMarket(observers []Stocker) *StockMarket {
	return &StockMarket{observers: observers}
}

func (s *StockMarket) Subscribe(o Stocker) {
	s.observers = append(s.observers, o)
}

func (s *StockMarket) Unsubscribe(o Stocker) {
	for i, v := range s.observers {
		if v == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *StockMarket) Notify(stock string, price float64) {
	for _, v := range s.observers {
		v.Update(stock, price)
	}
}

func (s *StockMarket) SetPrice(stock string, price float64) {
	s.Notify(stock, price)
}

//observers concrete
type TraderAlert struct {
	name string
}

func NewTradeAlert(name string) *TraderAlert {
	return &TraderAlert{
		name: name,
	}
}

type PriceLogger struct {
}

func NewPriceLogger() *PriceLogger {
	return &PriceLogger{}
}

func (s *TraderAlert) Update(string, float64) {
	fmt.Printf("Alerting the trader %s\n", s.name)
}

func (p *PriceLogger) Update(string, float64) {
	fmt.Printf("logging the price....\n")
}
