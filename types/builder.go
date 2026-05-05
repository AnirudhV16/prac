package types

import "fmt"

/*
Build a Pizza builder:

Pizza struct with fields — size string (required), crust string (required), sauce string, cheese string, toppings []string
PizzaBuilder with a constructor that takes size and crust
WithSauce, WithCheese, WithTopping(topping string) methods — each returns *PizzaBuilder
Build() (*Pizza, error) — validates size and crust are not empty
A Describe() string method on Pizza that prints all its details
In main, build two different pizzas using the builder
*/

type Pizza struct {
	size    string
	crust   string
	sauce   string
	cheese  string
	topping []string
}

type PizzaBuilder struct {
	size    string
	crust   string
	sauce   string
	cheese  string
	topping []string
}

func NewPizzaBuilder(size string, crust string) *PizzaBuilder {
	return &PizzaBuilder{size: size, crust: crust}
}
func (pb *PizzaBuilder) WithSauce(sauce string) *PizzaBuilder {
	pb.sauce = sauce
	return pb
}

func (pb *PizzaBuilder) WithCheese(cheese string) *PizzaBuilder {
	pb.cheese = cheese
	return pb
}

func (pb *PizzaBuilder) WithToppings(toppings ...string) *PizzaBuilder {
	pb.topping = append(pb.topping, toppings...)
	return pb
}

func (pb *PizzaBuilder) Build() (*Pizza, error) {
	if pb.size == "" {
		return nil, fmt.Errorf("size is required")
	}
	if pb.crust == "" {
		return nil, fmt.Errorf("crust is required")
	}
	return &Pizza{
		size:    pb.size,
		crust:   pb.crust,
		sauce:   pb.sauce,
		cheese:  pb.cheese,
		topping: pb.topping,
	}, nil
}

func (p *Pizza) Describe() {
	fmt.Printf("%s %s %s %s %v \n", p.size, p.crust, p.sauce, p.cheese, p.topping)
}
