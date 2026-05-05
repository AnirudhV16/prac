package types

/*
Build a Pizza builder:

Pizza struct with fields — size string (required), crust string (required), sauce string, cheese string, toppings []string
PizzaBuilder with a constructor that takes size and crust
WithSauce, WithCheese, WithTopping(topping string) methods — each returns *PizzaBuilder
Build() (*Pizza, error) — validates size and crust are not empty
A Describe() string method on Pizza that prints all its details
In main, build two different pizzas using the builder
*/
