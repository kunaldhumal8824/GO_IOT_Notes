package main

import (
	"fmt"
	"math"
)

// Define the Shape interface with methods Area() and Perimeter()
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type
type Circle struct {
	Radius float64
}

// Rectangle type
type Rectangle struct {
	Length, Width float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Implement the Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Implement the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Length: 4, Width: 6}

	// Print the area and perimeter of the Circle
	fmt.Println("Circle:")
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Perimeter: %.2f\n", circle.Perimeter())

	// Print the area and perimeter of the Rectangle
	fmt.Println("\nRectangle:")
	fmt.Printf("Area: %.2f\n", rectangle.Area())
	fmt.Printf("Perimeter: %.2f\n", rectangle.Perimeter())
}
