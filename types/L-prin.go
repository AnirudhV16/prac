package types

import "fmt"

type TwoDShape interface {
	Area() float64
}

type ThreeDShape interface {
	TwoDShape
	Volume() float64
}

type Circle struct {
	radius float64
}

type Sphere struct {
	radius float64
}

func NewCircle(radius float64) (*Circle, error) {
	var v = Circle{}
	err := v.SetRadius(radius)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c *Circle) GetRadius() float64 {
	return c.radius
}
func (c *Circle) SetRadius(rad float64) error {
	if rad <= 0 {
		return fmt.Errorf("radius should be greater than 0")
	}
	c.radius = rad
	return nil
}

func NewSphere(radius float64) (*Sphere, error) {
	var v = Sphere{}
	err := v.SetRadius(radius)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
func (s *Sphere) Area() float64 {
	return 4 * 3.14 * s.radius * s.radius
}
func (s *Sphere) Volume() float64 {
	return (4.0 / 3.0) * 3.14 * s.radius * s.radius * s.radius
}
func (s *Sphere) GetRadius() float64 {
	return s.radius
}
func (s *Sphere) SetRadius(rad float64) error {
	if rad <= 0 {
		return fmt.Errorf("radius should be greater than 0")
	}
	s.radius = rad
	return nil
}
