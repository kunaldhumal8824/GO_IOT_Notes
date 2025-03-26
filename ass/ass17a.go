package main

import (
	"fmt"
)

// Function that returns multiple values: sum, difference, product, and quotient
func calculate(a, b float64) (float64, float64, float64, float64) {
	sum := a + b
	diff := a - b
	product := a * b
	var quotient float64
	if b != 0 {
		quotient = a / b
	} else {
		quotient = 0 // Handle division by zero
	}
	return sum, diff, product, quotient
}

func main() {
	var a, b float64
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	// Call the function and capture the returned values
	sum, diff, product, quotient := calculate(a, b)

	// Print the results
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Difference: %.2f\n", diff)
	fmt.Printf("Product: %.2f\n", product)
	fmt.Printf("Quotient: %.2f\n", quotient)
}
