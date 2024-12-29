package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Implement Shape methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct
type Circle struct {
	Radius float64
}

// Implement Shape methods for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var s Shape

	r := Rectangle{Width: 5, Height: 4}
	c := Circle{Radius: 3}

	// Assign Rectangle to Shape
	s = r
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())

	// Assign Circle to Shape
	s = c
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())
}
