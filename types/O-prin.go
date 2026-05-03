package types

type Customer struct {
	Name               string
	DiscountCalculator DiscountCalculator
}

func (c *Customer) CalculateDiscount(price float64) float64 {
	return c.DiscountCalculator.Calculate(price)
}

type DiscountCalculator interface {
	Calculate(price float64) float64
}

type StudentDiscount struct{}
type PremiumDiscount struct{}
type RegularDiscount struct{}

func (r RegularDiscount) Calculate(price float64) float64 {
	return price * 0.95
}

func (r PremiumDiscount) Calculate(price float64) float64 {
	return price * 0.80
}

func (r StudentDiscount) Calculate(price float64) float64 {
	return price * 0.85
}
