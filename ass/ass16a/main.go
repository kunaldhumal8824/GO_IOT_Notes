package main

import (
	"fmt"
	"your_project/rectangle" // Replace with your actual module name
)

func main() {
	var length, width float64
	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scan(&length)
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&width)

	// Call the Area function from the rectangle package
	area := rectangle.Area(length, width)

	// Print the result
	fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
