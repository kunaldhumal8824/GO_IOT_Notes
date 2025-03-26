package main

import (
	"fmt"
)

// Define the Shape interface with methods Area() and Volume()
type Shape interface {
	Area() float64
	Volume() float64
}

// Square type (2D shape)
type Square struct {
	Side float64
}

// Rectangle type (2D shape)
type Rectangle struct {
	Length, Width float64
}

// Implement the Area method for Square
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Implement the Volume method for Square (In 2D shapes like square, volume is 0)
func (s Square) Volume() float64 {
	return 0
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Implement the Volume method for Rectangle (In 2D shapes like rectangle, volume is 0)
func (r Rectangle) Volume() float64 {
	return 0
}

// Function to print the shape details
func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Volume: %.2f\n", s.Volume())
}

func main() {
	// Create instances of Square and Rectangle
	square := Square{Side: 5}
	rectangle := Rectangle{Length: 4, Width: 6}

	// Print details of Square
	fmt.Println("Square:")
	printShapeDetails(square)

	// Print details of Rectangle
	fmt.Println("\nRectangle:")
	printShapeDetails(rectangle)
}
