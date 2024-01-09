package main

import "fmt"

// Shape is an interface with a single method, Area()
type Shape interface {
	Area() float64
}

// Square type implementing the Shape interface
type Square struct {
	Side float64
}

// Area method for Square satisfying the Shape interface
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Circle type implementing the Shape interface
type Circle struct {
	Radius float64
}

// Area method for Circle satisfying the Shape interface
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func testInterfaces() {
	// Create instances of Square and Circle
	square := Square{Side: 5}
	circle := Circle{Radius: 3}

	// Use the Shape interface to calculate areas
	printArea(square)
	printArea(circle)
}

// printArea is a function that accepts any type satisfying the Shape interface
func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}
