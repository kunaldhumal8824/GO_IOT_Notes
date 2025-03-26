package main

import "fmt"

// Function to print Fibonacci series
func fibonacci(n int) {
	// First two Fibonacci numbers
	a, b := 0, 1

	// Loop to print Fibonacci series
	for i := 0; i < n; i++ {
		// Print the current Fibonacci number
		fmt.Print(a, " ")

		// Update a and b to the next Fibonacci numbers
		a, b = b, a+b
	}
}

func main() {
	// Variable to store the number of terms
	var n int

	// Input from the user
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	// Call the fibonacci function to print the series
	fibonacci(n)
}
